package goamp

import "testing"

func TestString(t *testing.T) {
	a := "hello 你好嗎"
	buf := Marshal(a)
	b := Unmarshal(buf)
	n, ok := b.(string)
	if !ok {
		t.Errorf("expected string, got %T", b)
	}
	if n != a {
		t.Errorf("expected %s, got %s", a, n)
	}
}
