package goamp

import "bytes"

func StringInstruction(str string) (byte, []byte) {
	l := len(str)
	switch {
	case l <= 255:
		return S8, NumberToByteArray(uint8(l))
	case l <= 65535:
		return S16, NumberToByteArray(uint16(l))
	case l <= 4294967295:
		return S32, NumberToByteArray(uint32(l))
	default:
		panic("String too long")
	}
}

func StringToByteArray(str string) []byte {
	return []byte(str)
}

func BufferToString(buf *bytes.Buffer, l int) string {
	str := make([]byte, l)
	buf.Read(str)
	return string(str)
}
