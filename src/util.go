package binstruct

import (
	"fmt"
	"strings"
	"testing"
	"reflect"
)

// Convert bytes to hex string (ex)"\x01\x02\x03" -> "01 02 03"
func ToHexString(bs []byte) string {
	hs := fmt.Sprintf("%X", bs)
	ss := make([]string, len(bs), len(bs))
	for i := 0; i < len(bs); i++ {
		ss[i] = hs[i*2 : i*2+2]
	}
	return strings.Join(ss, " ")
}

func AssertDeepEquals(a, b interface{}, t *testing.T) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("assert bytes failed! %v != %v", a, b)
	}
}