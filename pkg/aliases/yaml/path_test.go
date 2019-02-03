package yaml_test

import (
	"fmt"
	"testing"

	"github.com/k-kinzal/aliases/pkg/aliases/yaml"
)

func ExampleSpecPath_Base() {
	var path yaml.SpecPath = "/path/to/command1.dependencies[0]./path/to/command2"
	fmt.Println(path.Base())
	// Output: /path/to/command2
}

func ExampleSpecPath_Base2() { // is path of root OptionSpec
	var path yaml.SpecPath = "/path/to/command1"
	fmt.Println(path.Base())
	// Output: /path/to/command1
}

func TestSpecPath_BaseInvalidPath(t *testing.T) {
	var path yaml.SpecPath = "."
	if base := path.Base(); base != "" {
		t.Errorf("expect `\"\"`, but actual `%#v`", base)
	}
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
