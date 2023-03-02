package spec

import (
	"encoding/binary"
)

// Def is the definition of a field.
type Def struct {
	// Type is the datatype of the field.
	Type Type
	// Name of the field.
	Name string
	// Whether or not this field is an array.  Note that 0 and 1 are both seen as 1, which means
	// not an array.
	Count int
	// Variable holdning the length of the field if derived from payload.  Must precede the field whose length it describes.
	LengthVar string
	// Which endianness the field has.
	Endian binary.ByteOrder
	// NullTerminated specifies that in the case of string type the field is null terminated.
	NullTerminated bool
}
