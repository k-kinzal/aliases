package yaml

import (
	"fmt"
	"math/big"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/iancoleman/strcase"
	validator "gopkg.in/go-playground/validator.v9"
)

// separator returns always true. because since `a|b,c|d` is `a|b|c|d`, it is like `a|b,.,c|d` as a delimiter.
func separator(fl validator.FieldLevel) bool {
	return true
}

// hasEnvironmentVariable is check that `$parameter` or `${parameter}` is included.
func hasEnvironmentVariable(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	dst := os.ExpandEnv(fl.Field().String())
	return dst != val
}

// asMax is check that less than or equal the parameter.
func asMax(fl validator.FieldLevel) bool {
	param := fl.Param()
	val := fl.Field().String()
	// 16 bit
	if n1, err := strconv.ParseInt(param, 10, 16); err == nil {
		if n2, e := strconv.ParseInt(val, 10, 16); e == nil {
			return n1 >= n2
		}
	}
	if n1, err := strconv.ParseUint(param, 10, 16); err == nil {
		if n2, e := strconv.ParseUint(val, 10, 16); e == nil {
			return n1 >= n2
		}
	}
	// 32 bit
	if n1, err := strconv.ParseInt(param, 10, 32); err == nil {
		if n2, e := strconv.ParseInt(val, 10, 32); e == nil {
			return n1 >= n2
		}
	}
	if n1, err := strconv.ParseUint(param, 10, 32); err == nil {
		if n2, e := strconv.ParseUint(val, 10, 32); e == nil {
			return n1 >= n2
		}
	}
	// 64 bit
	if n1, err := strconv.ParseInt(param, 10, 64); err == nil {
		if n2, e := strconv.ParseInt(val, 10, 64); e == nil {
			return n1 >= n2
		}
	}
	if n1, err := strconv.ParseUint(param, 10, 64); err == nil {
		if n2, e := strconv.ParseUint(val, 10, 64); e == nil {
			return n1 >= n2
		}
	}

	return false
}

// asMax is check that greater than or equal the parameter.
func asMin(fl validator.FieldLevel) bool {
	param := fl.Param()
	val := fl.Field().String()
	// 16 bit
	if n1, err := strconv.ParseInt(param, 10, 16); err == nil {
		if n2, e := strconv.ParseInt(val, 10, 16); e == nil {
			return n1 <= n2
		}
	}
	if n1, err := strconv.ParseUint(param, 10, 16); err == nil {
		if n2, e := strconv.ParseUint(val, 10, 16); e == nil {
			return n1 <= n2
		}
	}
	// 32 bit
	if n1, err := strconv.ParseInt(param, 10, 32); err == nil {
		if n2, e := strconv.ParseInt(val, 10, 32); e == nil {
			return n1 <= n2
		}
	}
	if n1, err := strconv.ParseUint(param, 10, 32); err == nil {
		if n2, e := strconv.ParseUint(val, 10, 32); e == nil {
			return n1 <= n2
		}
	}
	// 64 bit
	if n1, err := strconv.ParseInt(param, 10, 64); err == nil {
		if n2, e := strconv.ParseInt(val, 10, 64); e == nil {
			return n1 <= n2
		}
	}
	if n1, err := strconv.ParseUint(param, 10, 64); err == nil {
		if n2, e := strconv.ParseUint(val, 10, 64); e == nil {
			return n1 <= n2
		}
	}

	return false
}

// isNanoCPUs is check that Docker's nano CPU format.
func isNanoCPUs(fl validator.FieldLevel) bool {
	cpu, ok := new(big.Rat).SetString(fl.Field().String())
	if !ok {
		return false
	}
	nano := cpu.Mul(cpu, big.NewRat(1e9, 1))
	if !nano.IsInt() {
		return false
	}
	min := new(big.Rat).SetFloat64(0.01)
	if nano.Cmp(min.Mul(min, big.NewRat(1e9, 1))) == -1 {
		return false
	}
	max := new(big.Rat).SetInt64((int64)(runtime.NumCPU()))
	if nano.Cmp(max.Mul(max, big.NewRat(1e9, 1))) == 1 {
		return false
	}

	return true
}

var (
	sizeRegex = regexp.MustCompile(`^(\d+(\.\d+)*) ?([kKmMgGtTpP])?[iI]?[bB]?$`)
)

// isMemoryBytes is check that Docker's memory byte format.
func isMemoryBytes(fl validator.FieldLevel) bool {
	return sizeRegex.MatchString(fl.Field().String())
}

// isMemoryBytes is check that Docker's memory byte format or -1.
func isMemorySwapBytes(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	if val == "-1" {
		return true
	}
	return sizeRegex.MatchString(val)
}

// isBoolean is check that bool type,
func isBoolean(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	switch val {
	case "true":
		return true
	case "false":
		return true
	}
	return false
}

// isBoolean is check that int type,
func isInteger(fl validator.FieldLevel) bool {
	num, err := strconv.ParseInt(fl.Field().String(), 10, 32)
	if err != nil {
		return false
	}
	return fl.Field().String() == strconv.FormatInt(num, 10)
}

// isBoolean is check that int64 type,
func isInteger64(fl validator.FieldLevel) bool {
	num, err := strconv.ParseInt(fl.Field().String(), 10, 64)
	if err != nil {
		return false
	}
	return fl.Field().String() == strconv.FormatInt(num, 10)
}

// isBoolean is check that uint16 type,
func isUnsignedInteger16(fl validator.FieldLevel) bool {
	num, err := strconv.ParseUint(fl.Field().String(), 10, 16)
	if err != nil {
		return false
	}
	return fl.Field().String() == strconv.FormatUint(num, 10)
}

// isBoolean is check that uint64 type,
func isUnsignedInteger64(fl validator.FieldLevel) bool {
	num, err := strconv.ParseUint(fl.Field().String(), 10, 64)
	if err != nil {
		return false
	}
	return fl.Field().String() == strconv.FormatUint(num, 10)
}

// isBoolean is check that duration type,
func isDuration(fl validator.FieldLevel) bool {
	if _, err := time.ParseDuration(fl.Field().String()); err != nil {
		return false
	}
	return true
}

// Validate is wrapper to validator.v9.
type Validate struct {
	*validator.Validate
}

var (
	indexRegexp = regexp.MustCompile(`^(.*?)(\[\d+\])+$`)
)

// Struct is wrapper to validator.v9:Struct.
func (v *Validate) Struct(s interface{}) error {
	err := v.Validate.Struct(s)
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			tag := e.Tag()
			tag = strings.Replace(tag, "|env", "", -1)
			tag = strings.Replace(tag, "env|", "", -1)
			tag = strings.Replace(tag, fmt.Sprintf("=%s", e.Param()), "", -1)
			// tag = paramRegexp.ReplaceAllString(tag, "")
			switch tag {
			case "bool":
				return fmt.Errorf("invalid parameter `%s` for `%s` is not bool or environment variables", e.Value(), strcase.ToLowerCamel(e.Field()))
			case "int":
				return fmt.Errorf("invalid parameter `%s` for `%s` is not int or environment variables", e.Value(), strcase.ToLowerCamel(e.Field()))
			case "int64":
				return fmt.Errorf("invalid parameter `%s` for `%s` is not int64 or environment variables", e.Value(), strcase.ToLowerCamel(e.Field()))
			case "uint16":
				return fmt.Errorf("invalid parameter `%s` for `%s` is not uint16 or environment variables", e.Value(), strcase.ToLowerCamel(e.Field()))
			case "uint64":
				return fmt.Errorf("invalid parameter `%s` for `%s` is not uint64 or environment variables", e.Value(), strcase.ToLowerCamel(e.Field()))
			case "duration":
				return fmt.Errorf("invalid parameter `%s` for `%s` is not time duration or environment variables", e.Value(), strcase.ToLowerCamel(e.Field()))
			case "nanocpus":
				return fmt.Errorf("invalid parameter `%s` for `%s` is not CPU or environment variables (e.g. 0.5, 1)", e.Value(), strcase.ToLowerCamel(e.Field()))
			case "membytes":
				return fmt.Errorf("invalid parameter `%s` for `%s` is not memory or environment variables (e.g. 2MB, 2GB)", e.Value(), strcase.ToLowerCamel(e.Field()))
			case "memswapbytes":
				return fmt.Errorf("invalid parameter `%s` for `%s` is not swap memory or environment variables (e.g. -1, 2MB, 2GB)", e.Value(), strcase.ToLowerCamel(e.Field()))
			case "oneof":
				field := e.Field()
				matches := indexRegexp.FindStringSubmatch(field)
				if matches != nil {
					field = fmt.Sprintf("%s%s", strcase.ToLowerCamel(matches[1]), strings.Join(matches[2:], ""))
				}
				return fmt.Errorf("invalid parameter `%s` for `%s` is one of %s", e.Value(), field, strings.Replace(e.Param(), " ", ", ", -1))
			case "max":
				return fmt.Errorf("invalid parameter `%s` for `%s` is a number less than or equal to `%s`", e.Value(), strcase.ToLowerCamel(e.Field()), e.Param())
			case "min":
				return fmt.Errorf("invalid parameter `%s` for `%s` is a number greater than or equal to `%s`", e.Value(), strcase.ToLowerCamel(e.Field()), e.Param())
			case "required":
				return fmt.Errorf("invalid parameter for `%s` is required", strcase.ToLowerCamel(e.Field()))
			case "ipv4":
				return fmt.Errorf("invalid parameter `%s` for `%s` is not IPv4 format (e.g., 172.30.100.104)", e.Value(), strings.ToLower(e.Field()))
			case "ipv6":
				return fmt.Errorf("invalid parameter `%s` for `%s` is not IPv6 format (e.g., 2001:db8::33)", e.Value(), strings.ToLower(e.Field()))
			case "mac":
				return fmt.Errorf("invalid parameter `%s` for `%s` is not mac address format", e.Value(), strcase.ToLowerCamel(e.Field()))
			default:
				return fmt.Errorf("%#v", e)
			}

		}
	}

	return nil
}

// NewValidator creates a new Validator.
func NewValidator() (*Validate, error) {
	validate := validator.New()
	if err := validate.RegisterValidation("_", separator); err != nil {
		return nil, fmt.Errorf("logic error: %s", err)
	}
	if err := validate.RegisterValidation("env", hasEnvironmentVariable); err != nil {
		return nil, fmt.Errorf("logic error: %s", err)
	}
	if err := validate.RegisterValidation("max", asMax); err != nil {
		return nil, fmt.Errorf("logic error: %s", err)
	}
	if err := validate.RegisterValidation("min", asMin); err != nil {
		return nil, fmt.Errorf("logic error: %s", err)
	}
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

	v := Validate{validate}
	return &v, nil
}
