package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	// strings.Reader構造体を取得
	reader := strings.NewReader("Example of io.SelectionReader\n")
	// strings.Reader構造体をもとに、io.SectionReader構造体を取得
	sectionReader := io.NewSectionReader(reader, 14, 7)
	io.Copy(os.Stdout, sectionReader)
}
