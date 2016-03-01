package binstruct

import (
	"encoding"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
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
	return fmt.Errorf("not implemented")
}

func (dec *Decoder) decodeBinaryUnmarshaler(v encoding.BinaryUnmarshaler) error {
	data, err := ioutil.ReadAll(dec.r)
	if err != nil {
		return err
	}
	return v.UnmarshalBinary(data)
}
