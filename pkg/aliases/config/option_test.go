package config_test

import (
	"encoding/json"
	"fmt"

	"github.com/k-kinzal/aliases/pkg/aliases/config"
	"github.com/k-kinzal/aliases/pkg/aliases/yaml"
)

func ptr(s string) *string {
	return &s
}

func ExampleOption_Binary() {
	opt := config.Option{
		OptionSpec: &yaml.OptionSpec{
			Docker: struct {
				Image string `yaml:"image" default:"docker"`
				Tag   string `yaml:"tag" default:"18.09.0"`
			}{
				Image: "docker",
				Tag:   "18.09.0",
			},
			Image: "alpine",
			Tag:   "latest",
			Name:  ptr("alpine1"),
		},
		Namespace:    "/",
		Path:         "/path/to/command1",
		FileName:     "command1",
		Dependencies: nil,
	}
	out, err := json.Marshal(opt.Binary("/tmp"))
	if err != nil {
		panic(nil)
	}
	fmt.Println(string(out))
	// Output: {"Image":"docker","Tag":"18.09.0","Path":"/tmp/docker-18-09-0"}
}

func ExampleOption_Binaries() {
	opt := config.Option{
		OptionSpec: &yaml.OptionSpec{
			Docker: struct {
				Image string `yaml:"image" default:"docker"`
				Tag   string `yaml:"tag" default:"18.09.0"`
			}{
				Image: "docker",
				Tag:   "18.09.1",
			},
			Image: "alpine",
			Tag:   "latest",
			Name:  ptr("alpine1"),
		},
		Namespace: "/",
		Path:      "/path/to/command1",
		FileName:  "command1",
		Dependencies: []*config.Option{
			&config.Option{
				OptionSpec: &yaml.OptionSpec{
					Docker: struct {
						Image string `yaml:"image" default:"docker"`
						Tag   string `yaml:"tag" default:"18.09.0"`
					}{
						Image: "docker",
						Tag:   "18.06.1",
					},
					Image: "alpine",
					Tag:   "latest",
					Name:  ptr("alpine2"),
				},
				Namespace: "/",
				Path:      "/path/to/command2",
				FileName:  "command2",
			},
			&config.Option{
				OptionSpec: &yaml.OptionSpec{
					Docker: struct {
						Image string `yaml:"image" default:"docker"`
						Tag   string `yaml:"tag" default:"18.09.0"`
					}{
						Image: "docker",
						Tag:   "18.09.0",
					},
					Image: "alpine",
					Tag:   "latest",
					Name:  ptr("alpine3"),
				},
				Namespace: "/",
				Path:      "/path/to/command3",
				FileName:  "command3",
				Dependencies: []*config.Option{
					&config.Option{
						OptionSpec: &yaml.OptionSpec{
							Docker: struct {
								Image string `yaml:"image" default:"docker"`
								Tag   string `yaml:"tag" default:"18.09.0"`
							}{
								Image: "docker",
								Tag:   "18.06.1",
							},
							Image: "alpine",
							Tag:   "latest",
							Name:  ptr("alpine4"),
						},
						Namespace: "/",
						Path:      "/path/to/command4",
						FileName:  "command4",
					},
				},
			},
		},
	}
	for _, binary := range opt.Binaries("/tmp") {
		out, err := json.Marshal(binary)
		if err != nil {
			panic(nil)
		}
		fmt.Println(string(out))
	}
	// Output:
	// {"Image":"docker","Tag":"18.06.1","Path":"/tmp/docker-18-06-1"}
	// {"Image":"docker","Tag":"18.09.1","Path":"/tmp/docker-18-09-1"}
	// {"Image":"docker","Tag":"18.09.0","Path":"/tmp/docker-18-09-0"}
}
