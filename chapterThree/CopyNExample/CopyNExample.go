package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Open("CopyNExample.txt")
	if err != nil {
		panic(err)
	}
	// 全てコピー
	// io.Copy(os.Stdout, file)
	// 指定したサイズだけコピー
	io.CopyN(os.Stdout, file, 5)
}
