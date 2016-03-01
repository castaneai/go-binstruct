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

// MarshalBinary ...
func (s *StringWithTerminator) MarshalBinary() ([]byte, error) {
	enc := s.Encoding.NewEncoder()
	b, err := enc.Bytes([]byte(s.Value))
	if err != nil {
		return nil, err
	}

	termIdx := bytes.IndexByte(b, s.Terminator)
	if termIdx == -1 {
		return append([]byte(s.Value), s.Terminator), nil
	}
	return b[:termIdx+1], nil
}

// UnmarshalBinary ...
func (s *StringWithTerminator) UnmarshalBinary(data []byte) error {
	dec := s.Encoding.NewDecoder()
	termIdx := bytes.IndexByte(data, s.Terminator)
	if termIdx == -1 {
		return fmt.Errorf("bytes %v does not contain terminator: %v", data, s.Terminator)
	}
	u8b, err := dec.Bytes(data[:termIdx])
	if err != nil {
		return err
	}
	s.Value = string(u8b)
	return nil
}
