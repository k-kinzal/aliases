package util

import (
	"io/ioutil"
	"os"
)

func IsFilePath(name string) bool {
	if _, err := os.Stat(name); err == nil {
		return true
	}
	var d []byte
	if err := ioutil.WriteFile(name, d, 0644); err == nil {
		_ = os.Remove(name)
		return true
	}

	return false
}
