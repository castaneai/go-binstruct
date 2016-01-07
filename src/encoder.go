package binstruct

import (
	"bytes"
	"reflect"
	"io"
)

type Encoder struct {
	byteBuf encBuffer
}

type Decoder struct {
}

func NewEncoder(w *io.Writer) *Encoder {
	enc := new(Encoder)
	return enc
}

func (enc *Encoder) Encode(e interface{}) error {
}

func (enc *Encoder) EncodeValue(value reflect.Value) error {
	if value.Kind() == reflect.Ptr {
		panic("binstruct: cannot encode pointer")
	}
}

func (enc *Encoder) encode(b *encBuffer, value reflect.Value) {
	if value.Type().Kind() == reflect.Struct {
		// TODO: encode struct
	} else {
		enc.encodeSingle(b, value)
	}
}

func (enc *Encoder) encodeSingle(b *encBuffer, value reflect.Value) {

}

func NewDecoder(b *bytes.Buffer) *Decoder {

}
