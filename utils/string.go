package utils

import "github.com/iancoleman/strcase"

func TypeNamePublic(name string) string {
	return strcase.ToCamel(name)
}

func TypeNamePrivate(name string) string {
	return strcase.ToLowerCamel(name)
}
