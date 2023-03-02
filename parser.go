package spec

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

type Parser struct {
	Fields []Def
}

var (
	ErrUnhandledType = errors.New("unhandled type")
	ErrReading       = errors.New("read error")
)

func (d *Parser) Parse(data []byte) (*Record, error) {
	values := map[string]any{}

	r := bytes.NewBuffer(data)

	for _, field := range d.Fields {
		switch field.Type {
		case TypeUint8:
			vv := uint8(0)
			err := binary.Read(r, field.Endian, &vv)
			if err != nil {
				return nil, err
			}
			values[field.Name] = vv

		case TypeInt8:
			vv := int8(0)
			err := binary.Read(r, field.Endian, &vv)
			if err != nil {
				return nil, err
			}
			values[field.Name] = vv

		case TypeUint16:
			vv := uint16(0)
			err := binary.Read(r, field.Endian, &vv)
			if err != nil {
				return nil, err
			}
			values[field.Name] = vv

		case TypeInt16:
			vv := int16(0)
			err := binary.Read(r, field.Endian, &vv)
			if err != nil {
				return nil, err
			}
			values[field.Name] = vv

		case TypeFloat32:
			vv := float32(0)
			err := binary.Read(r, field.Endian, &vv)
			if err != nil {
				return nil, err
			}
			values[field.Name] = vv

		case TypeFloat64:
			vv := float64(0)
			err := binary.Read(r, field.Endian, &vv)
			if err != nil {
				return nil, err
			}
			values[field.Name] = vv

		default:
			return nil, fmt.Errorf("%w: %v", ErrUnhandledType, field.Type)
		}
	}
	return &Record{d.Fields, values}, nil
}
