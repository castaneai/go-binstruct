package binstruct

import (
	"bytes"
	"encoding/binary"
	"testing"
)


func TestEncodeSingleValue(t *testing.T) {
	var b bytes.Buffer
	enc := NewEncoder(&b, binary.LittleEndian)

	// int32
	b.Reset()
	if err := enc.Encode(int32(12345678)); err != nil {
		t.Fatalf(err.Error())
	}
	expected := []byte("\x4e\x61\xbc\x00")
	AssertDeepEquals(b.Bytes(), expected, t)

	// int64
	b.Reset()
	if err := enc.Encode(int64(123456789012345)); err != nil {
		t.Fatalf(err.Error())
	}
	expected = []byte("\x79\xdf\x0d\x86\x48\x70\x00\x00")
	AssertDeepEquals(b.Bytes(), expected, t)

	// float32
	b.Reset()
	if err := enc.Encode(float32(3.1415)); err != nil {
		t.Fatalf(err.Error())
	}
	expected = []byte("\x56\x0e\x49\x40")
	AssertDeepEquals(b.Bytes(), expected, t)

	// float64
	b.Reset()
	if err := enc.Encode(float64(1.234324235435435353)); err != nil {
		t.Fatalf(err.Error())
	}
	expected = []byte("\xb7\xaf\xfd\xc4\xca\xbf\xf3\x3f")
	AssertDeepEquals(b.Bytes(), expected, t)
}

func TestEncodeStringWithTerminator(t *testing.T) {
	s := BytesWithTerminator{
		Bytes:      []byte("hello"),
		Terminator: byte(0x00),
	}

	var actual bytes.Buffer
	enc := NewEncoder(&actual, binary.LittleEndian)
	if err := enc.Encode(&s); err != nil {
		t.Fatalf(err.Error())
	}
	expected := []byte("\x68\x65\x6C\x6C\x6F\x00")
	AssertDeepEquals(actual.Bytes(), expected, t)
}

func TestEncodeSimpleStruct(t *testing.T) {
	type SimpleStruct struct {
		Int32Value int32
		Float64Value float64
	}
	s := SimpleStruct{Int32Value: 12345678, Float64Value: 3.141592}

	var actual bytes.Buffer
	enc := NewEncoder(&actual, binary.LittleEndian)
	if err := enc.Encode(&s); err != nil {
		t.Fatalf(err.Error())
	}
	expected := []byte("\x4e\x61\xbc\x00\x7a\x00\x8b\xfc\xfa\x21\x09\x40")
	AssertDeepEquals(actual.Bytes(), expected, t)
}

func TestEncodeNestedStruct(t *testing.T) {
	type InnerStruct struct {
		UInt32Value uint32
	}
	type OuterStruct struct {
		UInt16Value uint16
		InnerStructValue InnerStruct
	}
	inner := InnerStruct{UInt32Value: 87654321}
	outer := OuterStruct{UInt16Value: 0xffff, InnerStructValue: inner}

	var actual bytes.Buffer
	enc := NewEncoder(&actual, binary.LittleEndian)
	if err := enc.Encode(&outer); err != nil {
		t.Fatalf(err.Error())
	}

	expected := []byte("\xFF\xFF\xB1\x7F\x39\x05")
	AssertDeepEquals(actual.Bytes(), expected, t)
}

func TestEncodeNestedBinaryMarshalerStruct(t *testing.T) {
	type InnerStruct struct {
		StrValue BytesWithTerminator
	}
	type OuterStruct struct {
		UInt16Value uint16
		InnerStructValue InnerStruct
	}
	inner := InnerStruct{StrValue: BytesWithTerminator{Bytes: []byte("hello"), Terminator: byte(0x00)}}
	outer := OuterStruct{UInt16Value: 0xffff, InnerStructValue: inner}

	var actual bytes.Buffer
	enc := NewEncoder(&actual, binary.LittleEndian)
	if err := enc.Encode(&outer); err != nil {
		t.Fatalf(err.Error())
	}

	expected := []byte("\xFF\xFF\x68\x65\x6C\x6C\x6F\x00")
	AssertDeepEquals(actual.Bytes(), expected, t)
}