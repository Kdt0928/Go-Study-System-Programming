package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	for {
		// 5byteのバッファを作成
		buffer := make([]byte, 5)
		// 標準入力を取得
		size, err := os.Stdin.Read(buffer)
		if err == io.EOF {
			fmt.Println("EOF")
		}
		fmt.Printf("size=%d input='%s'\n", size, string(buffer))
	}
}
