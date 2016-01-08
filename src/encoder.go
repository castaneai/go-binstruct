package binstruct

import (
	"encoding"
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
		return enc.encodeUint32(value)
	case reflect.Int64,
		reflect.Uint64:
		return enc.encodeUint64(value)
	case reflect.Float32:
		return enc.encodeFloat32(value)
	case reflect.Float64:
		return enc.encodeFloat64(value)
	case reflect.Ptr:
		return enc.encodeStruct(value)
	default:
		return fmt.Errorf("binstruct: cannot encode type: %v", value.Type().Kind())
	}
	return nil
}

func (enc *Encoder) encodePtr(value reflect.Value) error {
	switch value.Elem().Kind() {
	case reflect.Struct:
		return enc.encodeStruct(value.Elem())
	default:
		return fmt.Errorf("binstruct: cannot encode pointer to %v type", value.Elem().Kind())
	}
}

func (enc *Encoder) encodeStruct(value reflect.Value) error {
	_, has := value.Type().MethodByName("MarshalBinary")
	if !has {
		return fmt.Errorf("binstruct: type: %v has not MarhshalBinary method", value.Type().Name())
	}
	b, err := value.Interface().(encoding.BinaryMarshaler).MarshalBinary()
	fmt.Printf("marshal binary result: %v\n", b)
	if err != nil {
		return err
	}
	_, errr := enc.w.Write(b)
	return errr
}

func (enc *Encoder) encodeUint32(value reflect.Value) error {
	return binary.Write(enc.w, enc.byteOrder, uint32(value.Int()))
}

func (enc *Encoder) encodeUint64(value reflect.Value) error {
	return binary.Write(enc.w, enc.byteOrder, uint64(value.Int()))
}

func (enc *Encoder) encodeFloat32(value reflect.Value) error {
	return binary.Write(enc.w, enc.byteOrder, float32(value.Float()))
}

func (enc *Encoder) encodeFloat64(value reflect.Value) error {
	return binary.Write(enc.w, enc.byteOrder, float64(value.Float()))
}
