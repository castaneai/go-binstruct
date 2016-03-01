package binstruct

import (
	"bytes"
	"encoding/binary"
	"testing"

	"golang.org/x/text/encoding/japanese"
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
