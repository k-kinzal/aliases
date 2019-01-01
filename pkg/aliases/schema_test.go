package aliases_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/k-kinzal/aliases/pkg/aliases"
)

func TestNewSchema_ShouldbeLowerCamelCaseFieldname(t *testing.T) {
	val := reflect.New(reflect.TypeOf(aliases.Schema{})).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)

		tag, ok := field.Tag.Lookup("yaml")
		if !ok {
			continue
		}
		arr := strings.Split(tag, ",")
		name := arr[0]

		if strings.ToLower(name) != strings.ToLower(field.Name) {
			t.Errorf("expected yaml key name of %s is %s, but %s is defined", field.Name, field.Name, name)
		}
	}
}
