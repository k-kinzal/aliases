package validator

import validator "gopkg.in/go-playground/validator.v9"

func New() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("bool", isBoolean)
	validate.RegisterValidation("int", isInteger)
	validate.RegisterValidation("int64", isInteger64)
	validate.RegisterValidation("uint16", isUnsignedInteger16)
	validate.RegisterValidation("uint64", isUnsignedInteger64)
	validate.RegisterValidation("duration", isDuration)
	validate.RegisterValidation("nanocpus", isNanoCPUs)
	validate.RegisterValidation("membytes", isMemoryBytes)
	validate.RegisterValidation("memswapbytes", isMemorySwapBytes)
	validate.RegisterValidation("script", isScript)

	return validate
}
