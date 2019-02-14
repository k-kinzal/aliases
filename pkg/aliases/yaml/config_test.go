package yaml_test

import (
	"fmt"

	"github.com/k-kinzal/aliases/pkg/aliases/yaml"
)

func ptr(s string) *string {
	return &s
}

func hierarchy(path yaml.SpecPath) int {
	p := path.Parent()
	i := 1
	for p != nil {
		p = p.Parent()
		i++
	}
	return i
}

func ExampleConfigSpec_BreadthWalk() {
	config := yaml.ConfigSpec{
		"/path/to/command1": yaml.OptionSpec{
			Image: "alpine",
			Tag:   "latest",
			Name:  ptr("alpine1"),
		},
		"/path/to/command2": yaml.OptionSpec{
			Image: "alpine",
			Tag:   "latest",
			Name:  ptr("alpine2"),
			Dependencies: []yaml.DependencySpec{
				*yaml.NewDependencySpec("/path/to/command1"),
				*yaml.NewDependencySpec("/path/to/command2"),
				*yaml.NewDependencySpec(yaml.ConfigSpec{
					"/path/to/command3": yaml.OptionSpec{
						Image: "alpine",
						Tag:   "latest",
						Name:  ptr("alpine3"),
						Dependencies: []yaml.DependencySpec{
							*yaml.NewDependencySpec("/path/to/command1"),
							*yaml.NewDependencySpec("/path/to/command2"),
							*yaml.NewDependencySpec(yaml.ConfigSpec{
								"/path/to/command5": yaml.OptionSpec{
									Image: "alpine",
									Tag:   "latest",
									Name:  ptr("alpine5"),
								},
							}),
						},
					},
				}),
				*yaml.NewDependencySpec(yaml.ConfigSpec{
					"/path/to/command4": yaml.OptionSpec{
						Image: "alpine",
						Tag:   "latest",
						Name:  ptr("alpine4"),
					},
				}),
			},
		},
	}
	if err := config.BreadthWalk(func(path yaml.SpecPath, current yaml.OptionSpec) (spec *yaml.OptionSpec, e error) {
		fmt.Println(hierarchy(path)) // get number of hierarchy
		return &current, nil
	}); err != nil {
		panic(err)
	}
	// Output:
	// 1
	// 1
	// 2
	// 2
	// 3
}

func ExampleConfigSpec_DepthWalk() {
	config := yaml.ConfigSpec{
		"/path/to/command1": yaml.OptionSpec{
			Image: "alpine",
			Tag:   "latest",
			Name:  ptr("alpine2"),
			Dependencies: []yaml.DependencySpec{
				*yaml.NewDependencySpec("/path/to/command1"),
				*yaml.NewDependencySpec(yaml.ConfigSpec{
					"/path/to/command3": yaml.OptionSpec{
						Image: "alpine",
						Tag:   "latest",
						Name:  ptr("alpine3"),
						Dependencies: []yaml.DependencySpec{
							*yaml.NewDependencySpec("/path/to/command1"),
							*yaml.NewDependencySpec(yaml.ConfigSpec{
								"/path/to/command5": yaml.OptionSpec{
									Image: "alpine",
									Tag:   "latest",
									Name:  ptr("alpine5"),
								},
							}),
						},
					},
				}),
				*yaml.NewDependencySpec(yaml.ConfigSpec{
					"/path/to/command4": yaml.OptionSpec{
						Image: "alpine",
						Tag:   "latest",
						Name:  ptr("alpine4"),
					},
				}),
			},
		},
	}
	if err := config.DepthWalk(func(path yaml.SpecPath, current yaml.OptionSpec) (spec *yaml.OptionSpec, e error) {
		fmt.Println(hierarchy(path)) // get number of hierarchy
		return &current, nil
	}); err != nil {
		panic(err)
	}
	// Output:
	// 3
	// 2
	// 2
	// 1
}
