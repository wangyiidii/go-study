package main

import (
	"fmt"
	"sync"
)

var (
	m  = make(map[int]int) // 原生的map是有线程安全问题的
	m2 = sync.Map{}
	wg sync.WaitGroup
)

func main() {
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			//set01(i, i*100)
			//01fmt.Printf("key: %v, value: %v", i, get01(i))

			m2.Store(i, i*100)
			value, _ := m2.Load(i)
			fmt.Printf("key: %v, value: %v\n", i, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func get01(key int) int {
	return m[key]
}

func set01(key, value int) {
	m[key] = value
}
