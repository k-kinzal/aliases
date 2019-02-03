package yaml_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/k-kinzal/aliases/pkg/aliases/yaml"
)

func ExampleNewValidator() {
	i := struct {
		Data string `validate:"_"`
	}{
		Data: "foo",
	}
	v, err := yaml.NewValidator()
	if err != nil {
		panic(err)
	}

	if err := v.Struct(i); err != nil {
		panic(err)
	} else {
		fmt.Println(err)
	}

	// Output: <nil>
}

func TestValidate_StructSuccessIsShell(t *testing.T) {
	i := struct {
		Data1 string `validate:"shell"`
		Data2 string `validate:"shell"`
		Data3 string `validate:"shell"`
		Data4 string `validate:"shell"`
		Data5 string `validate:"shell"`
	}{
		Data1: "Hello $FOO",
		Data2: "Hello ${FOO}",
		Data3: "Hello $(echo 'foo')",
		Data4: "Hello `echo 'foo'`",
		Data5: "$(echo \")\")",
	}

	validator, _ := yaml.NewValidator()
	if err := validator.Struct(i); err != nil {
		t.Errorf("expect `<nil>`, but actual `%s`", err)
	}
}

func TestValidate_StructFailedIsShell(t *testing.T) {
	i := struct {
		Data1 string `validate:"shell"`
		Data2 string `validate:"shell"`
		Data3 string `validate:"shell"`
		Data4 string `validate:"shell"`
		Data5 string `validate:"shell"`
		Data6 string `validate:"shell"`
	}{
		Data1: "Hello World",
		Data2: "$",
		Data3: "${}",
		Data4: "$(",
		Data5: "$()",
		Data6: "$(echo (123)",
	}

	validator, _ := yaml.NewValidator()
	if err := validator.Struct(i); err == nil {
		t.Error("expected that an error will occur, but actual not error occur")
	}
}

func TestValidate_StructSuccessAsMax(t *testing.T) {
	i := struct {
		// int
		Data1 string `validate:"max=32767"`
		Data2 string `validate:"max=32767"`
		Data3 string `validate:"max=2147483647"`
		Data4 string `validate:"max=2147483647"`
		Data5 string `validate:"max=9223372036854775807"`
		//Data6 string `validate:"max=9223372036854775807"` // is no type to compare
		// uint
		Data7  string `validate:"max=65535"`
		Data8  string `validate:"max=65535"`
		Data9  string `validate:"max=4294967295"`
		Data10 string `validate:"max=4294967295"`
		Data11 string `validate:"max=18446744073709551615"`
		//Data12 string `validate:"max=18446744073709551615"` // is no type to compare
	}{
		Data1: "32767",
		Data2: "-32769",
		Data3: "2147483647",
		Data4: "-2147483649",
		Data5: "9223372036854775807",
		//Data6:  "-9223372036854775809",
		Data7:  "65535",
		Data8:  "-1",
		Data9:  "4294967295",
		Data10: "-1",
		Data11: "18446744073709551615",
		//Data12: "-1",
	}

	validator, _ := yaml.NewValidator()
	if err := validator.Struct(i); err != nil {
		t.Errorf("expect `<nil>`, but actual `%s`", err)
	}
}

func TestValidate_StructFailedAsMax(t *testing.T) {
	i := struct {
		// int
		Data1 string `validate:"max=32767"`
		Data2 string `validate:"max=2147483647"`
		Data3 string `validate:"max=9223372036854775807"`
		// uint
		Data4 string `validate:"max=65535"`
		Data5 string `validate:"max=4294967295"`
		Data6 string `validate:"max=18446744073709551615"`
	}{
		Data1: "32768",
		Data2: "2147483648",
		Data3: "9223372036854775808",
		Data4: "65536",
		Data5: "4294967296",
		Data6: "18446744073709551616",
	}

	validator, _ := yaml.NewValidator()
	if err := validator.Struct(i); err == nil {
		t.Error("expected that an error will occur, but actual not error occur")
	}
}

func TestValidate_StructSuccessAsMin(t *testing.T) {
	i := struct {
		// int
		Data1 string `validate:"min=-32768"`
		Data2 string `validate:"min=-32768"`
		Data3 string `validate:"min=-2147483648"`
		Data4 string `validate:"min=-2147483648"`
		Data5 string `validate:"min=-9223372036854775808"`
		//Data6 string `validate:"min=-9223372036854775808"` // is no type to compare
		// uint
		Data7  string `validate:"min=0"`
		Data8  string `validate:"min=0"`
		Data9  string `validate:"min=0"`
		Data10 string `validate:"min=0"`
		Data11 string `validate:"min=0"`
		//Data12 string `validate:"min=0"` // is no type to compare
	}{
		Data1: "-32768",
		Data2: "32768",
		Data3: "-2147483648",
		Data4: "2147483648",
		Data5: "-9223372036854775808",
		//Data6:  "9223372036854775808",
		Data7:  "0",
		Data8:  "65536",
		Data9:  "0",
		Data10: "4294967296",
		Data11: "0",
		//Data12: "18446744073709551616",
	}

	validator, _ := yaml.NewValidator()
	if err := validator.Struct(i); err != nil {
		t.Errorf("expect `<nil>`, but actual `%s`", err)
	}
}

func TestValidate_StructFailedAsMin(t *testing.T) {
	i := struct {
		// int
		Data1 string `validate:"min=-32768"`
		Data2 string `validate:"min=-2147483648"`
		Data3 string `validate:"min=-9223372036854775808"`
		// uint
		Data4 string `validate:"min=0"`
		Data5 string `validate:"min=0"`
		Data6 string `validate:"min=0"`
	}{
		Data1: "-32769",
		Data2: "-2147483649",
		Data3: "-9223372036854775809",
		Data4: "-1",
		Data5: "-1",
		Data6: "-1",
	}

	validator, _ := yaml.NewValidator()
	if err := validator.Struct(i); err == nil {
		t.Error("expected that an error will occur, but actual not error occur")
	}
}

func TestValidate_StructSuccessIsNanoCPUs(t *testing.T) {
	i := struct {
		Data1 string `validate:"nanocpus"`
		Data2 string `validate:"nanocpus"`
		Data3 string `validate:"nanocpus"`
	}{
		Data1: "1",
		Data2: "0.5",
		Data3: ".5",
	}

	validator, _ := yaml.NewValidator()
	if err := validator.Struct(i); err != nil {
		t.Errorf("expect `<nil>`, but actual `%s`", err)
	}
}

func TestValidate_StructFailedIsNanoCPUs(t *testing.T) {
	i := struct {
		Data1 string `validate:"nanocpus"`
		Data2 string `validate:"nanocpus"`
	}{
		Data1: "0.001",
		Data2: fmt.Sprintf("%d", runtime.NumCPU()+1),
	}

	validator, _ := yaml.NewValidator()
	if err := validator.Struct(i); err == nil {
		t.Error("expected that an error will occur, but actual not error occur")
	}
}

func TestValidate_StructSuccessIsMemoryBytes(t *testing.T) {
	i := struct {
		Data1  string `validate:"membytes"`
		Data2  string `validate:"membytes"`
		Data3  string `validate:"membytes"`
		Data4  string `validate:"membytes"`
		Data5  string `validate:"membytes"`
		Data6  string `validate:"membytes"`
		Data7  string `validate:"membytes"`
		Data8  string `validate:"membytes"`
		Data9  string `validate:"membytes"`
		Data10 string `validate:"membytes"`
		Data11 string `validate:"membytes"`
	}{
		Data1:  "1",
		Data2:  "1B",
		Data3:  "1KB",
		Data4:  "1MB",
		Data5:  "1GB",
		Data6:  "1PB",
		Data7:  "1KiB",
		Data8:  "1MiB",
		Data9:  "1GiB",
		Data10: "1PiB",
		Data11: "1.2",
	}

	validator, _ := yaml.NewValidator()
	if err := validator.Struct(i); err != nil {
		t.Errorf("expect `<nil>`, but actual `%s`", err)
	}
}

func TestValidate_StructFailedIsMemoryBytes(t *testing.T) {
	i := struct {
		Data1 string `validate:"membytes"`
		Data2 string `validate:"membytes"`
	}{
		Data1: "-1",
		Data2: "foo",
	}

	validator, _ := yaml.NewValidator()
	if err := validator.Struct(i); err == nil {
		t.Error("expected that an error will occur, but actual not error occur")
	}
}

func TestValidate_StructSuccessIsMemorySwapBytes(t *testing.T) {
	i := struct {
		Data1  string `validate:"memswapbytes"`
		Data2  string `validate:"memswapbytes"`
		Data3  string `validate:"memswapbytes"`
		Data4  string `validate:"memswapbytes"`
		Data5  string `validate:"memswapbytes"`
		Data6  string `validate:"memswapbytes"`
		Data7  string `validate:"memswapbytes"`
		Data8  string `validate:"memswapbytes"`
		Data9  string `validate:"memswapbytes"`
		Data10 string `validate:"memswapbytes"`
		Data11 string `validate:"memswapbytes"`
	}{
		Data1:  "1",
		Data2:  "1B",
		Data3:  "1KB",
		Data4:  "1MB",
		Data5:  "1GB",
		Data6:  "1PB",
		Data7:  "1KiB",
		Data8:  "1MiB",
		Data9:  "1GiB",
		Data10: "1PiB",
		Data11: "-1",
	}

	validator, _ := yaml.NewValidator()
	if err := validator.Struct(i); err != nil {
		t.Errorf("expect `<nil>`, but actual `%s`", err)
	}
}

func TestValidate_StructFailedIsMemorySwapBytes(t *testing.T) {
	i := struct {
		Data1 string `validate:"memswapbytes"`
	}{
		Data1: "foo",
	}

	validator, _ := yaml.NewValidator()
	if err := validator.Struct(i); err == nil {
		t.Error("expected that an error will occur, but actual not error occur")
	}
}

func TestValidate_StructSuccessIsBoolean(t *testing.T) {
	i := struct {
		Data1 string `validate:"bool"`
		Data2 string `validate:"bool"`
	}{
		Data1: "true",
		Data2: "false",
	}

	validator, _ := yaml.NewValidator()
	if err := validator.Struct(i); err != nil {
		t.Errorf("expect `<nil>`, but actual `%s`", err)
	}
}

func TestValidate_StructFailedIsBoolean(t *testing.T) {
	i := struct {
		Data1 string `validate:"bool"`
	}{
		Data1: "foo",
	}

	validator, _ := yaml.NewValidator()
	if err := validator.Struct(i); err == nil {
		t.Error("expected that an error will occur, but actual not error occur")
	}
}

func TestValidate_StructSuccessIsInteger(t *testing.T) {
	i := struct {
		Data1 string `validate:"int"`
		Data2 string `validate:"int"`
		Data3 string `validate:"int"`
		Data4 string `validate:"int"`
	}{
		Data1: "-2147483648",
		Data2: "0",
		Data3: "1",
		Data4: "2147483647",
	}

	validator, _ := yaml.NewValidator()
	if err := validator.Struct(i); err != nil {
		t.Errorf("expect `<nil>`, but actual `%s`", err)
	}
}

func TestValidate_StructFailedIsInteger(t *testing.T) {
	i := struct {
		Data1 string `validate:"int"`
		Data2 string `validate:"int"`
	}{
		Data1: "-2147483649",
		Data2: "2147483648",
	}

	validator, _ := yaml.NewValidator()
	if err := validator.Struct(i); err == nil {
		t.Error("expected that an error will occur, but actual not error occur")
	}
}

func TestValidate_StructSuccessIsInteger64(t *testing.T) {
	i := struct {
		Data1 string `validate:"int64"`
		Data2 string `validate:"int64"`
		Data3 string `validate:"int64"`
		Data4 string `validate:"int64"`
	}{
		Data1: "-9223372036854775808",
		Data2: "0",
		Data3: "1",
		Data4: "9223372036854775807",
	}

	validator, _ := yaml.NewValidator()
	if err := validator.Struct(i); err != nil {
		t.Errorf("expect `<nil>`, but actual `%s`", err)
	}
}

func TestValidate_StructFailedIsInteger64(t *testing.T) {
	i := struct {
		Data1 string `validate:"int64"`
		Data2 string `validate:"int64"`
	}{
		Data1: "-9223372036854775809",
		Data2: "9223372036854775808",
	}

	validator, _ := yaml.NewValidator()
	if err := validator.Struct(i); err == nil {
		t.Error("expected that an error will occur, but actual not error occur")
	}
}

func TestValidate_StructSuccessIsUnsignedInteger16(t *testing.T) {
	i := struct {
		Data1 string `validate:"uint16"`
		Data2 string `validate:"uint16"`
		Data3 string `validate:"uint16"`
	}{
		Data1: "0",
		Data2: "1",
		Data3: "65535",
	}

	validator, _ := yaml.NewValidator()
	if err := validator.Struct(i); err != nil {
		t.Errorf("expect `<nil>`, but actual `%s`", err)
	}
}

func TestValidate_StructFailedIsUnsignedInteger16(t *testing.T) {
	i := struct {
		Data1 string `validate:"uint16"`
		Data2 string `validate:"uint16"`
	}{
		Data1: "-1",
		Data2: "65536",
	}

	validator, _ := yaml.NewValidator()
	if err := validator.Struct(i); err == nil {
		t.Error("expected that an error will occur, but actual not error occur")
	}
}

func TestValidate_StructSuccessIsUnsignedInteger64(t *testing.T) {
	i := struct {
		Data1 string `validate:"uint64"`
		Data2 string `validate:"uint64"`
		Data3 string `validate:"uint64"`
	}{
		Data1: "0",
		Data2: "1",
		Data3: "18446744073709551615",
	}

	validator, _ := yaml.NewValidator()
	if err := validator.Struct(i); err != nil {
		t.Errorf("expect `<nil>`, but actual `%s`", err)
	}
}

func TestValidate_StructFailedIsUnsignedInteger64(t *testing.T) {
	i := struct {
		Data1 string `validate:"uint64"`
		Data2 string `validate:"uint64"`
	}{
		Data1: "-1",
		Data2: "18446744073709551616",
	}

	validator, _ := yaml.NewValidator()
	if err := validator.Struct(i); err == nil {
		t.Error("expected that an error will occur, but actual not error occur")
	}
}

func TestValidate_StructSuccessIsDuration(t *testing.T) {
	i := struct {
		Data1 string `validate:"duration"`
		Data2 string `validate:"duration"`
		Data3 string `validate:"duration"`
		Data4 string `validate:"duration"`
	}{
		Data1: "1ms",
		Data2: "1s",
		Data3: "1m",
		Data4: "1h",
	}

	validator, _ := yaml.NewValidator()
	if err := validator.Struct(i); err != nil {
		t.Errorf("expect `<nil>`, but actual `%s`", err)
	}
}

func TestValidate_StructFailedIsDuration(t *testing.T) {
	i := struct {
		Data1 string `validate:"duration"`
		Data2 string `validate:"duration"`
	}{
		Data1: "1",
		Data2: "foo",
	}

	validator, _ := yaml.NewValidator()
	if err := validator.Struct(i); err == nil {
		t.Error("expected that an error will occur, but actual not error occur")
	}
}
