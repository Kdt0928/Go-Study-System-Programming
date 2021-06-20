package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Server is running at localhot:8888")
	// データ送受信のためのnet.PacketConnインターフェースが即座に返る
	conn, err := net.ListenPacket("udp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buffer := make([]byte, 1500)
	for {
		// 通信内容を読み込むと同時に、接続してきた相手のアドレス情報を受け取れる
		length, remoteAdress, err := conn.ReadFrom(buffer)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Received from %v: %v\n", remoteAdress, string(buffer[:length]))
		_, err = conn.WriteTo([]byte("Hello from Server"), remoteAdress)
		if err != nil {
			panic(err)
		}
	}
}
