package yaml

import (
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

// BreadthWalk is executes the function when entry traverse OptionSpec.
func (spec *ConfigSpec) BreadthWalk(fn func(path SpecPath, current OptionSpec) (*OptionSpec, error)) error {
	type value struct {
		path    SpecPath
		current *OptionSpec
		parent  *OptionSpec
	}
	queue := make([]value, 0)
	for key, opt := range *spec {
		o := opt
		queue = append(queue, value{SpecPath(key), &o, nil})
	}
	for i := 0; i < len(queue); i++ {
		v := queue[i]

		current, err := fn(v.path, *v.current)
		if err != nil {
			return err
		}
		if v.parent == nil {
			(*spec)[v.path.String()] = *current
		} else {
			conf := v.parent.Dependencies[v.path.Index()].Config()
			conf[v.path.Name()] = *current
			v.parent.Dependencies[v.path.Index()] = *NewDependencySpec(conf)
		}

		for index, dependency := range v.current.Dependencies {
			if dependency.IsString() {
				continue
			}
			for key, opt := range dependency.Config() {
				o := opt
				queue = append(queue, value{*v.path.Dependencies(index, key), &o, v.current})
			}
		}
	}

	return nil
}

// Walk is executes the function when leaving traverse OptionSpec.
func (spec *ConfigSpec) DepthWalk(fn func(path SpecPath, current OptionSpec) (*OptionSpec, error)) error {
	type value struct {
		path    SpecPath
		current *OptionSpec
		parent  *OptionSpec
	}

	for key, opt := range *spec {
		stack := types.NewStack(nil)
		callstack := types.NewStack(nil)
		callstack.Push(value{SpecPath(key), &opt, nil})
		for val := callstack.Pop(); val != nil; val = callstack.Pop() {
			v := val.(value)
			for i, dependency := range v.current.Dependencies {
				if dependency.IsString() {
					continue
				}
				for k, c := range dependency.Config() {
					cmd := c
					callstack.Push(value{*v.path.Dependencies(i, k), &cmd, v.current})
				}
			}
			stack.Push(val)
		}
		for _, val := range stack.Slice() {
			v := val.(value)
			current, err := fn(v.path, *v.current)
			if err != nil {
				return err
			}
			if v.parent == nil {
				(*spec)[v.path.String()] = *current
			} else {
				conf := v.parent.Dependencies[v.path.Index()].Config()
				conf[v.path.Name()] = *current
				v.parent.Dependencies[v.path.Index()] = *NewDependencySpec(conf)
			}
		}
	}

	return nil
}
