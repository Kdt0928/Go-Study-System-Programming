package main

import (
	"archive/zip"
	"os"
	"strings"
)

func main() {
	zipFile, zipErr := os.Create("zip.zip")
	if zipErr != nil {
		panic(zipErr)
	}
	zipWriter := zip.NewWriter(zipFile)

	writer, err := zipWriter.Create("newFile.txt")
	if err != nil {
		panic(err)
	}
	defer zipWriter.Close()
	// strings.Reader構造体を取得
	reader := strings.NewReader("WriteZipExample\n")
	reader.WriteTo(writer)
}
