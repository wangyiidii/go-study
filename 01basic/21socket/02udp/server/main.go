package main

// udp server

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 20001,
	})
	if err != nil {
		fmt.Printf("listen failed, err: %v\n", err)
		return
	}
	defer listen.Close()

	for {
		var buf [1024]byte
		n, addr, err := listen.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Printf("ReadFromUDP failed, err: %v\n", err)
			return
		}
		fmt.Println("接收到的数据：", string(buf[:n]))
		_, err = listen.WriteToUDP(buf[:n], addr)
		if err != nil {
			fmt.Printf("WriteToUDP failed, err: %v\n", err)
			return
		}
	}
}
