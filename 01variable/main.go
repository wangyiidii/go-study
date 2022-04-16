package main

import "fmt"

// 单独声明
//var name string
//var age int
//var isOk bool

// 批量声明
var (
	name string
	age  int
	isOk bool
	test string
)

func main() {
	// 全局变量声明了但是可以不初始化，也可以不使用
	name = "理想"
	age = 16
	isOk = true

	// 局部变量声明必须初始化，必须使用
	//var a string = "s"
	fmt.Printf("name: %s, age: %d, isOk: %t\n", name, age, isOk)

	// 声明变量同时赋值
	// 类型推导，可省略不写
	// var s1 string = "wang"
	var s1 = "wang"
	fmt.Printf("s1: %s\n", s1)

	// 简短变量声明，不能在函数外使用
	s2 := "s3"
	fmt.Printf("s2: %s\n", s2)

	// 匿名变量 "_"
	var _ = ""

}
