package spec

import "reflect"

type Type byte

const (
	TypeUnspecified Type = iota
	TypeUint8
	TypeInt8
	TypeUint16
	TypeInt16
	TypeUint32
	TypeInt32
	TypeUint64
	TypeInt64
	TypeFloat32
	TypeFloat64
	TypeString
)

func (t Type) ReflectType() reflect.Type {
	return reflect.TypeOf(map[Type]any{
		TypeUint8:   uint8(0),
		TypeInt8:    int8(0),
		TypeUint16:  uint16(0),
		TypeInt16:   int16(0),
		TypeUint32:  uint32(0),
		TypeInt32:   int32(0),
		TypeUint64:  uint64(0),
		TypeInt64:   int64(0),
		TypeFloat32: float32(0.0),
		TypeFloat64: float64(0.0),
		TypeString:  "",
	}[t])
}

func (t Type) String() string {
	return [...]string{
		"unspecified",
		"uint8",
		"int8",
		"uint16",
		"int16",
		"uint32",
		"int32",
		"uint64",
		"int64",
		"float32",
		"float64",
		"string"}[t]
}

func TypeFromString(s string) Type {
	return map[string]Type{
		"unspecified": TypeUnspecified,
		"uint8":       TypeUint8,
		"int8":        TypeInt8,
		"uint16":      TypeUint16,
		"int16":       TypeInt16,
		"uint32":      TypeUint32,
		"int32":       TypeInt32,
		"uint64":      TypeUint64,
		"int64":       TypeInt64,
		"float32":     TypeFloat32,
		"float64":     TypeFloat64,
		"string":      TypeString,
	}[s]
}
