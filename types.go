package fields

// fields types
type (
	b        bool
	i        int
	i8       int8
	i16      int16
	i32      int32
	i64      int64
	u        uint
	u8       uint8
	u16      uint16
	u32      uint32
	u64      uint64
	f32      float32
	f64      float64
	str      string
	strs     []string
	anything struct {
		val interface{}
	}
)

func (v b) Extract(key string, extractor FieldExtractor) {
	extractor.Bool(key, bool(v))
}

func (v i) Extract(key string, extractor FieldExtractor) {
	extractor.Int(key, int(v))
}

func (v i8) Extract(key string, extractor FieldExtractor) {
	extractor.Int8(key, int8(v))
}

func (v i16) Extract(key string, extractor FieldExtractor) {
	extractor.Int16(key, int16(v))
}

func (v i32) Extract(key string, extractor FieldExtractor) {
	extractor.Int32(key, int32(v))
}

func (v i64) Extract(key string, extractor FieldExtractor) {
	extractor.Int64(key, int64(v))
}

func (v u) Extract(key string, extractor FieldExtractor) {
	extractor.Uint(key, uint(v))
}

func (v u8) Extract(key string, extractor FieldExtractor) {
	extractor.Uint8(key, uint8(v))
}

func (v u16) Extract(key string, extractor FieldExtractor) {
	extractor.Uint16(key, uint16(v))
}

func (v u32) Extract(key string, extractor FieldExtractor) {
	extractor.Uint32(key, uint32(v))
}

func (v u64) Extract(key string, extractor FieldExtractor) {
	extractor.Uint64(key, uint64(v))
}

func (v f32) Extract(key string, extractor FieldExtractor) {
	extractor.Float32(key, float32(v))
}

func (v f64) Extract(key string, extractor FieldExtractor) {
	extractor.Float64(key, float64(v))
}

func (v str) Extract(key string, extractor FieldExtractor) {
	extractor.Str(key, string(v))
}

func (v strs) Extract(key string, extractor FieldExtractor) {
	extractor.Strings(key, v)
}

func (v anything) Extract(key string, extractor FieldExtractor) {
	extractor.Any(key, v.val)
}
