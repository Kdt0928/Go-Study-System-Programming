// 2.9 問題
// Q2.1 ふぃあるに対するフォーマット出力

package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("text.txt")
	if err != nil {
		panic(err)
	}

	// ファイルに整形した文字列を書き出し
	fmt.Fprintf(file, "数値：%d, 文字列：%s, 浮動小数点：%f", 10, "100", 100.11)
	file.Close()
}
