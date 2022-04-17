package main

import "fmt"

// 结构体匿名字段

func main() {
	p1 := Person{
		"p1",
		18,
	}
	fmt.Printf("p1: %#v\n", p1)
	fmt.Printf("p1: name: %v, age: %v\n", p1.string, p1.int8)
}

type Person struct {
	string
	int8
}
