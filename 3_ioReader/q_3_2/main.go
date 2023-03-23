package main

import (
	"crypto/rand"
	"os"
)

// create random 1024 bytes data file
func main() {
	file, err := os.Create("random.dat")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	randData := make([]byte, 1024)
	_, err = rand.Read(randData)
	if err != nil {
		panic(err)
	}
	file.Write(randData)
}
