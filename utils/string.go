package utils

import (
	"encoding/json"

	"github.com/iancoleman/strcase"
)

func TypeNamePublic(name string) string {
	return strcase.ToCamel(name)
}

func TypeNamePrivate(name string) string {
	return strcase.ToLowerCamel(name)
}

func JSON(ins interface{}) string {
	bs, _ := json.Marshal(ins)
	return string(bs)
}

func JSONIndent(ins interface{}) string {
	bs, _ := json.MarshalIndent(ins, "", "  ")
	return string(bs)
}
