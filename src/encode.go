package binstruct

type encBuffer struct {
	data []byte
	scratch [64]byte
}

func (e *encBuffer) WriteByte(c byte) {
	e.data = append(e.data, c)
}

func (e *encBuffer) Write(p []byte) (int, error) {
	e.data = append(e.data, p...)
	return len(p), nil
}
