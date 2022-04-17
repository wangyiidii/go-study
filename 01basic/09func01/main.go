package main

import "fmt"

func main() {
	helloWorld()
	say("tom")
	sum := intSum(1, 2)
	fmt.Println("sum: ", sum)
	intSum01(2, 3)
	a := intSum03(1, 2, 3)
	fmt.Println("sum03: ", a)
	a1 := intSum04(5, 1, 2, 3)
	fmt.Println("sum04: ", a1)
	calc(5, 2)
}

// 无参数，无返回值
func helloWorld() {
	fmt.Println("hello world")
}

// 一个参数
func say(name string) {
	fmt.Println("hello ", name)
}

// 多个参数，一个返回值
func intSum(a int, b int) int {
	sum := a + b
	fmt.Println("sum: ", sum)
	return sum
}

// 多个参数，一个返回值. 给返回值命名
func intSum01(a int, b int) (sum int) {
	sum = a + b
	fmt.Println("sum: ", sum)
	return
}

// 接收可变参数
// ...表示参数是个slice类型
func intSum03(a ...int) int {
	fmt.Printf("type a: %T\n", a)
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

// 固定参数和可变参数同时出现时，可变参数要放在最后
func intSum04(a int, b ...int) int {
	sum := a
	for _, v := range b {
		sum += v
	}
	return sum
}

// go语言中参数类型简写
// 多个返回值的参数
func calc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	fmt.Println("sum: ", sum)
	fmt.Println("sub: ", sub)
	return
}
