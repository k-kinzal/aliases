package yaml

import (
	"fmt"

	"github.com/k-kinzal/aliases/pkg/types"
)

// ConfigSpec is the specification of aliases config
//
// ```
// /path/to/command:
//   image: alpine
//   tag: latest
//   ...
// ```
type ConfigSpec map[string]OptionSpec

// Walk is executes the function when leaving traverse OptionSpec.
func (spec *ConfigSpec) Walk(fn func(path SpecPath, current *OptionSpec) (*OptionSpec, error)) error {
	type value struct {
		path    SpecPath
		current *OptionSpec
		parent  *OptionSpec
	}

	var hasher types.Hasher = types.SHA256
	for key, opt := range *spec {
		stack := types.NewStack(hasher)
		callstack := types.NewStack(nil)
		var v interface{} = value{SpecPath(key), &opt, nil}
		for ; v != nil; v = callstack.Pop() {
			val := v.(value)
			for i, dependency := range val.current.Dependencies {
				if dependency.IsConfig() {
					for k, c := range dependency.Config() {
						cmd := c
						callstack.Push(value{*val.path.Dependencies(i, k), &cmd, val.current})
					}
				} else {
					k := dependency.String()
					if k == key {
						continue
					}
					c, ok := (*spec)[k]
					if !ok {
						return fmt.Errorf("spec error: invalid parameter `%s` for `dependencies[%d]` is an undefined dependency in `%s`", k, i, val.path)
					}
					w := value{*val.path.Dependencies(i, k), &c, val.current}
					if stack.Has(w) {
						continue
					}
					callstack.Push(w)
				}
			}
			stack.Push(val)
		}
		for _, v := range stack.Slice() {
			val := v.(value)
			path := val.path
			current := val.current
			parent := val.parent

			current, err := fn(path, current)
			if err != nil {
				return err
			}
			if parent == nil { // is parent
				(*spec)[key] = *current
			} else {
				parent.Dependencies[path.Index()] = *NewDependencySpec(ConfigSpec{path.Base(): *current})
			}
		}
	}

	return nil
}
