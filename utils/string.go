package utils

import (
	"bytes"
	"encoding/json"

	"azure-spec-of-go/utils/logs"

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

func JSONFormat(bs []byte, indent bool) (res []byte) {
	var buf bytes.Buffer
	if indent {
		if err := json.Indent(&buf, bs, "", "  "); err != nil {
			logs.Error("json indent: %+v", err)
		}
	} else {
		if err := json.Compact(&buf, bs); err != nil {
			logs.Error("json compact: %+v", err)
		}
	}
	if buf.Len() == 0 {
		return bs
	}
	return buf.Bytes()
}
