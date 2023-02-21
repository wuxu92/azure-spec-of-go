package specs

import (
	"fmt"
	"strings"

	"azure-spec-of-go/utils"

	"github.com/go-openapi/spec"
)

// define the model,mapping to the swagger definitions

type Definition struct {
}

type Spec struct {
	spec *spec.Swagger
}

// RenderDefinitions render all definitions as output
func (s *Spec) RenderDefinitions() []byte {

	return nil
}

func renderOneDefinition(name string, def *spec.Schema) {
	var bs strings.Builder
	wl := func(format string, args ...interface{}) {
		bs.WriteString(fmt.Sprintf(format, args...))
		bs.WriteByte('\n')
	}
	typeName := utils.TypeNamePublic(name)
	wl("type %s struct", typeName)
	for name, prop := range def.Properties {
		if prop.SchemaProps.Ref.String() != "" {
			wl("%s *%s", name, prop.SchemaProps.Ref.String())
			continue
		}
	}
}
