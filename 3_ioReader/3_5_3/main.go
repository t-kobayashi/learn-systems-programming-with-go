package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)


func main() {
	file, err := os.Open("PNG_transparency_demonstration_1.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	chunks := readChunks(file)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
}

func readChunks(file *os.File) []io.Reader {
	var chunks []io.Reader

	file.Seek(8, 0) // skip 8 bytes of PNG signature
	var offset int64 = 8

	for {
		var length int32
		err := binary.Read(file, binary.BigEndian, &length)
		if err == io.EOF {
			break
		}
		chunks = append(chunks, io.NewSectionReader(file, offset, int64(length)+12)) // 12=length+type+crc
		offset, _ = file.Seek(int64(length+8), 1) // whence: 0=seekStart, 1=seekCurrent, 2=seekEnd
	}
	return chunks
}

func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer) // read type
	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length);
}
