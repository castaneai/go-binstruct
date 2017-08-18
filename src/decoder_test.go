package binstruct

import (
	"bytes"
	"encoding/binary"
	"testing"

	"golang.org/x/text/encoding/japanese"
	"reflect"
)

func TestDecodeString(t *testing.T) {
	src := []byte("\x82\xb1\x82\xf1\x82\xc9\x82\xbf\x82\xcd\x00\x82\xb1\x82")
	expected := "こんにちは"

	r := bytes.NewReader(src)
	dec := NewDecoder(r, binary.LittleEndian)
	s := StringWithTerminator{Encoding: japanese.ShiftJIS, Terminator: 0x00}
	err := dec.Decode(&s)
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Logf("s.Value: %v", toHexString([]byte(s.Value)))
	t.Logf("expected: %v", toHexString([]byte(expected)))
	if s.Value != expected {
		t.Errorf("%v != %v\n", s.Value, expected)
	}
}

func TestDecodeNestedStruct(t *testing.T) {
	src := []byte("\xFF\xFF\xB1\x7F\x39\x05")

	type InnerStruct struct {
		UInt32Value uint32
	}
	type OuterStruct struct {
		UInt16Value uint16
		InnerStructValue InnerStruct
	}

	r := bytes.NewReader(src)
	dec := NewDecoder(r, binary.LittleEndian)

	actual := OuterStruct{}
	dec.Decode(&actual)

	t.Logf("res: %v", actual)
	expected := OuterStruct{UInt16Value: 0xffff, InnerStructValue: InnerStruct{UInt32Value: 87654321}}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v != %v\n", actual, expected)
	}
}