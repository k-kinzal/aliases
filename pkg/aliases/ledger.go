package aliases

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/k-kinzal/aliases/pkg/types"

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
	if _, ok := ledger.schemas[index]; ok {
		return fmt.Errorf("runtime error: %s: schema is alread exists", index)
	}

	if err := defaults.Set(&schema); err != nil {
		return err
	}

	if err := ledger.validator.Struct(schema); err != nil {
		return fmt.Errorf("schema error: %s in `%s`", err, index)
	}

	schema.Path = index
	schema.FileName = path.Base(index)

	ledger.schemas[index] = schema

	return nil
}

func (ledger *Ledger) Merge(index string, schema Schema) error {
	dst := schema
	src, ok := ledger.schemas[index]
	if !ok {
		return fmt.Errorf("runtime error: %s: schema is not exists", index)
	}

	// no inherit parameters
	src.Dependencies = nil
	src.Image = ""
	src.Args = nil
	src.Tag = ""
	src.Command = nil

	if err := mergo.Map(&dst, src, mergo.WithAppendSlice); err != nil {
		return fmt.Errorf("runtime error: %s", err)
	}

	if err := defaults.Set(&dst); err != nil {
		return err
	}

	if err := ledger.validator.Struct(dst); err != nil {
		return fmt.Errorf("schema error: %s in `%s`", err, index)
	}

	dst.Path = index
	dst.FileName = path.Base(index)

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

func NewLedgerFromConfig(configpath string) (*Ledger, error) {
	if _, err := os.Stat(configpath); os.IsNotExist(err) {
		return nil, fmt.Errorf("runtime error: %s", err)
	}

	buf, err := ioutil.ReadFile(configpath)
	if err != nil {
		return nil, fmt.Errorf("runtime error: %s", err)
	}

	schemas := make(map[string]Schema)
	if err := yaml.UnmarshalStrict(buf, &schemas); err != nil {
		if e, ok := err.(*yaml.TypeError); ok {
			return nil, fmt.Errorf("yaml error: %s", strings.Replace(e.Errors[0], "in type aliases.Schema", "", 1))
		}
		return nil, err
	}

	for index, schema := range schemas {
		if schema.Path != "" {
			return nil, fmt.Errorf("yaml error: field path not found in `%s`", index)
		}
		if schema.FileName != "" {
			return nil, fmt.Errorf("yaml error: field filename not found in `%s`", index)
		}
		schema.Path = index
		schema.FileName = path.Base(index)
		schemas[index] = schema
	}

	ledger, err := NewLedger()
	if err != nil {
		return nil, err
	}

	var hasher types.Hasher = types.SHA256
	for index, schema := range schemas {
		inherits := types.NewStack(hasher)
		callstack := types.NewStack(hasher)
		callstack.Push(&struct {
			Path   string
			Schema Schema
		}{index, schema})
		for {
			var value *struct {
				Path   string
				Schema Schema
			}
			v := callstack.Pop()
			if v == nil {
				break
			}
			value = v.(*struct {
				Path   string
				Schema Schema
			})
			for idx, dependency := range value.Schema.Dependencies {
				if dependency.IsSchema() {
					for i, d := range dependency.Schemas() {
						callstack.Push(&struct {
							Path   string
							Schema Schema
						}{fmt.Sprintf("%s.Dependencies[%d].%s", value.Path, idx, i), d})
					}
					continue
				}
				if dependency.IsString() {
					i := dependency.String()
					if i == index {
						break
					}
					sch, ok := schemas[i]
					if !ok {
						return nil, fmt.Errorf("yaml error: invalid parameter `%s` for `dependencies[%d]` is an undefined dependency in `%s`", i, idx, value.Path)
					}
					if inherits.Has(sch) {
						break
					}
					callstack.Push(&struct {
						Path   string
						Schema Schema
					}{value.Schema.Path, sch})
				}
			}
			inherits.Push(value.Schema)
		}
		for i, sch := range inherits.Slice() {
			if i == 0 {
				if err := ledger.Entry(index, sch.(Schema)); err != nil {
					return nil, errors.New(strings.Replace(err.Error(), "schema error:", "yaml error:", 1))
				}
			} else {
				if err := ledger.Merge(index, sch.(Schema)); err != nil {
					return nil, errors.New(strings.Replace(err.Error(), "schema error:", "yaml error:", 1))
				}
			}
		}
	}

	return ledger, nil
}
