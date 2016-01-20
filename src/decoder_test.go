package binstruct

import (
	"encoding/binary"
	"testing"
)

func TestDecodeString(t *testing.T) {
	src := []byte("\x82\xb1\x82\xf1\x82\xc9\x82\xbf\x82\xcd\x00\x82\xb1\x82")
	expected := "こんにちは"

	dec := NewDecoder(&src, binary.LittleEndian)
	var s StringWithTerminator
	dec.Decode(&s)
	if s.Value != expected {
		t.Errorf("%v != %v\n", s.Value, expected)
	}
}
