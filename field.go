package fields

import (
	"fmt"
)

type field struct {
	val fieldValue
	key string
}

func (f *field) Key() string {
	return f.key
}

func (f *field) Value() fieldValue {
	return f.val
}

func Bool(key string, val bool) Field {
	return &field{key: key, val: b(val)}
}

func Int(key string, val int) Field {
	return &field{key: key, val: i(val)}
}

func Int8(key string, val int8) Field {
	return &field{key: key, val: i8(val)}
}

func Int16(key string, val int16) Field {
	return &field{key: key, val: i16(val)}
}

func Int32(key string, val int32) Field {
	return &field{key: key, val: i32(val)}
}

func Int64(key string, val int64) Field {
	return &field{key: key, val: i64(val)}
}

func Uint(key string, val uint) Field {
	return &field{key: key, val: u(val)}
}

func Uint8(key string, val uint8) Field {
	return &field{key: key, val: u8(val)}
}

func Uint16(key string, val uint16) Field {
	return &field{key: key, val: u16(val)}
}

func Uint32(key string, val uint32) Field {
	return &field{key: key, val: u32(val)}
}

func Uint64(key string, val uint64) Field {
	return &field{key: key, val: u64(val)}
}

func Float32(key string, val float32) Field {
	return &field{key: key, val: f32(val)}
}

func Float64(key string, val float64) Field {
	return &field{key: key, val: f64(val)}
}

func Str(key string, val string) Field {
	return &field{key: key, val: str(val)}
}

func Strings(key string, val []string) Field {
	return &field{key: key, val: strs(val)}
}

func Stringer(key string, val fmt.Stringer) Field {
	return &field{key: key, val: str(val.String())}
}

func Any(key string, val interface{}) Field {
	return &field{key: key, val: anything{val: val}}
}

func Error(err error) Field {
	return &field{key: "error", val: str(err.Error())}
}
