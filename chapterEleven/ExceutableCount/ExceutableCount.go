package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func main() {
	count := exec.Command("../DisplayOnePerSecond/DisplayOnePerSecond")
	stdout, _ := count.StdoutPipe()
	go func() {
		fmt.Println("gorutione準備完了")
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			fmt.Printf("(stdout) %s\n", scanner.Text())
		}
	}()
	fmt.Println("準備中")
	err := count.Run()
	if err != nil {
		panic(err)
	}
}
