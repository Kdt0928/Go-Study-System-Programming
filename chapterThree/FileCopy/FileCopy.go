package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Open("old.txt")
	if err != nil {
		panic(err)
	}
	newFile, err := os.Create("new.txt")
	defer file.Close()
	defer newFile.Close()
	io.Copy(newFile, file)
}
