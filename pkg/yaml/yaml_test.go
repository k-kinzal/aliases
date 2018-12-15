package yaml_test

import (
	"github.com/k-kinzal/aliases/pkg/docker"
	"github.com/k-kinzal/aliases/pkg/yaml"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func toKebabCase(str string) string {
	kebab := matchFirstCap.ReplaceAllString(str, "${1}-${2}")
	kebab = matchAllCap.ReplaceAllString(kebab, "${1}-${2}")
	return strings.ToLower(kebab)
}

func TestUnmarshalConfFile(t *testing.T) {
	content := `---
/usr/local/bin/kubectl:
  image: chatwork/kubectl
  tag: 1.11.2
`

	defs, err := yaml.UnmarshalConfFile([]byte(content))
	if err != nil {
		t.Errorf("unmarshal configuration error: %v", err)
	}
	if len(defs) != 1 {
		t.Errorf("expected `1`, but in actual the unmarshaled configuration length is `%d`", len(defs))
	}
	if _, ok := defs["/usr/local/bin/kubectl"]; !ok {
		t.Error("/usr/local/bin/kubectl does not exist in unmarshaled configuration")
	}
	if defs["/usr/local/bin/kubectl"].Image != "chatwork/kubectl" {
		t.Errorf("expected `chatwork/kubectl`, but in actual `%s` has been set in image", defs["/usr/local/bin/kubectl"].Image)
	}
	if defs["/usr/local/bin/kubectl"].Tag != "1.11.2" {
		t.Errorf("expected `1.11.2`, but in actual `%s` has been set in tag", defs["/usr/local/bin/kubectl"].Tag)
	}
}

func TestUnmarshalConfFile_ShouldBeKebabCaseOfFieldName(t *testing.T) {
	val := reflect.New(reflect.TypeOf(yaml.YamlDefinition{})).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)

		tag, ok := field.Tag.Lookup("yaml")
		if !ok {
			t.Errorf("tag of yaml is undefined in the field of %s", field.Name)
		}
		arr := strings.Split(tag, ",")
		name := arr[0]

		if name != toKebabCase(field.Name) {
			t.Errorf("expected yaml key name of %s is %s, but %s is defined", field.Name, toKebabCase(field.Name), name)
		}
	}
}

func TestUnmarshalConfFile_ShouldBeSameFieldAsDockerRunOptsExist(t *testing.T) {
	val1 := reflect.New(reflect.TypeOf(docker.RunOpts{})).Elem()
	val2 := reflect.New(reflect.TypeOf(yaml.YamlDefinition{})).Elem()
	for i := 0; i < val1.NumField(); i++ {
		field := val1.Type().Field(i)
		if _, ok := val2.Type().FieldByName(field.Name); !ok {
			t.Errorf("field in %s does not exist in YamlDefinition", field.Name)
		}
	}
}
