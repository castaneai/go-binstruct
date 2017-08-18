package binstruct

import (
	"encoding"
	"encoding/binary"
	"io"
	"io/ioutil"
	"reflect"
	"fmt"
)

type Decoder struct {
	r         io.Reader
	byteOrder binary.ByteOrder
}

func NewDecoder(r io.Reader, byteOrder binary.ByteOrder) *Decoder {
	return &Decoder{
		r:         r,
		byteOrder: byteOrder,
	}
}

func (dec *Decoder) Decode(e interface{}) error {
	v, ok := e.(encoding.BinaryUnmarshaler)
	if ok {
		return dec.decodeBinaryUnmarshaler(v)
	}
	return dec.decode(e)
}

func (dec *Decoder) decodeBinaryUnmarshaler(v encoding.BinaryUnmarshaler) error {
	data, err := ioutil.ReadAll(dec.r)
	if err != nil {
		return err
	}
	return v.UnmarshalBinary(data)
}

func (dec *Decoder) decode(e interface{}) error {
	var err error

	// BinaryUnMarshalerをネストしている可能性があるものは、内側まで掘る
	v := reflect.ValueOf(e)

	if v.Type().Kind() != reflect.Ptr {
		return fmt.Errorf("decode value must be a pointer")
	}

	vp := v.Elem()
	switch vp.Type().Kind() {
	case reflect.Array, reflect.Slice:
		l := vp.Len()
		for i := 0; i < l; i++ {
			err = dec.Decode(vp.Index(i).Addr().Interface())
			if err != nil {
				return err
			}
		}
	case reflect.Struct:
		l := vp.NumField()
		for i := 0; i < l; i++ {
			err = dec.Decode(vp.Field(i).Addr().Interface())
			if err != nil {
				return err
			}
		}

	default:
		return binary.Read(dec.r, dec.byteOrder, e)
	}
	return err
}
