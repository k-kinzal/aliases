package yaml_test

import (
	"fmt"
	"testing"

	"github.com/k-kinzal/aliases/pkg/aliases/yaml"
)

func ExampleConfigSpec_BreadthWalk() {
	content := `
/path/to/command1:
  image: alpine
  tag: latest
  name: alpine1
/path/to/command2:
  image: alpine
  tag: latest
  name: alpine2
  dependencies:
  - /path/to/command1
  - /path/to/command2
  - /path/to/command3:
      image: alpine
      tag: latest
      name: alpine3
      dependencies:
      - /path/to/command1
      - /path/to/command2
      - /path/to/command5:
          image: alpine
          tag: latest
          name: alpine5
  - /path/to/command4:
      image: alpine
      tag: latest
      name: alpine4
`
	config, err := yaml.Unmarshal([]byte(content))
	if err != nil {
		panic(err)
	}
	if err := config.BreadthWalk(func(path yaml.SpecPath, current yaml.OptionSpec) (spec *yaml.OptionSpec, e error) {
		fmt.Println(path, *current.Name)
		return &current, nil
	}); err != nil {
		panic(err)
	}
	// Output:
	// /path/to/command1 alpine1
	// /path/to/command2 alpine2
	// /path/to/command2.dependencies[2]./path/to/command3 alpine3
	// /path/to/command2.dependencies[3]./path/to/command4 alpine4
	// /path/to/command2.dependencies[2]./path/to/command3.dependencies[2]./path/to/command5 alpine5
}

func ExampleConfigSpec_Walk() {
	content := `
/path/to/command1:
  image: alpine
  tag: latest
  name: alpine1
/path/to/command2:
  image: alpine
  tag: latest
  name: alpine2
  dependencies:
  - /path/to/command1
  - /path/to/command2
  - /path/to/command3:
      image: alpine
      tag: latest
      name: alpine3
      dependencies:
      - /path/to/command1
      - /path/to/command2
      - /path/to/command5:
          image: alpine
          tag: latest
          name: alpine5
  - /path/to/command4:
      image: alpine
      tag: latest
      name: alpine4
`
	config, err := yaml.Unmarshal([]byte(content))
	if err != nil {
		panic(err)
	}
	if err := config.DepthWalk(func(path yaml.SpecPath, current yaml.OptionSpec) (spec *yaml.OptionSpec, e error) {
		fmt.Println(path, *current.Name)
		return &current, nil
	}); err != nil {
		panic(err)
	}
	// Output:
	// /path/to/command1 alpine1
	// /path/to/command2.dependencies[2]./path/to/command3.dependencies[2]./path/to/command5 alpine5
	// /path/to/command2.dependencies[2]./path/to/command3 alpine3
	// /path/to/command2.dependencies[3]./path/to/command4 alpine4
	// /path/to/command2 alpine2
}

func TestConfigSpec_WalkUndefinedDependencyReference(t *testing.T) {
	content := `
/path/to/command1:
  image: alpine
  tag: latest
  name: alpine1
/path/to/command2:
  image: alpine
  tag: latest
  name: alpine2
  dependencies:
  - /path/to/command3
`
	_, err := yaml.Unmarshal([]byte(content))
	if err == nil || err.Error() != "yaml error: invalid parameter `/path/to/command3` for `dependencies[0]` is an undefined dependency in `/path/to/command2`" {
		t.Errorf("not expect message of \"%v\"", err)
	}
}
