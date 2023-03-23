package main

import (
	"archive/zip"
	"io"
	"os"
)

func main() {
	file, err := os.Create("test.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()
	zipFile, err := zipWriter.Create("test.txt")
	if err != nil {
		panic(err)
	}
	// write data to zip file
	io.WriteString(zipFile, "test data")
}
