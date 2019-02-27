package yaml_test

import (
	"fmt"
	"testing"

	"github.com/k-kinzal/aliases/pkg/aliases/yaml"
)

func ExampleUnmarshal() {
	content := `
/path/to/command1:
  image: alpine
  tag: latest
  name: alpine1
/path/to/command2:
  image: alpine
  tag: latest
  name: alpine2
`
	config, err := yaml.Unmarshal([]byte(content))
	if err != nil {
		panic(err)
	}
	fmt.Println(*(*config)["/path/to/command1"].Name)
	fmt.Println(*(*config)["/path/to/command2"].Name)
	// Output:
	// alpine1
	// alpine2
}

func TestUnmarshal_UndefinedDependencyReference(t *testing.T) {
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

func TestUnmarshalFailedEmpty(t *testing.T) {
	content := ``
	_, err := yaml.Unmarshal([]byte(content))
	if err != nil {
		t.Errorf("expect `<nil>`, but actual `%s`", err)
		return
	}
}

func TestUnmarshalFailedNotYamlFormat(t *testing.T) {
	actual := "yaml error: line 2: cannot unmarshal !!str `foo:bar...`"
	content := `
foo:bar:baz
`
	spec, err := yaml.Unmarshal([]byte(content))
	if spec != nil {
		t.Errorf("expect `<nil>`, but actual `%#v`", *spec)
		return
	}
	if err.Error() != actual {
		t.Errorf("expect `%s`, but actual `%s`", actual, err)
		return
	}
}

func TestUnmarshalFailedUndefinedProperty(t *testing.T) {
	actual := "yaml error: line 3: field foo not found"
	content := `
/path/to/command:
  foo: bar
`
	spec, err := yaml.Unmarshal([]byte(content))
	if spec != nil {
		t.Errorf("expect `<nil>`, but actual `%#v`", *spec)
		return
	}
	if err.Error() != actual {
		t.Errorf("expect `%s`, but actual `%s`", actual, err)
		return
	}
}

func TestUnmarshalFailedUnmatchedPropertyType(t *testing.T) {
	actual := "yaml error: line 3: cannot unmarshal !!str `STDERR` into []string"
	content := `
/path/to/command:
  attach: STDERR
`
	spec, err := yaml.Unmarshal([]byte(content))
	if spec != nil {
		t.Errorf("expect `<nil>`, but actual `%#v`", *spec)
		return
	}
	if err.Error() != actual {
		t.Errorf("expect `%s`, but actual `%s`", actual, err)
		return
	}
}

func TestUnmarshalFailedUnmatchedDependencyType(t *testing.T) {
	actual := "yaml error: cannot unmarshal !![]interface {} `[1 2 3]` into string or map"
	content := `
/path/to/command:
  image: alpine
  tag: latest
  dependencies:
  - [1, 2, 3]
`
	spec, err := yaml.Unmarshal([]byte(content))
	if spec != nil {
		t.Errorf("expect `<nil>`, but actual `%#v`", *spec)
		return
	}
	if err.Error() != actual {
		t.Errorf("expect `%s`, but actual `%s`", actual, err)
		return
	}
}

func TestUnmarshalFailedNotExistsReference(t *testing.T) {
	actual := "yaml error: invalid parameter `/path/to/other-command` for `dependencies[0]` is an undefined dependency in `/path/to/command`"
	content := `
/path/to/command:
  image: alpine
  tag: latest
  dependencies:
  - /path/to/other-command
`
	spec, err := yaml.Unmarshal([]byte(content))
	if spec != nil {
		t.Errorf("expect `<nil>`, but actual `%#v`", *spec)
		return
	}
	if err.Error() != actual {
		t.Errorf("expect `%s`, but actual `%s`", actual, err)
		return
	}
}

func TestUnmarshalFailedValidation(t *testing.T) {
	actual := "yaml error: invalid parameter for `tag` is required in `/path/to/command`"
	content := `
/path/to/command:
  image: alpine
`
	spec, err := yaml.Unmarshal([]byte(content))
	if spec != nil {
		t.Errorf("expect `<nil>`, but actual `%#v`", *spec)
		return
	}
	if err.Error() != actual {
		t.Errorf("expect `%s`, but actual `%s`", actual, err)
		return
	}
}

func TestUnmarshalFailedValidationEnum(t *testing.T) {
	actual := "yaml error: invalid parameter `foo` for `attach[0]` is one of STDIN, STDOUT, STDERR in `/path/to/command`"
	content := `
/path/to/command:
  image: alpine
  tag: latest
  attach:
  - foo
`
	spec, err := yaml.Unmarshal([]byte(content))
	if spec != nil {
		t.Errorf("expect `<nil>`, but actual `%#v`", *spec)
		return
	}
	if err.Error() != actual {
		t.Errorf("expect `%s`, but actual `%s`", actual, err)
		return
	}
}

func TestUnmarshalFailedValidationMin(t *testing.T) {
	actual := "yaml error: invalid parameter `1` for `blkioWeight` is a number greater than or equal to `10` in `/path/to/command`"
	content := `
/path/to/command:
  image: alpine
  tag: latest
  blkioWeight: 1
`
	spec, err := yaml.Unmarshal([]byte(content))
	if spec != nil {
		t.Errorf("expect `<nil>`, but actual `%#v`", *spec)
		return
	}
	if err.Error() != actual {
		t.Errorf("expect `%s`, but actual `%s`", actual, err)
		return
	}
}

func TestUnmarshalFailedValidationMax(t *testing.T) {
	actual := "yaml error: invalid parameter `1001` for `blkioWeight` is a number less than or equal to `1000` in `/path/to/command`"
	content := `
/path/to/command:
  image: alpine
  tag: latest
  blkioWeight: 1001
`
	spec, err := yaml.Unmarshal([]byte(content))
	if spec != nil {
		t.Errorf("expect `<nil>`, but actual `%#v`", *spec)
		return
	}
	if err.Error() != actual {
		t.Errorf("expect `%s`, but actual `%s`", actual, err)
		return
	}
}

func TestUnmarshalFailedValidationIPv4(t *testing.T) {
	actual := "yaml error: invalid parameter `foo` for `ip` is not IPv4 format (e.g., 172.30.100.104) in `/path/to/command`"
	content := `
/path/to/command:
  image: alpine
  tag: latest
  ip: foo
`
	spec, err := yaml.Unmarshal([]byte(content))
	if spec != nil {
		t.Errorf("expect `<nil>`, but actual `%#v`", *spec)
		return
	}
	if err.Error() != actual {
		t.Errorf("expect `%s`, but actual `%s`", actual, err)
		return
	}
}

func TestUnmarshalFailedValidationIPv6(t *testing.T) {
	actual := "yaml error: invalid parameter `foo` for `ip6` is not IPv6 format (e.g., 2001:db8::33) in `/path/to/command`"
	content := `
/path/to/command:
  image: alpine
  tag: latest
  ip6: foo
`
	spec, err := yaml.Unmarshal([]byte(content))
	if spec != nil {
		t.Errorf("expect `<nil>`, but actual `%#v`", *spec)
		return
	}
	if err.Error() != actual {
		t.Errorf("expect `%s`, but actual `%s`", actual, err)
		return
	}
}

func TestUnmarshalFailedValidationMacAddress(t *testing.T) {
	actual := "yaml error: invalid parameter `foo` for `macAddress` is not mac address format in `/path/to/command`"
	content := `
/path/to/command:
  image: alpine
  tag: latest
  macAddress: foo
`
	spec, err := yaml.Unmarshal([]byte(content))
	if spec != nil {
		t.Errorf("expect `<nil>`, but actual `%#v`", *spec)
		return
	}
	if err.Error() != actual {
		t.Errorf("expect `%s`, but actual `%s`", actual, err)
		return
	}
}
