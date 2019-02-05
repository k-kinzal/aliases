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
func (opt *Option) merge(src Option) *Option {
	dst := *opt
	src.OptionSpec.Dependencies = nil
	src.OptionSpec.Image = ""
	src.OptionSpec.Args = nil
	src.OptionSpec.Tag = ""
	src.OptionSpec.Command = nil
	if err := mergo.Map(dst.OptionSpec, src.OptionSpec, mergo.WithAppendSlice); err != nil {
		panic(err)
	}
	return &dst
}

// inherit dependencies.
func (opt *Option) inherit() *Option {
	src := Option{OptionSpec: &yaml.OptionSpec{}}
	for _, o := range opt.Dependencies {
		src = *o.merge(src)
	}
	return opt.merge(src)
}
