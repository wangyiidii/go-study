package main

// 指针

import "fmt"

func main() {
	// &: 取指针， *: q取指针的值
	s := "hello"
	sPointer := &s
	fmt.Printf("s: %v, type: %T\n", s, &s)
	fmt.Printf("sPointer: %v, type: %T, value: %v\n", sPointer, &sPointer, *sPointer)

	// 方法传递
	a := 1
	modify01(a)
	fmt.Println(a)
	b := 2
	modify02(&b)
	fmt.Println(b)

	// new 和 make
	// new
	a1 := new(int)
	b1 := new(bool)
	fmt.Printf("a1 type: %T\n", a1)
	fmt.Printf("b1 type: %T\n", b1)
	fmt.Println(*a1)
	fmt.Println(*b1)
	// make (make 只能用作slice、map 和 chan的内存创建，返回的是三个类型本身)
	a2 := make(map[string]string, 1)
	fmt.Printf("a2: v: %v, type: %T\n", a2, a2)
	/*
			new 和 make的区别：
			1. 两者都是用来做内存分配的
			2. make 只能用作slice、map 和 chan的内存创建，返回的是三个类型本身
		   	   new 用于类型的分配，并且对应的值为类型的零值，返回的是指向内存的指针
	*/
}

func modify01(x int) {
	x = 100
}

func modify02(y *int) {
	*y = 100
}
