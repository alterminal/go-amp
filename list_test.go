package goamp

import (
	"testing"
)

func TestList(t *testing.T) {
	var a []any
	a = append(a, int32(1))
	a = append(a, "hello")
	buf := Marshal(a)
	b, _ := Unmarshal(buf)
	n, ok := b.([]any)
	if !ok {
		t.Errorf("expected []any, got %T", b)
	}
	if len(n) != len(a) {
		t.Errorf("expected %d, got %d", len(a), len(n))
	}
	// for i := range n {
	// 	if n[i] != a[i] {
	// 		t.Errorf("expected %v, got %v", a[i], n[i])
	// 	}
	// }
}
