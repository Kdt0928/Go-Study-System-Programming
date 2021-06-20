// 2.4 io.Writerを使う構造体の例
// 2.4.1 ファイル出力

package main

import (
	"os"
)

func main() {
	// 新規ファイルのos.Fileのインスタンス作成
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	// 文字列をバイト列に変換しファイルに書き込み
	file.Write([]byte("os.File example\n"))
	file.Close()
}
