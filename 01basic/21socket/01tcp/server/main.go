package main

import (
	"bufio"
	"fmt"
	"net"
)

// tcp server

func main() {
	// 1. 开启服务
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("listen failed, err: %v\n", err)
		return
	}
	for {
		// 2. 等待客户端连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err: %v\n", err)
			return
		}
		// 3. 启动一个单独的goroutine去处理连接
		go process(conn)
	}
}

func process(conn net.Conn) {
	// 处理完之后要关闭这个连接
	defer conn.Close()
	// 针对当前的连接做数据的发送和连接操作
	for {
		reader := bufio.NewReader(conn)
		var buf [127]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("read from conn failed, err: %v\n", err)
			return
		}
		recv := string(buf[:n])
		fmt.Println("接收到的数据：", recv)
		// 回复客户端
		conn.Write([]byte("ok"))
	}
}
