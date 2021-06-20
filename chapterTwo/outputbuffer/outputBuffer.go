// 2.4.3 書かれた内容を記憶しておくバッファ(1) bytes.Bufffer

package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	// bytesの構造体Bufferを取得
	var buffer bytes.Buffer
	// 構造体Bufferに定義されているWriteメソッドを呼び出す
	buffer.Write([]byte("bytes.Buffer 1st example\n"))
	buffer.Write([]byte("bytes.Buffer 2nd example\n"))
	fmt.Println(buffer.String())

	// Write()メソッドを使わない方法
	buffer.WriteString("buffer.WriteString example\n")
	io.WriteString(&buffer, "io.WriteString example\n")
	fmt.Println(buffer.String())
}
