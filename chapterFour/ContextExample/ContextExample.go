// 第4章 低レベルアクセスへの入り口3：チャネル
// 4.2.5 コンテキスト
// 終了を受け取るctxと、そのコンテキストを終了させるcancel関数をcontext.WithCancel()関数を通じて取得して利用している

package main

import (
	"context"
	"fmt"
)

func main() {
	fmt.Println("start sub()")
	// 終了を受け取るための終了関数付きコンテキスト
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		fmt.Println("sub() is finished")
		// 終了を通知
		cancel()
	}()
	// 終了を待つ
	<-ctx.Done()
	fmt.Println("all tasks are finished")
}
