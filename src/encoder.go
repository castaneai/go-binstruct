package binstruct

import (
	"encoding"
	"encoding/binary"
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
	v, ok := e.(encoding.BinaryMarshaler)
	if ok {
		return enc.encodeBinaryMarshaler(v)
	}
	return enc.encode(e)
}

func (enc *Encoder) encodeBinaryMarshaler(v encoding.BinaryMarshaler) error {
	b, err := v.MarshalBinary()
	if err != nil {
		return err
	}
	_, errr := enc.w.Write(b)
	return errr
}

func (enc *Encoder) encode(e interface{}) error {
	var err error

	// BinaryMarshalerをネストしている可能性があるものは、内側まで掘る
	v := reflect.ValueOf(e)
	switch v.Type().Kind() {
	case reflect.Array, reflect.Slice:
		l := v.Len()
		for i := 0; i < l; i++ {
			err = enc.Encode(v.Index(i).Interface())
			if err != nil {
				return err
			}
		}
	case reflect.Struct:
		l := v.NumField()
		for i := 0; i < l; i++ {
			err = enc.Encode(v.Field(i).Interface())
			if err != nil {
				return err
			}
		}
	case reflect.Ptr:
		return enc.Encode(v.Elem().Interface())

	default:
		return binary.Write(enc.w, enc.byteOrder, e)
	}
	return err
}