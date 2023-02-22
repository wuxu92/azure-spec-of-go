package specs

import (
	"strings"

	"azure-spec-of-go/utils/logs"

	"github.com/go-openapi/spec"
)

type Definition struct {
	Name    string
	JSONTag string
	Fields  map[string]*Field
}

func (d *Definition) MockValue(bs *strings.Builder) {
	bs.WriteString("{\n")
	first := true
	for name, f := range d.Fields {
		if !first {
			bs.WriteString(",\n")
		}
		first = true
		bs.WriteString(sf(`"%s"`, name))
		f.MockJSON(bs)
	}
	bs.WriteString("}\n")
}

// Field definition to a go struct field
type Field struct {
	Type        FieldType   `json:"type,omitempty"`
	OriginType  []string    `json:"origin_type",omitempty`
	Format      FieldFormat `json:"format,omitempty"`
	Name        string      `json:"name,omitempty"`
	JSONTag     string      `json:"json_tag,omitempty"` // origin name defined in swagger
	Description string      `json:"description,omitempty"`
	Default     interface{} `json:"default,omitempty"`
	Enums       []string    `json:"enums,omitempty"`    // may be enum type, values in enum
	RefName     string      `json:"ref_name,omitempty"` // this property refers name, expand to Subs

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
	if name == "" {
		f.Type = TypeRoot
	}

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

// MockJSON try to mock this field to a string builder
func (f *Field) MockJSON(bs *strings.Builder) {
	bs.WriteString(f.JSONTag)
	bs.WriteByte(':')
	if val := f.Type.MockValue(); val != nil {
		_, isStr := val.(string)
		if isStr {
			bs.WriteByte('"')
		}
		bs.WriteString(sf("%v", val))
		if isStr {
			bs.WriteByte('"')
		}
	}
	if len(f.Subs) > 0 {
		bs.WriteByte('{')
		var first = true
		for name, sub := range f.Subs {
			if !first {
				bs.WriteString(",\n")
			}
			first = false
			bs.WriteString(sf(`"%s"`, name))
			sub.MockJSON(bs)
		}
		bs.WriteByte('}')
		bs.WriteByte('\n')
	}
}
