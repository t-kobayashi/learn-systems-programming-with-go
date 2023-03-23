package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	s := strings.NewReader("Hello, World!")
	copyN(os.Stdout, s, 5)
}

func copyN(dest io.Writer, src io.Reader, n int64) (written int64, err error) {
	return io.Copy(dest, io.LimitReader(src, n))
}
