// 2.4.6 io.Writerのデコレータ

package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// 新規ファイル作成
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}

	var builder strings.Builder
	// ファイルと標準出力に書き込み
	// io.MultWriter()は、複数のio.Writerを受け取り、すべてに対して書き込まれた内容を同時に書き込む
	writer := io.MultiWriter(file, os.Stdout, &builder)
	io.WriteString(writer, "io.MultiWriter example\n")
	fmt.Println(builder.String())

	//　gzipファイル作成
	gzipFile, gzipErr := os.Create("text.txt.gz")
	if gzipErr != nil {
		panic(gzipErr)
	}

	gzipFileWriter := gzip.NewWriter(gzipFile)
	// gzip.WriterのHeader構造体が持つNameにファイル名をセット
	gzipFileWriter.Header.Name = "text.txt"
	io.WriteString(gzipFileWriter, "gzip.Writer example\n")
	gzipFileWriter.Close()

	// ある程度の分量ごとに纏めて書き出す
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer")
	buffer.Flush()
	buffer.WriteString("example\n")
	buffer.Flush()

}
