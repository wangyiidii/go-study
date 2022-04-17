package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写互斥锁
var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwLock sync.RWMutex
)

func main() {
	// 互斥锁
	//start := time.Now()
	//for i := 0; i < 1000; i++ {
	//	wg.Add(1)
	//	go read01()
	//}
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go write01()
	//}
	//wg.Wait()
	//01fmt.Println(time.Now().Sub(start))

	// 读写锁
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read02()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write02()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}

func read01() {
	lock.Lock()
	time.Sleep(time.Millisecond)
	lock.Unlock()
	wg.Done()
}

func write01() {
	lock.Lock()
	x++
	time.Sleep(time.Millisecond * 10)
	lock.Unlock()
	wg.Done()
}

func read02() {
	rwLock.RLock()
	time.Sleep(time.Millisecond)
	rwLock.RUnlock()
	wg.Done()
}

func write02() {
	rwLock.Lock()
	x++
	time.Sleep(time.Millisecond * 10)
	rwLock.Unlock()
	wg.Done()
}
