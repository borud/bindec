package spec

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"reflect"
)

type Parser struct {
	Fields []Def
}

var (
	ErrUnhandledType            = errors.New("unhandled type")
	ErrReading                  = errors.New("read error")
	ErrEOFBeforeNullTermination = errors.New("reached EOF before null termination")
	ErrLengthVarNotDefined      = errors.New("length variable not defined")
	ErrMissingStringParameters  = errors.New("missing LengthVar, Count or NullTerminated parameter")
)

func (d *Parser) Parse(data []byte) (*Record, error) {
	values := map[string]any{}
	reader := bytes.NewBuffer(data)

	for _, field := range d.Fields {

		// Handle strings
		if field.Type == TypeString {
			vv, err := readString(reader, field, values)
			if err != nil {
				return nil, err
			}
			values[field.Name] = vv
			continue
		}

		// Handle single values
		if field.Count < 2 {

			vv, err := readSingle(reader, field)
			if err != nil {
				return nil, err
			}
			log.Printf("single: %+v", vv)
			values[field.Name] = vv
			continue
		}

		// Handle array of values
		vv, err := readArray(reader, field)
		if err != nil {
			return nil, err
		}
		values[field.Name] = vv

	}
	return &Record{d.Fields, values}, nil
}

func readString(reader io.Reader, field Def, values map[string]any) (string, error) {
	// Handle case where we get the length from another, previously defined value
	if field.LengthVar != "" {
		aa, ok := values[field.LengthVar]
		if !ok {
			return "", fmt.Errorf("%w: %s", ErrLengthVarNotDefined, field.LengthVar)
		}

		buffer := make([]byte, reflect.ValueOf(aa).Int())
		return string(buffer), binary.Read(reader, field.Endian, buffer)
	}

	// Handle case where we get the length from field.Count
	if field.Count != 0 {
		buffer := make([]byte, field.Count)
		return string(buffer), binary.Read(reader, field.Endian, buffer)
	}

	// Handle case where we have null terminated string
	if field.NullTerminated {
		return readNullTerminatedString(reader)
	}

	return "", ErrMissingStringParameters
}

func readNullTerminatedString(reader io.Reader) (string, error) {
	out := bytes.Buffer{}
	var b [1]byte
	for {
		_, err := reader.Read(b[:])
		if err == io.EOF {
			return out.String(), ErrEOFBeforeNullTermination
		}

		if b[0] == 0 {
			return out.String(), nil
		}

		out.Write(b[:])
	}
}

func readSingle(reader io.Reader, field Def) (any, error) {
	switch field.Type {
	case TypeUint8:
		var vv uint8
		err := binary.Read(reader, field.Endian, &vv)
		if err != nil {
			return nil, err
		}
		return vv, nil

	case TypeInt8:
		var vv int8
		err := binary.Read(reader, field.Endian, &vv)
		if err != nil {
			return nil, err
		}
		return vv, nil

	case TypeUint16:
		var vv uint16
		err := binary.Read(reader, field.Endian, &vv)
		if err != nil {
			return nil, err
		}
		return vv, nil

	case TypeInt16:
		var vv int16
		err := binary.Read(reader, field.Endian, &vv)
		if err != nil {
			return nil, err
		}
		return vv, nil

	case TypeUint32:
		var vv uint32
		err := binary.Read(reader, field.Endian, &vv)
		if err != nil {
			return nil, err
		}
		return vv, nil

	case TypeInt32:
		var vv int32
		err := binary.Read(reader, field.Endian, &vv)
		if err != nil {
			return nil, err
		}
		return vv, nil

	case TypeUint64:
		var vv uint64
		err := binary.Read(reader, field.Endian, &vv)
		if err != nil {
			return nil, err
		}
		return vv, nil

	case TypeInt64:
		var vv int64
		err := binary.Read(reader, field.Endian, &vv)
		if err != nil {
			return nil, err
		}
		return vv, nil

	case TypeFloat32:
		var vv float32
		err := binary.Read(reader, field.Endian, &vv)
		if err != nil {
			return nil, err
		}
		return vv, nil

	case TypeFloat64:
		var vv float64
		err := binary.Read(reader, field.Endian, &vv)
		if err != nil {
			return nil, err
		}
		return vv, nil

	default:
		return nil, ErrUnhandledType
	}
}

func readArray(reader io.Reader, field Def) (any, error) {
	switch field.Type {
	case TypeUint8:
		var vv = make([]uint8, field.Count)
		return vv, binary.Read(reader, field.Endian, vv)

	case TypeInt8:
		var vv = make([]int8, field.Count)
		return vv, binary.Read(reader, field.Endian, vv)

	case TypeUint16:
		var vv = make([]uint16, field.Count)
		return vv, binary.Read(reader, field.Endian, vv)

	case TypeInt16:
		var vv = make([]int16, field.Count)
		return vv, binary.Read(reader, field.Endian, vv)

	case TypeUint32:
		var vv = make([]uint32, field.Count)
		return vv, binary.Read(reader, field.Endian, vv)

	case TypeInt32:
		var vv = make([]int32, field.Count)
		return vv, binary.Read(reader, field.Endian, vv)

	case TypeUint64:
		var vv = make([]uint64, field.Count)
		return vv, binary.Read(reader, field.Endian, vv)

	case TypeInt64:
		var vv = make([]int64, field.Count)
		return vv, binary.Read(reader, field.Endian, vv)

	case TypeFloat32:
		var vv = make([]float32, field.Count)
		return vv, binary.Read(reader, field.Endian, vv)

	case TypeFloat64:
		var vv = make([]float64, field.Count)
		return vv, binary.Read(reader, field.Endian, vv)

	default:
		return nil, ErrUnhandledType
	}
}
