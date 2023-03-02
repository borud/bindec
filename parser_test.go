package spec

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	def := Parser{
		Fields: []Def{
			{TypeUint8, "uint8Var", 0, "", binary.BigEndian, false},
			{TypeInt8, "int8Var", 0, "", binary.BigEndian, false},
			{TypeUint16, "uint16Var", 0, "", binary.BigEndian, false},
			{TypeInt16, "int16Var", 0, "", binary.BigEndian, false},
			{TypeFloat32, "float32Var", 0, "", binary.BigEndian, false},
			{TypeFloat64, "float64var", 0, "", binary.BigEndian, false},
			{TypeString, "nullTermString", 0, "", binary.BigEndian, true},
		},
	}

	rec, err := def.Parse([]byte{
		0x01,
		0x02,
		0x00, 0x03,
		0x00, 0x04,
		0x40, 0x48, 0xf5, 0xc3,
		0x40, 0x9, 0x1e, 0xb8, 0x51, 0xeb, 0x85, 0x1f,
		't', 'e', 's', 't', 0})
	require.NoError(t, err)

	jsonData, err := json.MarshalIndent(rec.Values, "", "  ")
	require.NoError(t, err)

	fmt.Println(string(jsonData))

}

func TestDummy(t *testing.T) {
	buffer := bytes.Buffer{}
	binary.Write(&buffer, binary.BigEndian, float64(3.14))

	ss := make([]string, len(buffer.Bytes()))

	for i, b := range buffer.Bytes() {
		ss[i] = fmt.Sprintf("0x%x", b)
	}

	fmt.Println(strings.Join(ss, ","))

}
