package server

import "encoding/binary"

// Uint32ToByteArr takes a 32 bit integer as input, converts it to a byte slice, and returns the byte slice
func Uint32ToByteArr(num uint32) []byte {
	// Create slice
	// Since there are 8 bits per byte, and there are 32 bits, they length of the slice
	// must be 32 / 8
	bytes := make([]byte, 32/8)

	// Put num into slice
	binary.LittleEndian.PutUint32(bytes, num)

	return bytes
}

