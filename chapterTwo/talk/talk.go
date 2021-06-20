// Go言語のインタフェース
// Go言語のインタフェースは、構造体などの具象型が満たすべき仕様、つまりは持つべきメソッドを表現するための言語機能

package main

import "fmt"

// インタフェースを定義
type Talker interface {
	Talk()
}

// 構造体を定義
type Greeter struct {
	name string
}

// 構造体はTalkerインターフェイスで定義されているメソッドを持っている
func (g Greeter) Talk() {
	fmt.Printf("Hello, my name is %s\n", g.name)
}

func main() {
	// インターフェイスの型を持つ変数を宣言
	var talker Talker = &Greeter{"wozozo"}
	talker.Talk()

	var greeter Greeter
	greeter.name = "hoge"
	greeter.Talk()
}
