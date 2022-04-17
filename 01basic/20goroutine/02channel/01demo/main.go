package main

import "fmt"

func main() {
	//var ch1 chan int
	//ch1 = make(chan int, 1) // 引用类型，需要初始化后才能使用

	//ch1 := make(chan int)// 无缓冲区通道，又称为同步通道
	ch1 := make(chan int, 1) // 带缓冲区通道，又称为异步通道
	ch1 <- 10                // 发送值
	x := <-ch1
	fmt.Println(x)
	close(ch1)
}
