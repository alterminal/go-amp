package goamp

import (
	"bytes"
	"fmt"
	"reflect"
	"unicode/utf8"
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
	case int64:
		buf.WriteByte(INT)
		buf.Write(Int64ToByteArray(v))
	case float64:
		buf.WriteByte(FLO)
		buf.Write(Float6464ToByteArray(v))
	case string:
		buf.WriteByte(U8)
		var buffer []byte
		for _, r := range v {
			buffer = utf8.AppendRune(buffer, r)
		}
		buf.Write((Int64ToByteArray(int64(len(buffer)))))
		for _, b := range buffer {
			buf.WriteByte(b)
		}
	case bool:
		if v {
			buf.WriteByte(TRU)
		} else {
			buf.WriteByte(FAL)
		}
	default:
		val := reflect.ValueOf(v)
		if val.IsNil() {
			buf.WriteByte(END)
			return
		}
		typ := reflect.TypeOf(v)
		if typ.Kind() == reflect.Slice {
			buf.WriteByte(LIS)
			return
		}
		if typ.Kind() == reflect.Map {
			buf.WriteByte(MAP)
			MarshalMap(val, typ, buf)
			return
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
	case INT:
		data := make([]byte, 8)
		buf.Read(data)
		return ByteArrayToInt64(data)
	case FLO:
		data := make([]byte, 8)
		buf.Read(data)
		return ByteArrayToFloat64(data)
	case U8:
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

func UnmarshalInt64(buf *bytes.Buffer) int64 {
	data, _ := buf.ReadBytes(8)
	return ByteArrayToInt64(data)
}
