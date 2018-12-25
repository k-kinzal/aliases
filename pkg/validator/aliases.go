package validator

import (
	"strings"

	validator "gopkg.in/go-playground/validator.v9"
)

func isScript(fl validator.FieldLevel) bool {
	return strings.Index(fl.Field().String(), "$") != -1
}
