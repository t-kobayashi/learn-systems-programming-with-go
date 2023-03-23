package main

import (
	"archive/zip"
	"io"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=test.zip")
	zipWriter := zip.NewWriter(w)
	zipFile, err := zipWriter.Create(strings.Repeat("a", 10000000))
	if err != nil {
		panic(err)
	}
	io.WriteString(zipFile, "test data")
	zipWriter.Close()
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
