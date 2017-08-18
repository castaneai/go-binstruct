package binstruct

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
	"testing"

	"golang.org/x/text/encoding/japanese"
)

// Convert bytes to hex string (ex)"\x01\x02\x03" -> "01 02 03"
func toHexString(bs []byte) string {
	hs := fmt.Sprintf("%X", bs)
	ss := make([]string, len(bs), len(bs))
	for i := 0; i < len(bs); i++ {
		ss[i] = hs[i*2 : i*2+2]
	}
	return strings.Join(ss, " ")
}

func assertBytes(a, b []byte, t *testing.T) {
	if !bytes.Equal(a, b) {
		t.Errorf("assert bytes failed! %v != %v",
			toHexString(a),
			toHexString(b),
		)
	}
}

func TestEncodeSingleValue(t *testing.T) {
	var b bytes.Buffer
	enc := NewEncoder(&b, binary.LittleEndian)

	// int32
	b.Reset()
	if err := enc.Encode(int32(12345678)); err != nil {
		t.Fatalf(err.Error())
	}
	expected := []byte("\x4e\x61\xbc\x00")
	assertBytes(b.Bytes(), expected, t)

	// int64
	b.Reset()
	if err := enc.Encode(int64(123456789012345)); err != nil {
		t.Fatalf(err.Error())
	}
	expected = []byte("\x79\xdf\x0d\x86\x48\x70\x00\x00")
	assertBytes(b.Bytes(), expected, t)

	// float32
	b.Reset()
	if err := enc.Encode(float32(3.1415)); err != nil {
		t.Fatalf(err.Error())
	}
	expected = []byte("\x56\x0e\x49\x40")
	assertBytes(b.Bytes(), expected, t)

	// float64
	b.Reset()
	if err := enc.Encode(float64(1.234324235435435353)); err != nil {
		t.Fatalf(err.Error())
	}
	expected = []byte("\xb7\xaf\xfd\xc4\xca\xbf\xf3\x3f")
	assertBytes(b.Bytes(), expected, t)
}

func TestEncodeString(t *testing.T) {
	var b bytes.Buffer
	enc := NewEncoder(&b, binary.LittleEndian)

	s := StringWithTerminator{
		Value:      "こんにちは\x00いい天気ですね",
		Encoding:   japanese.ShiftJIS,
		Terminator: byte(0x00),
	}
	if err := enc.Encode(&s); err != nil {
		t.Fatalf(err.Error())
	}
	expected := []byte("\x82\xb1\x82\xf1\x82\xc9\x82\xbf\x82\xcd\x00")
	assertBytes(b.Bytes(), expected, t)
}

func TestEncodeSimpleStruct(t *testing.T) {
	var b bytes.Buffer
	enc := NewEncoder(&b, binary.LittleEndian)

	type SimpleStruct struct {
		Int32Value int32
		Float64Value float64
	}

	s := SimpleStruct{Int32Value: 12345678, Float64Value: 3.141592}
	if err := enc.Encode(&s); err != nil {
		t.Fatalf(err.Error())
	}
	expected := []byte("\x4e\x61\xbc\x00\x7a\x00\x8b\xfc\xfa\x21\x09\x40")
	assertBytes(b.Bytes(), expected, t)
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

	var b bytes.Buffer
	enc := NewEncoder(&b, binary.LittleEndian)
	if err := enc.Encode(&outer); err != nil {
		t.Fatalf(err.Error())
	}

	t.Logf("nested struct: %v", toHexString(b.Bytes()))
}