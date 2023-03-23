package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	// 32bit Big Endian data (0x00002710=10000)
	data := []byte{0x10, 0x27, 0x00, 0x00}
	var i int32
	// Read Big Endian data
	binary.Read(bytes.NewReader(data), binary.LittleEndian, &i)
	fmt.Println(i) // "10000"
}
