package main

import (
	"flag"
	"io"
	"os"
)


func main() {
	from := flag.String("from", "old.txt", "from")
	to := flag.String("to", "new.txt", "to")
	flag.Parse()

	r, err := os.Open(*from)
	if err != nil {
		panic(err)
	}
	defer r.Close()
	w, err := os.Create(*to)
	if err != nil {
		panic(err)
	}
	defer w.Close()
	if _, err := io.Copy(w, r); err != nil {
		panic(err)
	}
}
