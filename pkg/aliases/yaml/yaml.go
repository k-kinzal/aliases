package yaml

import (
	"fmt"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

func Unmarshal(buf []byte) (*ConfigSpec, error) {
	spec := ConfigSpec{}
	if err := yaml.UnmarshalStrict(buf, &spec); err != nil {
		if e, ok := err.(*yaml.TypeError); ok {
			message := e.Errors[0]
			message = strings.Replace(message, " into yaml.DependencySpec", "", 1)
			message = strings.Replace(message, " into yaml.OptionSpec", "", 1)
			message = strings.Replace(message, " into yaml.ConfigSpec", "", 1)
			message = strings.Replace(message, " in type yaml.DependencySpec", "", 1)
			message = strings.Replace(message, " in type yaml.OptionSpec", "", 1)
			message = strings.Replace(message, " in type yaml.ConfigSpec", "", 1)
			return nil, fmt.Errorf("yaml error: %s", message)
		}
		return nil, err
	}
	v, err := NewValidator()
	if err != nil {
		return nil, err
	}

	if err := spec.Walk(func(path SpecPath, current *OptionSpec) (spec *OptionSpec, e error) {
		if err := v.Struct(*current); err != nil {
			return nil, fmt.Errorf("yaml error: %s in `%s`", err, path)
		}
		return current, nil
	}); err != nil {
		message := err.Error()
		message = strings.Replace(message, "spec error", "yaml error", 1)
		return nil, fmt.Errorf(message)
	}

	return &spec, nil
}
