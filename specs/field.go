package specs

import (
	"bytes"

	"azure-spec-of-go/utils/logs"

	"github.com/go-openapi/spec"
)

type Definition struct {
	Name    string
	JSONTag string
	RefBy   []*Field // this definition is referenced by
	Fields  map[string]*Field

	// remember has mock values from
	mockFrom map[*Field]struct{}
}

func (d *Definition) tryMockFor(f *Field) bool {
	if _, ok := d.mockFrom[f]; ok {
		return false
	}
	if d.mockFrom == nil {
		d.mockFrom = map[*Field]struct{}{}
	}
	d.mockFrom[f] = struct{}{}
	return true
}

func (d *Definition) MockValue(bs *bytes.Buffer, from *Field) {
	if !d.tryMockFor(from) {
		bs.WriteString("null")
		return
	}
	bs.WriteString("{\n")
	first := true
	for name, f := range d.Fields {
		if !first {
			bs.WriteString(",\n")
		}
		first = false
		bs.WriteString(sf(`"%s":`, name))
		f.MockValue(bs)
	}
	bs.WriteString("\n}")
}

// Field definition to a go struct field
type Field struct {
	Type        FieldType   `json:"type,omitempty"`
	OriginType  []string    `json:"origin_type,omitempty"`
	Format      FieldFormat `json:"format,omitempty"`
	Name        string      `json:"name,omitempty"`
	JSONTag     string      `json:"json_tag,omitempty"` // origin name defined in swagger
	Description string      `json:"description,omitempty"`
	Default     interface{} `json:"default,omitempty"`
	Enums       []string    `json:"enums,omitempty"`    // may be enum type, values in enum
	RefName     string      `json:"ref_name,omitempty"` // this property refers name, expand to Subs
	RefTo       *Definition // refer to a definition
	ArrayItem   *Field      `json:"array_item,omitempty"` // if type is array, this field for item detail

	// Subs properties of object type, expand by spec, it will stop when cycle ref exists
	Subs     map[string]*Field `json:"subs,omitempty"`
	subNames []string

	AdditionalPropertiesType FieldType `json:"additional_properties_type,omitempty"` // string, or object
}

func (f *Field) AddSub(f2 *Field) {
	f.Subs[f2.Name] = f2
	f.subNames = append(f.subNames, f2.Name)
}

func NewFieldFromSchema(name string, sch *spec.Schema) *Field {
	if sch == nil {
		return nil
	}
	var f Field
	props := sch.SchemaProps
	f.OriginType = props.Type
	f.Type = NewFieldType(props.Type, props.Format)
	if add := sch.AdditionalProperties; add != nil && add.Schema != nil {
		f.AdditionalPropertiesType = NewFieldType(add.Schema.Type, add.Schema.Format)
	}

	f.Format = NewFormat(props.Format)
	f.Default = sch.Default
	f.Description = sch.Description
	f.Name = name
	f.JSONTag = name

	if len(sch.Enum) > 0 {
		if sch.Type[0] != "string" {
			logs.Warn("not supported enum type: %s", sch.Type[0])
		} else {
			for _, val := range sch.Enum {
				f.Enums = append(f.Enums, val.(string))
			}
		}
	}

	if sch.Ref.String() != "" {
		f.RefName = sch.Ref.String()
	}

	if sch.Items != nil {
		item := sch.Items.Schema
		if item == nil || len(sch.Items.Schemas) > 0 {
			item = &sch.Items.Schemas[0]
		}
		f.ArrayItem = NewFieldFromSchema("", item)
	}

	// process properties
	// for circular reference, loaders.Expands will stop there so sch.Properties will be nil
	f.Subs = map[string]*Field{}
	for name, sch := range sch.Properties {
		subField := NewFieldFromSchema(name, &sch)
		f.AddSub(subField)
	}

	// expanded AllOf, AnyOf, OneOf
	for _, schema := range append(append(sch.AllOf, sch.OneOf...), sch.AllOf...) {
		for name, sch := range schema.Properties {
			subField := NewFieldFromSchema(name, &sch)
			f.AddSub(subField)
		}
	}

	return &f
}

func appendSimpleValue(bs *bytes.Buffer, val interface{}) {
	if s, ok := val.(string); ok {
		bs.WriteString(sf(`"%s"`, s))
	} else {
		bs.WriteString(sf("%v", val))
	}
}

// MockValue try to mock this field to a string builder
func (f *Field) MockValue(bs *bytes.Buffer) {
	if val := f.Type.MockValue(); val != nil {
		appendSimpleValue(bs, val)
		return
	}
	if f.RefTo != nil {
		f.RefTo.MockValue(bs, f)
		return
	}

	if len(f.Subs) > 0 {
		bs.WriteString("{\n")
		var first = true
		for name, sub := range f.Subs {
			if !first {
				bs.WriteString(",\n")
			}
			first = false
			bs.WriteString(sf(`"%s"`, name))
			sub.MockValue(bs)
		}
		bs.WriteString("\n}")
	} else if f.ArrayItem != nil {
		bs.WriteString("[")
		f.ArrayItem.MockValue(bs)
		bs.WriteString("]")
	} else if f.AdditionalPropertiesType > 0 {
		bs.WriteString("{\n\"key\":")
		appendSimpleValue(bs, f.AdditionalPropertiesType.MockValue())
		bs.WriteString("\n}")
	} else if f.RefName != "" {
		bs.WriteString("null")
	} else {
		// not support for now?
		logs.Warn("not support to mock value for: %s", f.Name)
		bs.WriteString(`""`)
	}
}
