package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	for i := 0; i < 100; i++ {
		wg.Add(1) // 计数器 +1
		go func(i int) {
			fmt.Println("demo02 ", i)
			wg.Done() // 计数器 -1
		}(i)

	}

	fmt.Println("main")
	wg.Wait()
}
