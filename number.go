package goamp

import (
	"bytes"
	"unsafe"
)

func NumberInstruction[T int8 | int16 | int32 | int64 | float32 |
	float64 | uint8 | uint16 | uint32 | uint64](num T) byte {
	switch any(num).(type) {
	case int8:
		return I8
	case int16:
		return I16
	case int32:
		return I32
	case int64:
		return I64
	case float32:
		return F32
	case float64:
		return F64
	case uint8:
		return U8
	case uint16:
		return U16
	case uint32:
		return U32
	case uint64:
		return U64
	default:
		panic("Unsupported type")
	}
}

func NumberToByteArray[T int8 | int16 | int32 | int64 | float32 |
	float64 | uint8 | uint16 | uint32 | uint64](num T) []byte {
	size := int(unsafe.Sizeof(num))
	arr := make([]byte, size)
	for i := 0; i < size; i++ {
		byt := *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&num)) + uintptr(i)))
		arr[i] = byt
	}
	return arr
}

func BufferToNumber[T int8 | int16 | int32 | int64 | float32 |
	float64 | uint8 | uint16 | uint32 | uint64](buf *bytes.Buffer) T {
	var v T
	size := int(unsafe.Sizeof(v))
	arr := make([]byte, size)
	buf.Read(arr)
	for i := 0; i < size; i++ {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&v)) + uintptr(i))) = arr[i]
	}
	return v
}

func ByteArrayToFloat64(arr []byte) float64 {
	val := float64(0)
	size := len(arr)
	for i := 0; i < size; i++ {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&val)) + uintptr(i))) = arr[i]
	}
	return val
}
