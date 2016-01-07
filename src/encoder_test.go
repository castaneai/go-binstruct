package binstruct

import (
	"testing"
	"bytes"
)

func Marshal(t *testing.T) {
	type TestStruct struct {
		uint32
		uint16
		uint8
	}

	ts := TestStruct{12345678, 12345, 123}
	var bytes bytes.Buffer
	enc := NewEncoder(&bytes)
	dec := NewDecoder(&bytes)

	enc.Encode(ts)
}
