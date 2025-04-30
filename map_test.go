package goamp

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m := make(map[any]any)
	m["hello"] = "world"
	m["world"] = "false2"
	m["good"] = false
	buf := Marshal(m)
	um := Unmarshal(buf)
	umm := um.(map[any]any)
	if len(umm) != len(m) {
		t.Errorf("expected %d, got %d", len(m), len(umm))
	}
	for k, v := range umm {
		fmt.Println(k, v)
	}
}
