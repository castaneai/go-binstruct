package binstruct

import (
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
)

type Encoder struct {
	w         io.Writer
	byteOrder binary.ByteOrder
}

func NewEncoder(w io.Writer, byteOrder binary.ByteOrder) *Encoder {
	return &Encoder{
		w:         w,
		byteOrder: byteOrder,
	}
}

func (enc *Encoder) Encode(e interface{}) error {
	return enc.EncodeValue(reflect.ValueOf(e))
}

func (enc *Encoder) EncodeValue(value reflect.Value) error {
	switch value.Type().Kind() {
	case reflect.Int,
		reflect.Int32,
		reflect.Uint,
		reflect.Uint32:
		enc.encodeInt32(value)
	default:
		panic(fmt.Sprintf("go-binstruct: cannot encode type: %v", value.Type().Kind()))
	}
	return nil
}

func (enc *Encoder) encodeInt32(value reflect.Value) {
	b := make([]byte, 4, 4)
	enc.byteOrder.PutUint32(b, uint32(value.Int()))
	enc.w.Write(b)
}
