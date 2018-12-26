package validator

import (
	"strings"

	validator "gopkg.in/go-playground/validator.v9"
)

func isScript(fl validator.FieldLevel) bool {
	index := strings.Index(fl.Field().String(), "$")
	return index != -1
}
