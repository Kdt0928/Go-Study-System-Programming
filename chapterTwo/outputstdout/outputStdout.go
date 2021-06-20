// 2.4 io.Writerを使う構造体の例
// 2.4.2 画面出力

package main

import (
	"os"
)

func main() {
	os.Stdout.Write([]byte("os.Stdout example\n")) // 標準出力
	os.Stderr.Write([]byte("os.Stderr example\n")) // 標準エラー出力
}
