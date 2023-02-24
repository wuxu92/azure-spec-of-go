package specs

import (
	"bytes"
	"fmt"
	"strings"

	"azure-spec-of-go/utils"
	"azure-spec-of-go/utils/logs"

	"github.com/go-openapi/spec"
)

// define the model,mapping to the swagger definitions

type Spec struct {
	spec        *spec.Swagger
	expandSpec  *spec.Swagger
	Definitions map[string]*Definition
}

func NewSpec(swag *spec.Swagger) *Spec {
	ins := &Spec{
		spec: swag,
	}
	if err := ins.ParseDefinitions(); err != nil {
		logs.Error("parse definition for %s err: %+v", swag.BasePath, err)
		return nil
	}
	return ins
}

func (s *Spec) ParseDefinitions() (err error) {
	if s.Definitions == nil {
		// pass 1: generate all definitions
		s.Definitions = make(map[string]*Definition, len(s.spec.Definitions))
		for name, specDef := range s.spec.Definitions {
			var def Definition
			def.Name = utils.TypeNamePublic(name)
			def.JSONTag = name
			def.Fields = make(map[string]*Field, len(specDef.Properties))
			fields := NewFieldFromSchema("", &specDef)
			def.Fields = fields.Subs
			s.Definitions[name] = &def
		}

		// pass 2, refer parse
		for _, def := range s.Definitions {
			if err = s.deRef(def); err != nil {
				return err
			}
		}
	}
	return nil
}

// de-reference for a field for one level
func (s *Spec) deRef(def *Definition) (err error) {
	for _, sub := range def.Fields {
		if sub.RefName == "" {
			continue
		}
		tokens := strings.Split(sub.RefName, "definitions/")
		if len(tokens) < 2 {
			logs.Warn("not support this kind of ref: %s", sub.RefName)
			continue
		}
		if refDef, ok := s.Definitions[tokens[1]]; ok {
			refDef.RefBy = append(refDef.RefBy, sub)
			sub.RefTo = refDef
		} else {
			return fmt.Errorf("no definition for ref %s", sub.RefName)
		}
	}
	return nil
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
		def.MockValue(&bs, nil)
	}
	bs.WriteString("}\n")
	return bs.Bytes()
}
