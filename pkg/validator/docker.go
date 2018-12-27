package validator

import (
	"math/big"
	"regexp"

	validator "gopkg.in/go-playground/validator.v9"
)

var (
	sizeRegex = regexp.MustCompile(`^(\d+(\.\d+)*) ?([kKmMgGtTpP])?[iI]?[bB]?$`)
)

func isNanoCPUs(fl validator.FieldLevel) bool {
	cpu, ok := new(big.Rat).SetString(fl.Field().String())
	if !ok {
		return false
	}
	if nano := cpu.Mul(cpu, big.NewRat(1e9, 1)); !nano.IsInt() {
		return false
	}
	return true
}

func isMemoryBytes(fl validator.FieldLevel) bool {
	return sizeRegex.MatchString(fl.Field().String())
}

func isMemorySwapBytes(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	if val == "-1" {
		return true
	}
	return sizeRegex.MatchString(val)
}
