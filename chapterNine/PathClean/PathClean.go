package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

// ~,環境変数展開後のパス作成
func Clean2(path string) string {
	if len(path) > 1 && path[0:2] == "~/" {
		my, err := user.Current()
		if err != nil {
			panic(err)
		}
		path = my.HomeDir + path[1:]
	}
	path = os.ExpandEnv(path)
	return filepath.Clean(path)
}

func main() {
	// パスをそのままクリーンする
	fmt.Println(filepath.Clean("./path/filepath/../path.go"))
	// path/path.go

	// パスを絶対パスに整形
	abspath, _ := filepath.Abs("path/filepath/path/path_unix.go")
	fmt.Println(abspath)

	// パスを相対パスに整形
	relpath, _ := filepath.Rel("/usr/local/go/src", "/usr/local/go/src/path/filepath/path.go")
	fmt.Println(relpath)

	// 環境変数の展開
	path := os.ExpandEnv("${GOPATH}/src/github.com/golang/example/")
	fmt.Println(path)

	// ホームディレクトリの取得
	fmt.Println(os.UserHomeDir())

	// ~&環境変数展開
	fmt.Println(Clean2("~/${GOPATH}"))

	// 指定したファイル名にパターンが一致するかどうか調べる
	result, _ := filepath.Match("image-*.png", "image-100.png")
	if result {
		fmt.Println("ファイル名:image-100.pngはパターン:image-*.pngに一致")
	}

	// ルールに合致したファイル名の一覧を取得する
	files, err := filepath.Glob("./*.go")
	if err != nil {
		panic(err)
	}
	fmt.Println(files)
}
