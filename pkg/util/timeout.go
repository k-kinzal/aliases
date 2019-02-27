package util

import (
	"time"
)

// Timeout stops the slow function at the specified time
func Timeout(t time.Duration, fn func() error) error {
	ch := make(chan error, 1)
	go func() {
		ch <- fn()
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(t):
		return &TimeoutError{"timeout"}
	}
}
