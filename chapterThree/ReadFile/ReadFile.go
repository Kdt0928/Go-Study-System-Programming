package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Open("text.txt")
	if err != nil {
		panic(err)
	}
	// deferで現在実行中の処理が終了した段階でfile.Close()を実行する
	defer file.Close()
	io.Copy(os.Stdout, file)
}
