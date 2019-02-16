package config

import (
	"fmt"
	"io/ioutil"
	"os"

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

// Config is aliases configuration.
type Config struct {
	options map[string]Option
}

// add option into configuration.
func (c *Config) add(path Path, option Option) {
	c.options[path.index()] = option
}

// has returns true if an option exists.
func (c *Config) has(path Path) bool {
	_, ok := c.options[path.index()]
	return ok
}

// Get options from configuration.
func (c *Config) Get(index string) (*Option, error) {
	opt, ok := c.options[index]
	if !ok {
		return nil, fmt.Errorf("%s: index that does not exist in the config", index)
	}
	return &opt, nil
}

// Slice converts from configuration to slice.
func (c *Config) Slice() []Option {
	i := 0
	options := make([]Option, len(c.options))
	for _, opt := range c.options {
		options[i] = opt
		i++
	}
	return options
}

// newConfig creates a new Config.
func newConfig() *Config {
	return &Config{options: make(map[string]Option)}
}

// Unmarshal parses YAML-encoded data and returns configuration.
func Unmarshal(buf []byte) (*Config, error) {
	spec, err := yaml.Unmarshal(buf)
	if err != nil {
		return nil, err
	}

	config := newConfig()
	resolve := resolver(spec)
	for key, opt := range *spec {
		path := Path(key)
		if config.has(path) {
			continue
		}

		o, err := transform(resolve, yaml.SpecPath(key), opt)
		if err != nil {
			return nil, err
		}

		config.add(path, *o)
	}

	return config, nil
}

// LoadConfig load the configuration from the conifiguration file.
func LoadConfig(path string) (*Config, error) {
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
