// 第4章 低レベルアクセスへの入り口3：チャネル
// 4.1 goroutine

package main

import (
	"fmt"
	"time"
)

// 新しく作られるgoroutioneが呼ぶ関数
func sub() {
	fmt.Println("sub() is running")
	time.Sleep(time.Second)
	fmt.Println("sub() is finished")
}

func main() {
	fmt.Println("start sub()")
	// goroutineと作って関数を実行
	go sub()
	time.Sleep(2 * time.Second)

	fmt.Println("start sub()")
	// インラインで無形関数を作ってその場でgoroutioneで実行
	go func() {
		fmt.Println("sub() is running")
		time.Sleep(time.Second)
		fmt.Println("sub() is finished")
	}()
	time.Sleep(2 * time.Second)
}
