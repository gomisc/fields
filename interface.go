package fields

import (
	"fmt"
)

type FieldContainer interface {
	Bool(key string, value bool) FieldContainer
	Int(key string, value int) FieldContainer
	Int8(key string, value int8) FieldContainer
	Int16(key string, value int16) FieldContainer
	Int32(key string, value int32) FieldContainer
	Int64(key string, value int64) FieldContainer
	Uint(key string, value uint) FieldContainer
	Uint8(key string, value uint8) FieldContainer
	Uint16(key string, value uint16) FieldContainer
	Uint32(key string, value uint32) FieldContainer
	Uint64(key string, value uint64) FieldContainer
	Float32(key string, value float32) FieldContainer
	Float64(key string, value float64) FieldContainer
	Str(key string, value string) FieldContainer
	Strings(key string, values []string) FieldContainer
	Stringer(key string, value fmt.Stringer) FieldContainer
	Any(key string, value interface{}) FieldContainer
	Extract(out FieldExtractor)
}

type FieldExtractor interface {
	Bool(key string, value bool)
	Int(key string, value int)
	Int8(key string, value int8)
	Int16(key string, value int16)
	Int32(key string, value int32)
	Int64(key string, value int64)
	Uint(key string, value uint)
	Uint8(key string, value uint8)
	Uint16(key string, value uint16)
	Uint32(key string, value uint32)
	Uint64(key string, value uint64)
	Float32(key string, value float32)
	Float64(key string, value float64)
	Str(key string, value string)
	Strings(key string, values []string)
	Any(key string, value interface{})
}

type Extractor interface {
	Extract(extractor FieldExtractor)
}

type Field interface {
	Key() string
	Value() fieldValue
}

type fieldValue interface {
	Extract(key string, extractor FieldExtractor)
}
