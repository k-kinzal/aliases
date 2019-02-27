package yaml_test

import (
	"fmt"

	"github.com/k-kinzal/aliases/pkg/aliases/yaml"
)

func ExampleError() {
	err := yaml.Error("parse error")
	fmt.Println(err)
	// Output: yaml error: parse error
}

func ExampleErrorf() {
	reason := "undefined key"
	err := yaml.Errorf("parse error: %s", reason)
	fmt.Println(err)
	// Output: yaml error: parse error: undefined key
}
