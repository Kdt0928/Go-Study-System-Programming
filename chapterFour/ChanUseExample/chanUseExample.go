// 第4章 低レベルアクセスへの入り口3：チャネル
// 4.2.1 チャネルの使用方法

package main

import "fmt"

func main() {
	fmt.Println("start sub()")
	// 終了を受け取るためのチャネル
	done := make(chan bool)
	go func() {
		fmt.Println("sub() is finished")
		// 終了を通知
		done <- true
	}()
	// 終了を待つ
	<-done
	fmt.Println("all tasks are finished")
}
