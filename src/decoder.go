package binstruct

import (
  "encoding"
	"encoding/binary"
	"io"
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
    return dec.
  }
}

func (dec *Decoder) decodeBinaryUnmarshaler(v encoding.BinaryUnmarshaler) error {
  // TODO: io.Reader だときつい・・・
}
