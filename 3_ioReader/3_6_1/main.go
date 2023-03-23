package main

import (
	"bufio"
	"fmt"
	"strings"
)

var source = `1 行め
2 行め
3 行め`

// func main() {
// 	reader := bufio.NewReader(strings.NewReader(source))
// 	for {
// 		line, err := reader.ReadString('\n')
// 		fmt.Printf("%#v\n", line)
// 		if err == io.EOF {
// 			break
// 		}
// 	}
// }

func main() {
	scanner := bufio.NewScanner(strings.NewReader(source))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}
}
