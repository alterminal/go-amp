package goamp

import (
	"bytes"
	"fmt"
	"reflect"
)

func MarshalMap(v reflect.Value, t reflect.Type, buf *bytes.Buffer) {
	for _, k := range v.MapKeys() {
		marshal(buf, k)
		marshal(buf, v.MapIndex(k))
	}
	buf.WriteByte(END)
}

func UnmarshalMap(buf *bytes.Buffer) map[any]any {
	m := make(map[any]any)
	for {
		key, _ := Unmarshal(buf)
		if key == nil {
			return m
		}
		value, _ := Unmarshal(buf)
		fmt.Println(key)
		m[key] = value
		fmt.Println(key, value)
	}
}

func walk(v reflect.Value) {
	fmt.Printf("Visiting %v\n", v)
	// Indirect through pointers and interfaces
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			walk(v.Index(i))
		}
	case reflect.Map:
		for _, k := range v.MapKeys() {
			walk(v.MapIndex(k))
		}
	default:
		// handle other types
	}
}
