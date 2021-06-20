// 2.9 問題
// Q2.3 gzipされたJSON出力しながら、標準出力にログを出力

package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	// json化する元データ
	source := map[string]string{
		"Hello": "World",
	}
	data, err := json.Marshal(source)
	if err != nil {
		panic(err)
	}

	// gzipファイル作成
	gzipFile, gzipErr := os.Create("text.gz")
	if gzipErr != nil {
		panic(gzipErr)
	}
	gzipFileWriter := gzip.NewWriter(gzipFile)
	gzipFileWriter.Header.Name = "text.txt"

	// JSON出力と標準出力に書き込み
	writer := io.MultiWriter(gzipFileWriter, os.Stdout)
	io.WriteString(writer, string(data))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
