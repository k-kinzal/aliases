package yaml

import (
	"fmt"
	"reflect"

	"github.com/k-kinzal/aliases/pkg/types"
)

// DependencySpec is a dependency on ConfigSpec or a string of references to other options.
//
// ```yaml
// dependencies:
// - command1
// - command2:
//     ...
// ```
type DependencySpec struct {
	types.Union
}

// IsString returns true if the value is a string.
func (spec *DependencySpec) IsString() bool {
	return spec.Type() == spec.Left()
}

// IsConfig returns true if the value is a ConfigSpec.
func (spec *DependencySpec) IsConfig() bool {
	return spec.Type() == spec.Right()
}

// String is return string value.
func (spec *DependencySpec) String() string {
	v := spec.Value()
	return v.(string)
}

// Config is return ConfigSpec value.
func (spec *DependencySpec) Config() ConfigSpec {
	v := spec.Value()
	return v.(ConfigSpec)
}

func (spec *DependencySpec) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v interface{}
	err := unmarshal(&v)
	if err != nil {
		return err
	}

	switch raw := v.(type) {
	case string:
		*spec = *NewDependencySpec(raw)
	case map[interface{}]interface{}:
		var v ConfigSpec
		err := unmarshal(&v)
		if err != nil {
			return err
		}
		*spec = *NewDependencySpec(v)
	default:
		return fmt.Errorf("cannot unmarshal !!%T `%v` into string or map", v, v)
	}

	return nil
}

// NewDependencySpec creates a new DependencySpec.
func NewDependencySpec(value interface{}) *DependencySpec {
	return &DependencySpec{*types.NewUnion(reflect.String, reflect.TypeOf(ConfigSpec{}).Kind(), value)}
}
