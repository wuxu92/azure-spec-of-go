package specs

import (
	"bytes"

	"azure-spec-of-go/utils"

	"github.com/go-openapi/spec"
)

// define the model,mapping to the swagger definitions

type Spec struct {
	spec        *spec.Swagger
	Definitions map[string]Definition
}

func NewSpec(swag *spec.Swagger) *Spec {
	ins := &Spec{
		spec: swag,
	}
	ins.ParseDefinitions()
	return ins
}

func (s *Spec) ParseDefinitions() {
	if s.Definitions == nil {
		s.Definitions = make(map[string]Definition, len(s.spec.Definitions))
		for name, specDef := range s.spec.Definitions {
			var def Definition
			def.Name = utils.TypeNamePublic(name)
			def.JSONTag = name
			def.Fields = make(map[string]*Field, len(specDef.Properties))
			fields := NewFieldFromSchema("", &specDef)
			def.Fields = fields.Subs
			s.Definitions[name] = def
		}
	}
}

// RenderDefinitions render all definitions as output
func (s *Spec) RenderDefinitions() []byte {
	var bs bytes.Buffer
	bs.WriteString("{\n")
	var first = true
	for name, def := range s.Definitions {
		if !first {
			bs.WriteString(",\n")
		}
		first = false
		bs.WriteString(sf(`"%s": `, name))
		def.MockValue(&bs)
	}
	bs.WriteString("}\n")
	return bs.Bytes()
}
