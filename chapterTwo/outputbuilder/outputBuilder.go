// 2.4.4 書かれた内容を記憶しておくバッファ(2) strings.Builder

package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder
	builder.Write([]byte("string.Builder 1st example\n"))
	builder.Write([]byte("string.Builder 2nd example\n"))
	fmt.Println(builder.String())
}
