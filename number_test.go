package goamp

import (
	"testing"
)

func TestInt64(t *testing.T) {
	num := int64(1234567890123456789)
	arr := Marshal(num)
	result := Unmarshal(arr)
	if result == nil {
		t.Errorf("Expected non-nil result, got nil")
	}
	if result != num {
		t.Errorf("Expected %d, got %d", num, result)
	}
	num = int64(0)
	arr = Marshal(num)
	result = Unmarshal(arr)
	if result == nil {
		t.Errorf("Expected non-nil result, got nil")
	}
	if result != num {
		t.Errorf("Expected %d, got %d", num, result)
	}
	num = int64(-1234567890123456789)
	arr = Marshal(num)
	result = Unmarshal(arr)
	if result == nil {
		t.Errorf("Expected non-nil result, got nil")
	}
	if result != num {
		t.Errorf("Expected %d, got %d", num, result)
	}
}

func TestFloat64(t *testing.T) {
	num := float64(1234567890.123456789)
	arr := Marshal(num)
	result := Unmarshal(arr)
	if result == nil {
		t.Errorf("Expected non-nil result, got nil")
	}
	if result != num {
		t.Errorf("Expected %f, got %f", num, result)
	}
	num = float64(0.0)
	arr = Marshal(num)
	result = Unmarshal(arr)
	if result == nil {
		t.Errorf("Expected non-nil result, got nil")
	}
	if result != num {
		t.Errorf("Expected %f, got %f", num, result)
	}
	num = float64(-1234567890.123456789)
	arr = Marshal(num)
	result = Unmarshal(arr)
	if result == nil {
		t.Errorf("Expected non-nil result, got nil")
	}
	if result != num {
		t.Errorf("Expected %f, got %f", num, result)
	}
}

func TestMap(t *testing.T) {
	m := make(map[string]any)
	m["key1"] = "value1"
	m["key2"] = "1234567890"
	m["key3"] = "true"
	arr := Marshal(m)
	result := Unmarshal(arr)
	if result == nil {
		t.Errorf("Expected non-nil result, got nil")
	}
}
