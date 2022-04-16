package main

import "fmt"

// 任意值类型的添加方法

func main() {
	var m1 MyInt
	m1 = 100
	m1.sayHi("m1")
}

// MyInt 基于内置的基本类型造一个我们自己的类型
type MyInt int

func (i MyInt) sayHi(s string) {
	fmt.Println("hello ", s)
}
