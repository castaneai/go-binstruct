package binstruct

import (
	"encoding/binary"
	"errors"
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
		enc.encodeUint32(value)
	case reflect.Int64,
		reflect.Uint64:
		enc.encodeUint64(value)
	case reflect.Float32:
		enc.encodeFloat32(value)
	case reflect.Float64:
		enc.encodeFloat64(value)
	default:
		return errors.New(fmt.Sprintf("go-binstruct: cannot encode type: %v", value.Type().Kind()))
	}
	return nil
}

func (enc *Encoder) encodeUint32(value reflect.Value) {
	binary.Write(enc.w, enc.byteOrder, uint32(value.Int()))
}

func (enc *Encoder) encodeUint64(value reflect.Value) {
	binary.Write(enc.w, enc.byteOrder, uint64(value.Int()))
}

func (enc *Encoder) encodeFloat32(value reflect.Value) {
	binary.Write(enc.w, enc.byteOrder, float32(value.Float()))
}

func (enc *Encoder) encodeFloat64(value reflect.Value) {
	binary.Write(enc.w, enc.byteOrder, float64(value.Float()))
}
