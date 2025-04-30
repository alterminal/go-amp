package goamp

import "testing"

func TestBinary(t *testing.T) {
	a := []byte{0x01, 0x02, 0x03, 0x04, 0x05}
	Marshal(a)
}
