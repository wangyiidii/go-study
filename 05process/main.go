package main

import "fmt"

func main() {

	// 1. if
	var score = 60
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
	// score1只在if作用域有效
	if score1 := 60; score1 >= 90 {
		fmt.Println("A")
	} else if score1 > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	// 2. for
	// 2.1 for循环基本写法
	for i := 0; i < 10; i++ {
		fmt.Println("i: ", i)
	}
	// 2.2 省略初始语句，但是必须保留初始语句后面的分号
	var j = 0
	for ; j < 10; j++ {
		fmt.Println("j: ", j)
	}
	// 2.3 省略初始语句和结束语句
	var k = 10
	for k > 0 {
		fmt.Println("k: ", k)
		k--
	}
	// 2.4 无限循环
	//for {
	//	fmt.Println("hello")
	//}
	// 2.5 跳出循环
	for i := 0; i < 10; i++ {
		fmt.Println("2.5: ", i)
		if i == 3 {
			break
		}
	}
	// 2.6 continue 继续下一次循环
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}
		fmt.Println("2.6: ", i)
	}
	// 2.7 for range 循环

	// 3 switch case
	// case 可以一次使用多个值
	// case 可以使用表达式
	finger := 1
	switch finger {
	case 1, 5:
		fmt.Println("1")
		fmt.Println("5")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	}
	age := 30
	switch {
	case age > 18:
		fmt.Println(">18")
	case age < 18:
		fmt.Println("<18")
	default:
		fmt.Println("default")
	}

	// 编写代码打印9*9乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}
}
