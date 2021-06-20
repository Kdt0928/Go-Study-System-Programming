package main

import "os"

func main() {
	file, err := os.Create("text.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write([]byte("system call example\n"))
}
