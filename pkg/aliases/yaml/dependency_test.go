package yaml_test

import (
	"fmt"
	"testing"

	"github.com/k-kinzal/aliases/pkg/aliases/yaml"

	y "gopkg.in/yaml.v2"
)

func ExampleNewDependencySpec() {
	dep := yaml.NewDependencySpec("/path/to/command")

	fmt.Println(dep.IsString())
	// Output: true
}

func TestNewDependencySpecPassConfigMap(t *testing.T) {
	dep := yaml.NewDependencySpec(yaml.ConfigSpec{})

	fmt.Println(dep.IsConfig())
	// Output: true
}

func TestNewDependencySpecPassInteger(t *testing.T) {
	defer func() {
		err := recover()
		if err != "value is expected to be type `string` or `map`, but the actual is `int`" {
			t.Errorf("not expect message of \"%v\"", err)
		}
	}()
	yaml.NewDependencySpec(1)
	t.Error("expected that `panic()` but did not occur")
}

func ExampleDependencySpec_IsConfig() {
	dep := yaml.NewDependencySpec(yaml.ConfigSpec{})

	fmt.Println(dep.IsConfig())
	// Output: true
}

func ExampleDependencySpec_IsString() {
	dep := yaml.NewDependencySpec("/path/to/command")

	fmt.Println(dep.IsString())
	// Output: true
}

func ExampleDependencySpec_Config() {
	config1 := yaml.ConfigSpec{
		"/path/to/command": yaml.OptionSpec{
			Image: "alpine",
			Tag:   "latest",
		},
	}

	dep := yaml.NewDependencySpec(config1)

	for index, option := range dep.Config() {
		fmt.Println(index)
		fmt.Println(option.Image)
		fmt.Println(option.Tag)
	}

	fmt.Println()
	// Output:
	// /path/to/command
	// alpine
	// latest
}

func ExampleDependencySpec_String() {
	dep := yaml.NewDependencySpec("/path/to/command")

	fmt.Println(dep.String())
	// Output: /path/to/command
}

func ExampleDependencySpec_UnmarshalYAML() {
	content := `
- /path/to/command1
- /path/to/command2:
    images: alpine
    tag: latest
    dependencies:
    - /path/to/command3
    - /path/to/command4:
        images: alpine
        tag: latest
`
	var dependencies []yaml.DependencySpec
	if err := y.Unmarshal([]byte(content), &dependencies); err != nil {
		panic(err)
	}

	fmt.Println(dependencies[0].IsString())
	fmt.Println(dependencies[1].IsConfig())
	for _, option := range dependencies[1].Config() {
		fmt.Println(option.Dependencies[0].IsString())
		fmt.Println(option.Dependencies[1].IsConfig())
	}
	// Output:
	// true
	// true
	// true
	// true
}
