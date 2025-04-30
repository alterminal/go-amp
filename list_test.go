package goamp

import (
	"testing"
)

func TestList(t *testing.T) {
	var a []any
	a = append(a, 1)
	a = append(a, "hello")
	Marshal(a)
}
