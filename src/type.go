package binstruct

import (
	"bytes"
	"fmt"

	"golang.org/x/text/encoding"
)

// StringWithTerminator is string terminated with terminator byte (ex: 0x00)
type StringWithTerminator struct {
	Value      string
	Encoding   encoding.Encoding
	Terminator byte
}

// MarshalBinary
func (s *StringWithTerminator) MarshalBinary() ([]byte, error) {
	enc := s.Encoding.NewEncoder()
	b, err := enc.Bytes([]byte(s.Value))
	if err != nil {
		return nil, err
	}

	termIdx := bytes.IndexByte(b, s.Terminator)
	if termIdx == -1 {
		return nil, fmt.Errorf("string %v does not contain terminator: %v", s.Value, s.Terminator)
	}
	return b[:termIdx+1], nil
}
