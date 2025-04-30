package goamp

func BinaryHeader(data []byte) (byte, []byte) {
	// Binary header for the protocol
	l := len(data)
	switch {
	case l <= 255:
		return B8, NumberToByteArray(uint8(l))
	case l <= 65535:
		return B16, NumberToByteArray(uint16(l))
	case l <= 4294967295:
		return B32, NumberToByteArray(uint32(l))
	default:
		panic("String too long")
	}
}
