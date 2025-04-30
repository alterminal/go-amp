package goamp

import "testing"

func TestNumber(t *testing.T) {
	var a int64 = 1234567890123456789
	buf := Marshal(a)
	// buf := Marshal(a)
	b, _ := Unmarshal(buf)
	n, ok := b.(int64)
	if !ok {
		t.Errorf("expected int64, got %T", b)
	}
	if n != a {
		t.Errorf("expected %d, got %d", a, n)
	}
	var c int64 = -12312
	buf2 := Marshal(c)
	// buf := Marshal(a)
	d, _ := Unmarshal(buf2)
	n2, ok := d.(int64)
	if !ok {
		t.Errorf("expected int64, got %T", b)
	}
	if n2 != c {
		t.Errorf("expected %d, got %d", a, n)
	}
}
