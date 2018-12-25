package validator

import (
	"math/big"

	units "github.com/docker/go-units"
	validator "gopkg.in/go-playground/validator.v9"
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
	if _, err := units.RAMInBytes(fl.Field().String()); err != nil {
		return false
	}
	return true
}

func isMemorySwapBytes(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	if val == "-1" {
		return true
	}
	if _, err := units.RAMInBytes(val); err != nil {
		return false
	}
	return true
}
