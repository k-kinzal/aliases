package validator

import (
	"fmt"

	"github.com/iancoleman/strcase"
	validator "gopkg.in/go-playground/validator.v9"
)

type Validate struct {
	*validator.Validate
}

func (v *Validate) Struct(s interface{}) error {
	err := v.Validate.Struct(s)
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			switch e.Tag() {
			case "bool|script":
				return fmt.Errorf("invalid parameter `%s` for `%s` is not bool or script string", e.Value(), strcase.ToLowerCamel(e.Field()))
			case "int|script":
				return fmt.Errorf("invalid parameter `%s` for `%s` is not int or script string", e.Value(), strcase.ToLowerCamel(e.Field()))
			case "int64|script":
				return fmt.Errorf("invalid parameter `%s` for `%s` is not int64 or script string", e.Value(), strcase.ToLowerCamel(e.Field()))
			case "uint16|script":
				return fmt.Errorf("invalid parameter `%s` for `%s` is not uint16 or script string", e.Value(), strcase.ToLowerCamel(e.Field()))
			case "uint64|script":
				return fmt.Errorf("invalid parameter `%s` for `%s` is not uint64 or script string", e.Value(), strcase.ToLowerCamel(e.Field()))
			case "duration|script":
				return fmt.Errorf("invalid parameter `%s` for `%s` is not time duration or script string", e.Value(), strcase.ToLowerCamel(e.Field()))
			case "nanocpus|script":
				return fmt.Errorf("invalid parameter `%s` for `%s` is not CPU or script string (e.g. 0.5, 1)", e.Value(), strcase.ToLowerCamel(e.Field()))
			case "membytes|script":
				return fmt.Errorf("invalid parameter `%s` for `%s` is not memory or script string (e.g. 2MB, 2GB)", e.Value(), strcase.ToLowerCamel(e.Field()))
			case "memswapbytes|script":
				return fmt.Errorf("invalid parameter `%s` for `%s` is not swap memory or script string (e.g. -1, 2MB, 2GB)", e.Value(), strcase.ToLowerCamel(e.Field()))
			case "required":
				return fmt.Errorf("invalid parameter for `%s` is required", strcase.ToLowerCamel(e.Field()))
			}
		}
	}

	return err
}

func New() (*Validate, error) {
	validate := validator.New()
	if err := validate.RegisterValidation("bool", isBoolean); err != nil {
		return nil, fmt.Errorf("logic error: %s", err)
	}
	if err := validate.RegisterValidation("int", isInteger); err != nil {
		return nil, fmt.Errorf("logic error: %s", err)
	}
	if err := validate.RegisterValidation("int64", isInteger64); err != nil {
		return nil, fmt.Errorf("logic error: %s", err)
	}
	if err := validate.RegisterValidation("uint16", isUnsignedInteger16); err != nil {
		return nil, fmt.Errorf("logic error: %s", err)
	}
	if err := validate.RegisterValidation("uint64", isUnsignedInteger64); err != nil {
		return nil, fmt.Errorf("logic error: %s", err)
	}
	if err := validate.RegisterValidation("duration", isDuration); err != nil {
		return nil, fmt.Errorf("logic error: %s", err)
	}
	if err := validate.RegisterValidation("nanocpus", isNanoCPUs); err != nil {
		return nil, fmt.Errorf("logic error: %s", err)
	}
	if err := validate.RegisterValidation("membytes", isMemoryBytes); err != nil {
		return nil, fmt.Errorf("logic error: %s", err)
	}
	if err := validate.RegisterValidation("memswapbytes", isMemorySwapBytes); err != nil {
		return nil, fmt.Errorf("logic error: %s", err)
	}
	if err := validate.RegisterValidation("script", isScript); err != nil {
		return nil, fmt.Errorf("logic error: %s", err)
	}

	v := Validate{validate}
	return &v, nil
}
