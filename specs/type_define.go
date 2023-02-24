package specs

import (
	"strings"

	"azure-spec-of-go/utils"
	"azure-spec-of-go/utils/logs"

	"github.com/go-openapi/spec"
)

// FieldType https://github.com/OAI/OpenAPI-Specification/blob/main/versions/2.0.md#data-types
type FieldType int

const (
	TypeDefault FieldType = iota
	TypeRoot

	TypeBool
	TypeInt
	TypeInt32
	TypeInt64

	TypeFloat
	TypeFloat64

	TypeString
	TypeBytes

	TypeDate
	TypeDateTime
	TypeDuration
	TypeUUID

	TypeArray
	TypeMap
	TypeFile

	TypeObject
)

// NewFieldType The value MUST be one of "string", "number", "integer", "boolean", "array" or "file"
// https://github.com/OAI/OpenAPI-Specification/blob/main/versions/2.0.md#parametertype
// https://datatracker.ietf.org/doc/html/draft-zyp-json-schema-04#section-3.5
func NewFieldType(f spec.StringOrArray, format string) (res FieldType) {
	if len(f) == 0 {
		return
	}
	if format != "" {
		if res = NewFormat(format).FieldType(); res != TypeDefault {
			return
		}
	}
	switch f[0] {
	case "string":
		return TypeString
	case "number":
		return TypeFloat
	case "integer":
		return TypeInt
	case "boolean":
		return TypeBool
	case "array":
		return TypeArray
	case "object":
		return TypeObject
	case "file":
		return TypeFile
	}
	logs.Warn("not support type: %s, format: %s", f[0], format)
	return TypeDefault
}

func (f FieldType) MockValue() interface{} {
	switch f {
	case TypeInt, TypeInt32, TypeInt64:
		return utils.MockInt()
	case TypeFloat, TypeFloat64:
		return 0.5 + float32(utils.MockInt())
	case TypeDate:
		return utils.MockDate()
	case TypeDateTime:
		return utils.MockDateTime()
	case TypeString:
		return utils.MockString()
	case TypeUUID:
		return utils.MockUUID()
	case TypeBytes:
		return utils.MockByte()
	case TypeBool:
		return true
	}
	return nil
}

type FieldFormat string

const (
	FormatDefault   FieldFormat = ""
	FormatInt32     FieldFormat = "int32"
	FormatInt64     FieldFormat = "int64"
	FormatFloat     FieldFormat = "float"
	FormatDouble    FieldFormat = "double"
	FormatByte      FieldFormat = "byte"
	FormatBinary    FieldFormat = "binary"
	FormatDate      FieldFormat = "date"
	FormatDateTime  FieldFormat = "date-time"
	FormatPassword  FieldFormat = "password"
	FormatUUID      FieldFormat = "uuid"
	FormatDuration  FieldFormat = "duration"
	FormatDecimal   FieldFormat = "decimal"
	FormatCSV       FieldFormat = "csv"
	FormatUnixTime  FieldFormat = "unixtime"
	FormatBase64URL FieldFormat = "base64url"
	FormatARMID     FieldFormat = "arm-id"
	FormatURI       FieldFormat = "uri"
	FormatFile      FieldFormat = "file"
	FormatJSON      FieldFormat = "json"
	FormatPFX       FieldFormat = "pfx"
	FormatSOLR      FieldFormat = "solr"
)

func NewFormat(f string) FieldFormat {
	if f != "" {
		f = strings.ReplaceAll(strings.ToLower(f), "_", "-")
	}
	return FieldFormat(f)
}

func (f FieldFormat) FieldType() FieldType {
	switch f {
	case FormatInt32:
		return TypeInt32
	case FormatInt64:
		return TypeInt64
	case FormatFloat:
		return TypeFloat
	case FormatDouble:
		return TypeFloat64
	case FormatByte:
		return TypeBytes
	case FormatDate:
		return TypeDate
	case FormatDateTime:
		return TypeDateTime
	case FormatDuration:
		return TypeDuration
	case FormatUUID:
		return TypeUUID
	case FormatBinary:
		return TypeBytes
	case FormatDecimal:
		return TypeFloat64
	case FormatUnixTime:

		return TypeInt64
	}
	return TypeDefault
}
