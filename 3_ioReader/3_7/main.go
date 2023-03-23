package main

import (
	"bytes"
	"fmt"
	"io"
)

// func main() {
// 	header := bytes.NewBufferString("----- HEADER -----\n")
// 	content := bytes.NewBufferString("Example of io.MultiReader\n")
// 	footer := bytes.NewBufferString("----- FOOTER -----\n")
// 	reader := io.MultiReader(header, content, footer)
// 	// すべての reader をつなげた出力が表示
// 	io.Copy(os.Stdout, reader)
// }


func main() {
	var buffer bytes.Buffer
	reader := bytes.NewBufferString("Example of io.TeeReader\n")
	teeReader := io.TeeReader(reader, &buffer)
	// データを読み捨てる
	_, _ = io.ReadAll(teeReader)
	// けどバッファに残ってる
	fmt.Println(buffer.String())
}
