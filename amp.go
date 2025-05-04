package goamp

import (
	"bytes"
)

func marshal(buf *bytes.Buffer, v any) {
	if v == nil {
		buf.WriteByte(NIL)
		return
	}
	switch v := v.(type) {
	case bool:
		if v {
			buf.WriteByte(TRU)
		} else {
			buf.WriteByte(FAL)
		}
	case int8:
		buf.WriteByte(NumberInstruction(v))
		buf.Write(NumberToByteArray(v))
	case int16:
		buf.WriteByte(NumberInstruction(v))
		buf.Write(NumberToByteArray(v))
	case int32:
		buf.WriteByte(NumberInstruction(v))
		buf.Write(NumberToByteArray(v))
	case int64:
		buf.WriteByte(NumberInstruction(v))
		buf.Write(NumberToByteArray(v))
	case float32:
		buf.WriteByte(NumberInstruction(v))
		buf.Write(NumberToByteArray(v))
	case float64:
		buf.WriteByte(NumberInstruction(v))
		buf.Write(NumberToByteArray(v))
	case uint8:
		buf.WriteByte(NumberInstruction(v))
		buf.Write(NumberToByteArray(v))
	case uint16:
		buf.WriteByte(NumberInstruction(v))
		buf.Write(NumberToByteArray(v))
	case uint32:
		buf.WriteByte(NumberInstruction(v))
		buf.Write(NumberToByteArray(v))
	case uint64:
		buf.WriteByte(NumberInstruction(v))
		buf.Write(NumberToByteArray(v))
	case string:
		ins, data := StringInstruction(v)
		buf.WriteByte(ins)
		buf.Write(data)
		buf.Write(StringToByteArray(v))
	case []byte:
		ins, data := BinaryHeader(v)
		buf.WriteByte(ins)
		buf.Write(data)
		buf.Write(v)
	case []any:
		buf.WriteByte(LIT)
		for i := range v {
			marshal(buf, v[i])
		}
		buf.WriteByte(END)
	case map[any]any:
		buf.WriteByte(MAP)
		for k, v := range v {
			marshal(buf, k)
			marshal(buf, v)
		}
		buf.WriteByte(END)
	case *int8:
		marshal(buf, *v)
	case *int16:
		marshal(buf, *v)
	case *int32:
		marshal(buf, *v)
	case *int64:
		marshal(buf, *v)
	case *float32:
		marshal(buf, *v)
	case *float64:
		marshal(buf, *v)
	case *uint8:
		marshal(buf, *v)
	case *uint16:
		marshal(buf, *v)
	case *uint32:
		marshal(buf, *v)
	case *uint64:
		marshal(buf, *v)
	case *string:
		marshal(buf, *v)
	case *[]byte:
		marshal(buf, *v)
	case *[]any:
		marshal(buf, *v)
	case *map[any]any:
		marshal(buf, *v)
	default:
		panic("unsupported type")
	}
}

func Marshal(v any) *bytes.Buffer {
	var buf bytes.Buffer
	marshal(&buf, v)
	return &buf
}

func UnmarshalWithEnd(buf *bytes.Buffer) (any, bool) {
	ins, _ := buf.ReadByte()
	switch ins {
	case END:
		return nil, true
	case NIL:
		return nil, false
	case TRU:
		return true, false
	case FAL:
		return false, false
	case U8:
		return BufferToNumber[uint8](buf), false
	case U16:
		return BufferToNumber[uint16](buf), false
	case U32:
		return BufferToNumber[uint32](buf), false
	case U64:
		return BufferToNumber[uint64](buf), false
	case I8:
		return BufferToNumber[int8](buf), false
	case I16:
		return BufferToNumber[int16](buf), false
	case I32:
		return BufferToNumber[int32](buf), false
	case I64:
		return BufferToNumber[int64](buf), false
	case F32:
		return BufferToNumber[float32](buf), false
	case F64:
		return BufferToNumber[float64](buf), false
	case S8:
		l := BufferToNumber[uint8](buf)
		return BufferToString(buf, int(l)), false
	case S16:
		l := BufferToNumber[uint16](buf)
		return BufferToString(buf, int(l)), false
	case S32:
		l := BufferToNumber[uint32](buf)
		return BufferToString(buf, int(l)), false
	case A8:
		l := BufferToNumber[uint8](buf)
		return BufferToString(buf, int(l)), false
	case A16:
		l := BufferToNumber[uint16](buf)
		return BufferToString(buf, int(l)), false
	case A32:
		l := BufferToNumber[uint32](buf)
		return BufferToString(buf, int(l)), false
	case LIT:
		var list []any
		for {
			v, end := UnmarshalWithEnd(buf)
			if end {
				return list, false
			}
			list = append(list, v)
		}
	case MAP:
		m := make(map[any]any)
		for {
			k, end := UnmarshalWithEnd(buf)
			if end {
				return m, false
			}
			v, end := UnmarshalWithEnd(buf)
			if end {
				return m, false
			}
			m[k] = v
		}
	case TUP:
		l := BufferToNumber[uint8](buf)
		tuple := make([]any, l)
		for i := 0; i < int(l); i++ {
			tuple[i], _ = UnmarshalWithEnd(buf)
		}
	}
	return nil, false
}

func Unmarshal(buf *bytes.Buffer) any {
	v, _ := UnmarshalWithEnd(buf)
	return v
}
