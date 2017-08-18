package binstruct

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestDecodeBytesWithTerminator(t *testing.T) {
	src := []byte("\x68\x65\x6C\x6C\x6F\x00")
	r := bytes.NewReader(src)

	dec := NewDecoder(r, binary.LittleEndian)
	actual := BytesWithTerminator{Terminator: 0x00}
	err := dec.Decode(&actual)
	if err != nil {
		t.Errorf("%v", err)
	}

	expected := []byte("hello")
	AssertDeepEquals(actual.Bytes, expected, t)
}

func TestDecodeNestedStruct(t *testing.T) {
	type InnerStruct struct {
		UInt32Value uint32
	}
	type OuterStruct struct {
		UInt16Value uint16
		InnerStructValue InnerStruct
	}

	src := []byte("\xFF\xFF\xB1\x7F\x39\x05")
	r := bytes.NewReader(src)
	dec := NewDecoder(r, binary.LittleEndian)

	actual := OuterStruct{}
	dec.Decode(&actual)

	t.Logf("res: %v", actual)
	expected := OuterStruct{UInt16Value: 0xffff, InnerStructValue: InnerStruct{UInt32Value: 87654321}}
	AssertDeepEquals(actual, expected, t)
}

func TestDecodeNestedBinaryUnMarshalerStruct(t *testing.T) {
	type InnerStruct struct {
		StrValue BytesWithTerminator
	}
	type OuterStruct struct {
		UInt16Value uint16
		InnerStructValue InnerStruct
	}

	src := []byte("\xFF\xFF\x68\x65\x6C\x6C\x6F\x00")
	r := bytes.NewReader(src)
	dec := NewDecoder(r, binary.LittleEndian)

	actual := OuterStruct{}
	dec.Decode(&actual)

	t.Logf("res: %v", actual)
	inner := InnerStruct{StrValue: BytesWithTerminator{Bytes: []byte("hello"), Terminator: byte(0x00)}}
	expected := OuterStruct{UInt16Value: 0xffff, InnerStructValue: inner}
	AssertDeepEquals(actual, expected, t)
}