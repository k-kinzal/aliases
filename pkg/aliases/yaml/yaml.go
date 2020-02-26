package yaml

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/k-kinzal/aliases/pkg/types"

	"github.com/creasty/defaults"

	"github.com/imdario/mergo"

	yaml "gopkg.in/yaml.v2"
)

type Option struct {
	*OptionSpec
	Path         SpecPath
	Dependencies []*Option
}

func (option *Option) Namespace() string {
	bases := make([]string, 0)
	for p := option.Path.Parent(); p != nil; p = p.Parent() {
		bases = append(bases, types.MD5(p.Base()))
	}
	return fmt.Sprintf("/%s", strings.Join(bases, "/"))
}

// index return configuration index.
func (option *Option) Index() string {
	ns := option.Namespace()
	name := option.Path.Name()
	return strings.Replace(fmt.Sprintf("%s%s", ns, name), "//", "/", -1)
}

type Config map[string]*Option

func (config *Config) Set(path SpecPath, option Option) {
	if path.Parent() == nil {
		(*config)[path.Name()] = &option
	} else {
		opt := config.Get(*path.Parent())
		if opt == nil {
			return
		}
		opt.Dependencies[path.Index()] = &option
	}
}

func (config *Config) Get(path SpecPath) *Option {
	var parts []SpecPath
	for p := &path; p != nil; p = p.Parent() {
		parts = append(parts, *p)
	}
	if len(parts) == 0 {
		return nil
	}
	opt := (*config)[parts[len(parts)-1].Name()]
	for i := len(parts) - 2; i >= 0; i-- {
		opt = opt.Dependencies[parts[i].Index()]
	}
	return opt
}

// Unmarshal parses YAML-encoded data and returns config specification.
func Unmarshal(buf []byte) (*Config, error) {
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
			return nil, Error(message)
		}
		return nil, Error(err)
	}

	conf := Config{}
	// initialize configuration
	if err := spec.BreadthWalk(func(path SpecPath, current OptionSpec) (*OptionSpec, error) {
		conf.Set(path, Option{&current, path, make([]*Option, len(current.Dependencies))})
		return &current, nil
	}); err != nil {
		return nil, err
	}
	// resolve dependencies
	if err := spec.DepthWalk(func(path SpecPath, current OptionSpec) (*OptionSpec, error) {
		option := conf.Get(path)
		var dependencies []*Option
		for index, dependency := range current.Dependencies {
			if dependency.IsConfig() {
				for key := range dependency.Config() {
					dependencies = append(dependencies, conf.Get(*path.Dependencies(index, key)))
				}
			} else {
				key := dependency.String()
				opt := conf.Get(SpecPath(key))
				if opt == nil {
					return nil, Errorf("invalid parameter `%s` for `dependencies[%d]` is an undefined dependency in `%s`", key, index, path)
				}
				dependencies = append(dependencies, opt)
			}
		}
		option.Dependencies = dependencies

		return &current, nil
	}); err != nil {
		return nil, err
	}
	// first inherit parameter
	if err := spec.DepthWalk(func(path SpecPath, current OptionSpec) (*OptionSpec, error) {
		option := conf.Get(path)
		dst := *option.OptionSpec
		for i := len(option.Dependencies) - 1; i >= 0; i-- { // merge what is defined later
			src := *option.Dependencies[i].OptionSpec
			src.Dependencies = nil
			src.Image = ""
			src.Args = nil
			src.Tag = ""
			src.Entrypoint = nil
			src.Command = nil
			if err := mergo.Map(&dst, src, mergo.WithAppendSlice); err != nil {
				panic(err)
			}
		}
		option.OptionSpec = &dst
		conf.Set(path, *option)

		return &current, nil
	}); err != nil {
		return nil, err
	}
	// second inherit parameter
	// if it is circulating, it will not be the correct parameter unless inherited twice
	//     cmd1:
	//       ...
	//       dependencies:
	//       - cmd2
	//     cmd2:
	//       ...
	//       dependencies:
	//       - cmd1
	if err := spec.DepthWalk(func(path SpecPath, current OptionSpec) (*OptionSpec, error) {
		option := conf.Get(path)
		dst := *option.OptionSpec
		for i := len(option.Dependencies) - 1; i >= 0; i-- {
			src := *option.Dependencies[i].OptionSpec
			src.Dependencies = nil
			src.Image = ""
			src.Args = nil
			src.Tag = ""
			src.Entrypoint = nil
			src.Command = nil
			if err := mergo.Map(&dst, src, mergo.WithAppendSlice); err != nil {
				panic(err)
			}
		}
		option.OptionSpec = &dst
		conf.Set(path, *option)

		return &current, nil
	}); err != nil {
		return nil, err
	}
	// validate & set default parameter
	v := NewValidator()
	if err := spec.DepthWalk(func(path SpecPath, current OptionSpec) (*OptionSpec, error) {
		option := conf.Get(path)
		if err := v.Struct(option.OptionSpec); err != nil {
			return nil, Errorf("%s in `%s`", err, path)
		}
		if len(option.Dependencies) > 0 && option.Docker == nil {
			dockerSpec := &DockerSpec{}
			if err := defaults.Set(dockerSpec); err != nil {
				panic(err)
			}
			option.Docker = dockerSpec
		}
		if err := defaults.Set(option.OptionSpec); err != nil {
			panic(err)
		}
		dependencies := make([]*Option, len(option.Dependencies))
		for index, dependency := range option.Dependencies {
			dependencies[index] = conf.Get(dependency.Path)
		}
		option.Dependencies = dependencies
		conf.Set(path, *option)

		return &current, nil
	}); err != nil {
		return nil, err
	}

	return &conf, nil
}

func LoadFile(path string) (*Config, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("%s: no such file or directory", path)
	}

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config, err := Unmarshal(buf)
	if err != nil {
		return nil, err
	}

	return config, nil
}
