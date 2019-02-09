package config

import (
	"github.com/imdario/mergo"
	"github.com/k-kinzal/aliases/pkg/aliases/yaml"
)

// Option is configuration option.
type Option struct {
	*yaml.OptionSpec
	Namespace    string
	Path         string
	FileName     string
	Dependencies []*Option
}

// merge Option and Option
func (opt *Option) merge(source Option) *Option {
	dst := *opt.OptionSpec
	src := *source.OptionSpec
	src.Dependencies = nil
	src.Image = ""
	src.Args = nil
	src.Tag = ""
	src.Command = nil
	if err := mergo.Map(&dst, src, mergo.WithAppendSlice); err != nil {
		panic(err)
	}

	return &Option{
		OptionSpec:   &dst,
		Namespace:    opt.Namespace,
		Path:         opt.Path,
		FileName:     opt.FileName,
		Dependencies: opt.Dependencies,
	}
}

// inherit dependencies.
func (opt *Option) inherit() *Option {
	src := Option{OptionSpec: &yaml.OptionSpec{}}
	for _, o := range opt.Dependencies {
		src = *o.merge(src)
	}
	return opt.merge(src)
}
