// Windows環境では動作しない

package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func main() {
	path := filepath.Join(os.TempDir(), "unixdomainsocket-server")
	os.Remove(path)
	fmt.Println("Server is running at " + path)
	// データ送受信のためのnet.PacketConnインターフェースが即座に返る
	conn, err := net.ListenPacket("unixgram", path)
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
