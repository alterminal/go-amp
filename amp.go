package goamp

import (
	"bytes"
	"fmt"
	"reflect"
)

func marshal(v any, buf *bytes.Buffer) {
	switch v := v.(type) {
	case reflect.Value:
		switch v.Kind() {
		case reflect.Ptr, reflect.Interface:
			fmt.Println(reflect.TypeOf(v).Kind())
		case reflect.String:
			marshal(v.String(), buf)
		case reflect.Int64:
			marshal(v.Int(), buf)
		case reflect.Float64:
			marshal(v.Float(), buf)
		}
	case int8:
	case int16:
	case int32:
	case int64:
		buf.WriteByte(IntInstruction(v))
		buf.Write(IntToByteArray(v))
	case float32:
		buf.WriteByte(F32)
		buf.Write(FloatToByteArray(v))
	case float64:
		buf.WriteByte(F64)
		buf.Write(FloatToByteArray(v))
	case bool:
		if v {
			buf.WriteByte(TRU)
		} else {
			buf.WriteByte(FAL)
		}
	}
}

func Marshal(v any) *bytes.Buffer {
	var buf bytes.Buffer
	marshal(v, &buf)
	return &buf
}

func Unmarshal(buf *bytes.Buffer) any {
	ins, _ := buf.ReadByte()
	switch ins {
	case END:
		return nil
	case NIL:
		return nil
	case I64:
		data := make([]byte, 8)
		buf.Read(data)
		return ByteArrayToInt64(data)
	case F64:
		data := make([]byte, 8)
		buf.Read(data)
		return ByteArrayToFloat64(data)
	case S32:
		data := make([]byte, 8)
		buf.Read(data)
		stringLen := ByteArrayToInt64(data)
		bs := make([]byte, stringLen)
		buf.Read(bs)
		return string(bs)
	case TUP:
		return true
	case FAL:
		return false
	case MAP:
		return UnmarshalMap(buf)
	}
	return nil
}
