package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(8)
	a()
	b()
}

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("a ", i)
	}
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("b ", i)
	}
}
