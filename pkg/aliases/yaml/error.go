package yaml

import (
	"fmt"
	"strings"
)

// UnmarshalError is an error occurred in YAML-encoded data.
type UnmarshalError struct {
	err error
}

// Error returns error message.
func (e *UnmarshalError) Error() string {
	return fmt.Sprintf("yaml error: %s", strings.Replace(fmt.Sprintf("%s", e.err), "yaml error: ", "", -1))
}

// Error return new UnmarshalError.
func Error(i interface{}) error {
	if i, ok := i.(error); ok {
		return &UnmarshalError{i}
	}

	return &UnmarshalError{fmt.Errorf("%s", i)}
}

// Errorf return new UnmarshalError with formated message.
func Errorf(format string, a ...interface{}) error {
	return &UnmarshalError{fmt.Errorf(format, a...)}
}
