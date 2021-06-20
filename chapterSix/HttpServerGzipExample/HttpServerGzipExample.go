package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

// クライアントはgzipを受け入れ可能か？
func isGzipAcceptalbe(request *http.Request) bool {
	return strings.Index(
		strings.Join(request.Header["Accept-Encoding"], ","), "gzip") != -1
}

// 1セッションの処理をする
func processSession(conn net.Conn) {
	fmt.Printf("Aceept %v\n", conn.RemoteAddr())
	defer conn.Close()
	for {
		// タイムアウトを設定
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		// リクエストを読み込む
		// HTTPリクエストのヘッダー、メソッド、パスなどの情報を切り出し
		request, err := http.ReadRequest(bufio.NewReader(conn))
		if err != nil {
			// タイムアウトもしくはソケットクローズ時は終了
			// それ以外はエラーにする
			neterr, ok := err.(net.Error) // 型アサーション インターフェースerrがnet.Errorを保持している場合は、OKがTrueになる
			if ok && neterr.Timeout() {
				fmt.Println("Timeout")
				break
			} else if err == io.EOF {
				break
			}
			panic(err)
		}
		// リクエストを表示
		dump, err := httputil.DumpRequest(request, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))
		// レスポンスを書き込む
		// HTTP/1.1かつ、ContentLengthの設定が必要
		response := http.Response{
			StatusCode: 200,
			ProtoMajor: 1,
			ProtoMinor: 1,
			Header:     make(http.Header),
		}
		if isGzipAcceptalbe(request) {
			content := "Hello World (gzippend)\n"
			// コンテンツをgzip化して転送
			var buffer bytes.Buffer
			writer := gzip.NewWriter(&buffer)
			io.WriteString(writer, content)
			writer.Close()
			response.Body = ioutil.NopCloser(&buffer)
			response.ContentLength = int64(buffer.Len())
			response.Header.Set("Content-Encoding", "gzip")
		} else {
			content := "Hello World\n"
			response.Body = io.NopCloser(strings.NewReader(content))
			response.ContentLength = int64(len(content))
		}
		response.Write(conn)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(nil)
	}
	fmt.Println("Server is runnning at localhost:8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(nil)
		}
		go processSession(conn)
	}
}
