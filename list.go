package goamp

import "bytes"

func SliceToByteArray(buf *bytes.Buffer, data []any) {
	for i := range data {
		marshal(buf, data[i])
	}
	buf.WriteByte(END)
}
