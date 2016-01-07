package binstruct

import (
	"testing"
	"bytes"
	"encoding/binary"
	"encoding/hex"
)

func TestEncodeSingleValue(t *testing.T) {
	type TestStruct struct {
		uint32
		uint16
		uint8
	}

	var b bytes.Buffer
	enc := NewEncoder(&b, binary.LittleEndian)
	if (enc.Encode(int32(12345678)) != nil) {
		t.Fatalf("Encode")
	}

	if b.String() != "\x4e\x61\xbc\x00" {
		t.Errorf("Invalid encode! %v", hex.EncodeToString(b.Bytes()))
	}
}
