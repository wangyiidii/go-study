package main

import "fmt"

// 函数进阶之-变量作用域

// 全局变量
var num = 10

func main() {
	testGlobal01()
	testGlobal02()
	fmt.Println("main 全局变量: ", num)
	// 外层访问不到函数内部的变量（局部变量），通过返回值访问

	// 函数作为 变量 或 参数
	abc := testGlobal01
	fmt.Printf("abc: %T\n", abc)
	abc()
	r1 := calc(1, 2, add)
	println(r1)
	r2 := calc(3, 1, sub)
	println(r2)
}

func testGlobal01() {
	// 可以在函数中访问全局变量
	fmt.Println("testGlobal01 全局变量: ", num)
}

func testGlobal02() {
	// 先在自己函数中查找，找到了用自己函数中的变量
	num := 100
	fmt.Println("testGlobal02全局变量: ", num)
}

func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
