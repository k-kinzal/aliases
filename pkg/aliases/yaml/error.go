package yaml

import (
	"fmt"
	"strings"
)

// YAMLError is an error occurred in YAML-encoded data.
type YAMLError struct {
	err error
}

// Error returns error message.
func (e *YAMLError) Error() string {
	return fmt.Sprintf("yaml error: %s", strings.Replace(fmt.Sprintf("%s", e.err), "yaml error: ", "", -1))
}

// Error return new YamlError.
func Error(i interface{}) error {
	if i, ok := i.(error); ok {
		return &YAMLError{i}
	}

	return &YAMLError{fmt.Errorf("%s", i)}
}

// Errorf return new YamlError with formated message.
func Errorf(format string, a ...interface{}) error {
	return &YAMLError{fmt.Errorf(format, a...)}
}
