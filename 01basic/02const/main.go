package main

import "fmt"

const pi = 3.14

// 批量声明常量
const (
	OK       = 200
	notFound = 404
)

// 批量声明常量时，如果某一行声明时没有赋值，默认就和上一行一致
const (
	n1 = 100
	n2
	n3
)

// iota. 是go语言的常量计数器，只能在常量的表达式中使用。记住一下两点
// 1.iota在const关键字出现时将被重置为0。
// 2.const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。
const (
	a1 = iota //0
	a2        //1
	a3        //2
)

const (
	b1 = iota //0
	b2 = iota //1
	_  = iota
	b3 = iota //3
)

const (
	c1 = iota // 0
	c2 = 100  // 100
	c3 = iota //2
	c4        //3
)

// 表示流量单位
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

// 定义在一行
const (
	a, b = iota + 1, iota + 2 //1,2
	c, d                      //2,3
	e, f                      //3,4
)

func main() {
	fmt.Printf("n1: %d\n", n1)
	fmt.Printf("n2: %d\n", n2)
	fmt.Printf("n3: %d\n", n3)

	fmt.Printf("a1: %d\n", a1)
	fmt.Printf("a2: %d\n", a2)
	fmt.Printf("a3: %d\n", a3)

	fmt.Printf("b1: %d\n", b1)
	fmt.Printf("b2: %d\n", b2)
	fmt.Printf("b3: %d\n", b3)

	fmt.Printf("c1: %d\n", c1)
	fmt.Printf("c2: %d\n", c2)
	fmt.Printf("c3: %d\n", c3)
	fmt.Printf("c4: %d\n", c4)

	fmt.Printf("KB: %d\n", KB)
	fmt.Printf("MB: %d\n", MB)
	fmt.Printf("GB: %d\n", GB)
	fmt.Printf("TB: %d\n", TB)
	fmt.Printf("PB: %d\n", PB)
}
