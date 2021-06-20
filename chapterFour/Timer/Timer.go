package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("開始:", time.Now())
	second := 10
	timer := time.After(time.Duration(second) * time.Second)
	t := make(chan int)
	exit := make(chan struct{})

	go func(chint chan<- int, second int, end chan<- struct{}) {
		for i := 1; i <= second; i++ {
			chint <- i
			time.Sleep(time.Second)
		}
		close(end)
	}(t, second, exit)

	for {
		select {
		case time := <-t:
			fmt.Printf("%vs\n", time)
		case end := <-timer:
			fmt.Println(end)
		case <-exit:
			fmt.Println("終了")
			return
		}
	}
}
