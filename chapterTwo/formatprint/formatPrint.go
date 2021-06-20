// 2.4.7 formatしてデータをio.Writerに書き出す

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	// %vを日付に置き換えて標準出力に出力する
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v\n", time.Now())

	// JSONを整形して書き出す
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "    ")
	encoder.Encode(map[string]string{
		"example": "encodeing/json",
		"hello":   "world",
	})

	// http.Request構造体を取得
	request, err := http.NewRequest("GET", "http://ascii.jp", nil)
	if err != nil {
		panic(err)
	}
	// 構造体にヘッダーを追加
	request.Header.Set("X-TEST", "ヘッダーも追加できます。")
	// http.Requestの内容を標準出力に書き込み
	request.Write(os.Stdout)
}
