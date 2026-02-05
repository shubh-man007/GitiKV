package store

import "encoding/binary"

func encodeInt32(i int32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(i))
	return b
}

func decodeInt32(b []byte) int32 {
	return int32(binary.BigEndian.Uint32(b))
}
