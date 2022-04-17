package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp client

func main() {
	// 1. 与服务端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("Dial failed, err: %v\n", err)
		return
	}
	// 2. 利用该连接进行数据的发送和接收
	input := bufio.NewReader(os.Stdin)
	for {
		s, err := input.ReadString('\n')
		if err != nil {
			fmt.Printf("ReadString failed, err: %v\n", err)
			return
		}
		s = strings.TrimSpace(s)
		if strings.ToUpper(s) == "Q" {
			return
		}
		// 3. 给服务端发送消息
		_, err = conn.Write([]byte(s))
		if err != nil {
			fmt.Printf("client Write failed, err: %v\n", err)
			return
		}
		// 4. 从服务端接收回复的消息
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("client Read failed, err: %v\n", err)
			return
		}
		fmt.Printf("收到服务端回复：%v\n", string(buf[:n]))
	}
}
