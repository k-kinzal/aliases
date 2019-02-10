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
	if err := spec.DepthWalk(func(path SpecPath, current OptionSpec) (*OptionSpec, error) {
		if err := v.Struct(current); err != nil {
			return nil, fmt.Errorf("yaml error: %s in `%s`", err, path)
		}
		for i, d := range current.Dependencies {
			if d.IsConfig() {
				continue
			}
			_, ok := spec[d.String()]
			if !ok {
				return nil, fmt.Errorf("yaml error: invalid parameter `%s` for `dependencies[%d]` is an undefined dependency in `%s`", d.String(), i, path)
			}
		}
		return &current, nil
	}); err != nil {
		return nil, err
	}

	return &spec, nil
}
