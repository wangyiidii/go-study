package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	for i := 0; i < 100; i++ {

		go hello(i) // 开启一个goroutine去执行hello函数
		wg.Add(1)   // 计数器 +1
	}

	fmt.Println("main")

	wg.Wait()
}

func hello(i int) {
	fmt.Println("hello ", i)
	wg.Done() // 计数器 -1
}
