// 第4章 低レベルアクセスへの入り口3：チャネル
// 4.2.3 for文

package main

import (
	"fmt"
	"math"
)

func primeNumber() chan int {
	// int型を入れるチャネル
	result := make(chan int)
	go func() {
		// resultチャネルに送信
		result <- 2
		for i := 3; i < 100000; i += 2 {
			l := int(math.Sqrt(float64(i)))
			found := false
			for j := 3; j < l; j += 2 {
				if i%j == 0 {
					found = true
					break
				}
			}
			if !found {
				result <- i
			}
		}
		close(result)
	}()
	return result
}

func main() {
	pn := primeNumber()
	// ここがポイント
	for n := range pn {
		fmt.Println(n)
	}
}
