package cmd

import "fmt"

// InvalidArgumentError is an invalid flag parameter error.
type InvalidFlagError struct {
	name   string
	value  interface{}
	reason string
}

// Error returns error message.
func (e *InvalidFlagError) Error() string {
	return fmt.Sprintf("invalid argument \"%s\" for \"%s\" flag: %s", e.value, e.name, e.reason)
}

// Error return new InvalidFlagError.
func FlagError(name string, value interface{}, reason string) error {
	return &InvalidFlagError{name, value, reason}
}
