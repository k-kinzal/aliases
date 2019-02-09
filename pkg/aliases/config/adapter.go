package config

import (
	"fmt"
	"strings"

	"github.com/k-kinzal/aliases/pkg/types"

	"github.com/k-kinzal/aliases/pkg/aliases/yaml"
)

// Path is an adapter handled by config from yaml.SpecPath.
type Path yaml.SpecPath

// namespace returns configuration namespace.
func (path *Path) namespace() string {
	bases := make([]string, 0)
	for p := (*yaml.SpecPath)(path).Parent(); p != nil; p = p.Parent() {
		bases = append(bases, types.MD5(p.Base()))
	}
	return fmt.Sprintf("/%s", strings.Join(bases, "/"))
}

// index return configuration index.
func (path *Path) index() string {
	ns := path.namespace()
	name := (*yaml.SpecPath)(path).Name()
	return strings.Replace(fmt.Sprintf("%s%s", ns, name), "//", "/", -1)
}

// resolver resolves reference string to Option.
func resolver(spec *yaml.ConfigSpec) func(key string) (*Option, error) {
	return func(key string) (*Option, error) {
		opt, ok := (*spec)[key]
		if !ok {
			return nil, fmt.Errorf("")
		}
		o, err := transform(resolver(spec), yaml.SpecPath(key), opt)
		if err != nil {
			return nil, err
		}
		return o, nil
	}
}

// transform yaml.OptionSpec to config.Option.
func transform(resolve func(key string) (*Option, error), path yaml.SpecPath, current yaml.OptionSpec) (*Option, error) {
	relatives := make([]*Option, 0)
	for i, dependency := range current.Dependencies {
		if dependency.IsConfig() {
			for k, o := range dependency.Config() {
				opt, err := transform(resolve, *path.Dependencies(i, k), o)
				if err != nil {
					panic(err)
				}
				relatives = append(relatives, opt)
			}
		} else {
			opt, err := resolve(dependency.String())
			if err != nil {
				return nil, err
			}
			relatives = append(relatives, opt)
		}
	}

	option := &Option{OptionSpec: &current}
	option.Namespace = (*Path)(&path).namespace()
	option.Path = path.Name()
	option.FileName = path.Base()
	option.Dependencies = relatives

	return option, nil
}
