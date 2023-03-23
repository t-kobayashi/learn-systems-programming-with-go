package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	//                           012345678901234567890123456789
	reader := strings.NewReader("Example of io.SecitonReader\n")
	sectionReader := io.NewSectionReader(reader, 14, 7)
	io.Copy(os.Stdout, sectionReader)
}
