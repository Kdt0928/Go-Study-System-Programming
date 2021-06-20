// 2.4.5 インターネットアクセスの送信

package main

import (
	"io"
	"net"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "http.ResponseWriter sample")
}

func main() {
	// 通信のコネクションをあらわすインタフェースを取得
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	// net.Connに書き込み
	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n")
	// io.Copyを使って標準出力に表示
	io.Copy(os.Stdout, conn)

	// HTTPリクエストを作成し、リクエストの内容を書き出す
	req, httpErr := http.NewRequest("GET", "http://ascii.jp", nil)
	if httpErr != nil {
		panic(err)
	}
	req.Write(conn)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
