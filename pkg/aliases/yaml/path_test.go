package yaml_test

import (
	"fmt"
	"testing"

	"github.com/k-kinzal/aliases/pkg/aliases/yaml"
)

func ExampleSpecPath_Name() {
	var path yaml.SpecPath = "/path/to/command1.dependencies[0]./path/to/command2"
	fmt.Println(path.Name())
	// Output: /path/to/command2
}

func ExampleSpecPath_Name2() { // is path of root OptionSpec
	var path yaml.SpecPath = "/path/to/command1"
	fmt.Println(path.Name())
	// Output: /path/to/command1
}

func TestSpecPath_NameInvalidPath(t *testing.T) {
	var path yaml.SpecPath = "."
	if name := path.Name(); name != "" {
		t.Errorf("expect `\"\"`, but actual `%#v`", name)
	}
}

func ExampleSpecPath_Base() {
	var path yaml.SpecPath = "/path/to/command1.dependencies[0]./path/to/command2"
	fmt.Println(path.Base())
	// Output: command2
}
func ExampleSpecPath_Dependencies() {
	var path yaml.SpecPath = "/path/to/command1"
	fmt.Println(path.Dependencies(0, "/path/to/command2"))
	// Output: /path/to/command1.dependencies[0]./path/to/command2
}

func TestSpecPath_DependenciesInvalidPath(t *testing.T) {
	var path yaml.SpecPath = "."
	if dep := path.Dependencies(0, "/path/to/command2"); dep != nil {
		t.Errorf("expect `<nil>`, but actual `%#v`", dep)
	}
}

func ExampleSpecPath_Index() {
	var path yaml.SpecPath = "/path/to/command1.dependencies[0]./path/to/command2"
	fmt.Println(path.Index())
	// Output: 0
}

func TestSpecPath_IndexInvalidPath(t *testing.T) {
	var path yaml.SpecPath = "."
	if index := path.Index(); index != -1 {
		t.Errorf("expect `-1`, but actual `%#v`", index)
	}
}

func ExampleSpecPath_Parent() {
	var path yaml.SpecPath = "/path/to/command1.dependencies[0]./path/to/command2"
	fmt.Println(path.Parent())
	// Output: /path/to/command1
}

func TestSpecPath_Parent(t *testing.T) {
	var path yaml.SpecPath = "/path/to/command1"
	if p := path.Parent(); p != nil {
		t.Errorf("expect `<nil>`, but actual `%#v`", *p)
	}
}

func TestSpecPath_ParentInvalidPath(t *testing.T) {
	var path yaml.SpecPath = "."
	if p := path.Parent(); p != nil {
		t.Errorf("expect `<nil>`, but actual `%#v`", *p)
	}
}
