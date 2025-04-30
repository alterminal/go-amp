package goamp

import (
	"bytes"
)

func marshal(v any, buf *bytes.Buffer) {
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
	case TRU:
		return true
	case FAL:
		return false
	case U8:
		return BufferToNumber[uint8](buf)
	case U16:
		return BufferToNumber[uint16](buf)
	case U32:
		return BufferToNumber[uint32](buf)
	case U64:
		return BufferToNumber[uint64](buf)
	case I8:
		return BufferToNumber[int8](buf)
	case I16:
		return BufferToNumber[int16](buf)
	case I32:
		return BufferToNumber[int32](buf)
	case I64:
		return BufferToNumber[int64](buf)
	case F32:
		return BufferToNumber[float32](buf)
	case F64:
		return BufferToNumber[float64](buf)
	case S8:
		l := BufferToNumber[uint8](buf)
		return BufferToString(buf, int(l))
	case S16:
		l := BufferToNumber[uint16](buf)
		return BufferToString(buf, int(l))
	case S32:
		l := BufferToNumber[uint32](buf)
		return BufferToString(buf, int(l))
	}
	return nil
}
