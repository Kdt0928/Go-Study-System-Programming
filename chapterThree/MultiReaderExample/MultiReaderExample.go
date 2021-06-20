package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	header := bytes.NewBufferString("----- HEADER -----\n")
	content := bytes.NewBufferString("Example of io.MulitReader\n")
	footer := bytes.NewBufferString("----- FOOTER ------\n")

	reader := io.MultiReader(header, content, footer)
	// 全てのreaderをつなげた結果が表示
	io.Copy(os.Stdout, reader)
}
