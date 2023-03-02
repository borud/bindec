package spec

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	testInt8Fields = []Def{
		{TypeUint8, "uint8Var", 0, "", binary.BigEndian, false},
		{TypeInt8, "int8Var", 0, "", binary.BigEndian, false},
	}

	testSingleFields = []Def{
		{TypeUint8, "uint8Var", 0, "", binary.BigEndian, false},
		{TypeInt8, "int8Var", 0, "", binary.BigEndian, false},
		{TypeUint16, "uint16Var", 0, "", binary.BigEndian, false},
		{TypeInt16, "int16Var", 0, "", binary.BigEndian, false},
		{TypeFloat32, "float32Var", 0, "", binary.BigEndian, false},
		{TypeFloat64, "float64var", 0, "", binary.BigEndian, false},
	}

	testArrayFields = []Def{
		{TypeInt8, "int8Arr", 3, "", binary.BigEndian, false},
		{TypeUint8, "uint8Arr", 3, "", binary.BigEndian, false},
		{TypeUint16, "uint16Arr", 3, "", binary.BigEndian, false},
	}

	testAllFields = append(testSingleFields, testArrayFields...)

	testSingleData = []byte{
		0x01,
		0x02,
		0x00, 0x03,
		0x00, 0x04,
		0x40, 0x48, 0xf5, 0xc3,
		0x40, 0x9, 0x1e, 0xb8, 0x51, 0xeb, 0x85, 0x1f,
	}

	testArrayData = []byte{
		0xa, 0xb, 0xc,
		0xa, 0xb, 0xc,
		0x0, 0xa, 0x0, 0xb, 0x0, 0xc,
	}

	testAllData = append(testSingleData, testArrayData...)
)

func TestParse(t *testing.T) {
	def := Parser{Fields: testSingleFields}

	rec, err := def.Parse(testSingleData)
	require.NoError(t, err)

	jsonData, err := json.MarshalIndent(rec.Values, "", "  ")
	require.NoError(t, err)
	fmt.Println(string(jsonData))
}

func BenchmarkParser(b *testing.B) {
	def := Parser{Fields: testSingleFields}

	for i := 0; i < b.N; i++ {
		def.Parse(testSingleData)
	}
}
