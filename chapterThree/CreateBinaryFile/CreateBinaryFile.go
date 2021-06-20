package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	buffer := make([]byte, 1024)
	_, err := io.ReadFull(rand.Reader, buffer)
	if err != nil {
		panic(nil)
	}
	file, fileErr := os.Create("binaryFile")
	if fileErr != nil {
		panic(nil)
	}
	file.Write(buffer)
	file.Close()
}
