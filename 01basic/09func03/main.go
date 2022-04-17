package main

import (
	"fmt"
	"strings"
)

func main() {
	anonymousFunc()
	closure()
}

// 匿名函数
func anonymousFunc() {
	// 1
	sayHello := func() {
		fmt.Println("hello 01mysql")
	}
	sayHello()
	// 2
	func() {
		fmt.Println("hello 02")
	}()
}

// 闭包
// 闭包 = 函数 + 引用环境
func closure() {
	f1 := test01()
	f1() // 相当于执行了test01函数里面的匿名函数

	f2 := test02() // f2此时就是一个闭包
	f2()           // 相当于执行了test02函数里面的匿名函数
	f3 := test03("java")
	f3()

	f4 := makeSuffixFunc(".txt")
	r4 := f4("log.tx")
	fmt.Println("r4: ", r4)

	add, sub := calc(100)
	addR := add(1) // base =  100 + 1
	subR := sub(2) // base =  101 -2
	fmt.Println("add: ", addR)
	fmt.Println("sub: ", subR)

}

// 返回值是一个func
func test01() func() {
	return func() {
		fmt.Println("hello closure")
	}
}
func test02() func() {
	name := "tom"
	return func() {
		fmt.Println("hello ", name)
	}
}

func test03(name string) func() {
	return func() {
		fmt.Println("hello ", name)
	}
}

// 闭包进阶2
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 闭包进阶4
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}
