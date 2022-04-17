package main

// 多个goroutine并发操作全局变量x
import (
	"fmt"
	"sync"
)

var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwLock sync.RWMutex
)

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}

func add() {
	for i := 0; i < 50000; i++ {
		lock.Lock() // 加锁
		x++
		lock.Unlock() // 释放锁
	}
	wg.Done()
}
