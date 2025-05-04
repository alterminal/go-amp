package goamp

import "testing"

type St struct {
	Name string
	Age  int
}

func TestStruct(t *testing.T) {
	var hello int8

	Marshal(&hello)
}
