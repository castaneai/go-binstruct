package binstruct

import (
	"bytes"
	"fmt"
)

// BytesWithTerminator is string terminated with terminator byte (ex: 0x00)
type BytesWithTerminator struct {
	Bytes      []byte
	Terminator byte
}

// MarshalBinary ...
func (s BytesWithTerminator) MarshalBinary() ([]byte, error) {
	termIdx := bytes.IndexByte(s.Bytes, s.Terminator)
	if termIdx == -1 {
		return append(s.Bytes, s.Terminator), nil
	}
	return s.Bytes[:termIdx+1], nil
}

// UnmarshalBinary ...
func (s *BytesWithTerminator) UnmarshalBinary(data []byte) error {
	termIdx := bytes.IndexByte(data, s.Terminator)
	if termIdx == -1 {
		return fmt.Errorf("bytes %v does not contain terminator: %v", data, s.Terminator)
	}
	s.Bytes = data[:termIdx]
	return nil
}
