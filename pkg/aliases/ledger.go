package aliases

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/creasty/defaults"
	"github.com/k-kinzal/aliases/pkg/aliases/validator"
	yaml "gopkg.in/yaml.v2"

	"github.com/imdario/mergo"
)

type Ledger struct {
	schemas   map[string]Schema
	validator *validator.Validate
}

func (ledger *Ledger) Entry(index string, schema Schema) error {
	dst := schema
	if src, ok := ledger.schemas[index]; ok {
		if err := mergo.Map(&dst, src, mergo.WithAppendSlice); err != nil {
			return fmt.Errorf("logic error: %s", err)
		}
	}
	if dst.Path == "" {
		dst.Path = index
		dst.FileName = path.Base(index)
	}
	if err := defaults.Set(&dst); err != nil {
		return err
	}
	if err := ledger.validator.Struct(dst); err != nil {
		return fmt.Errorf("schema error: %s in `%s`", err, index)
	}
	ledger.schemas[index] = dst

	return nil
}

func (ledger *Ledger) LookUp(index string) (*Schema, error) {
	schema, ok := ledger.schemas[index]
	if !ok {
		return nil, fmt.Errorf("runtime error: %s: schema not found", index)
	}
	return &schema, nil
}

func (ledger *Ledger) Indexes() []string {
	indexes := make([]string, 0)
	for index := range ledger.schemas {
		indexes = append(indexes, index)
	}
	return indexes
}

func (ledger *Ledger) Schemas() []Schema {
	schemas := make([]Schema, 0)
	for _, schema := range ledger.schemas {
		schemas = append(schemas, schema)
	}
	return schemas
}

func NewLedger() (*Ledger, error) {
	ledger := new(Ledger)
	ledger.schemas = make(map[string]Schema)
	v, err := validator.New()
	if err != nil {
		return nil, err
	}
	ledger.validator = v

	return ledger, nil
}

func NewLedgerFromConfig(path string) (*Ledger, error) {
	ledger, err := NewLedger()
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("runtime error: %s", err)
	}

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("runtime error: %s", err)
	}

	schemas := make(map[string]Schema)
	if err := yaml.UnmarshalStrict(buf, &schemas); err != nil {
		if e, ok := err.(*yaml.TypeError); ok {
			return nil, fmt.Errorf("yaml error: %s", strings.Replace(e.Errors[0], "in type yaml.Schema", "", 1))
		}
		return nil, err
	}
	for index, schema := range schemas {
		for index, dep := range schema.Dependencies {
			if _, ok := schemas[dep]; !ok {
				return nil, fmt.Errorf("yaml error: invalid parameter `%s` for `dependencies[%d]` is an undefined dependency in `%s`", dep, index, path)
			}

		}
		if err := ledger.Entry(index, schema); err != nil {
			return nil, err
		}
	}

	return ledger, nil
}
