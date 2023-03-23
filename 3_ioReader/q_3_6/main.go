package main

import (
	"io"
	"os"
	"strings"
)

var (
	computer    = strings.NewReader("COMPUTER")
	system      = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

func main() {
	var stream io.Reader
	// ここに io パッケージを使ったコードを書く
	stream = io.MultiReader(
		io.NewSectionReader(programming, 5, 1),
		io.NewSectionReader(system, 0, 1),
		io.NewSectionReader(computer, 0, 1),
		io.NewSectionReader(programming, 8, 1),
		io.NewSectionReader(programming, 8, 1),
	)
	io.Copy(os.Stdout, stream)
}
