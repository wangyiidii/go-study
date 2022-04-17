package main

import (
	"fmt"
	"sort"
)

// 切片
func main() {

	a := []string{"a", "b", "c"}
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	// 基于数组得到切片
	a1 := [6]int{1, 2, 3, 4, 5, 6}
	b1 := a1[1:5]
	fmt.Println(b1)
	fmt.Printf("%T\n", b1)
	// 切片再次切片
	c1 := b1[0:len(b1)]
	fmt.Println(c1)
	fmt.Printf("%T\n", c1)
	// make构造切片 (类型，长度，容量)
	d1 := make([]int, 5, 10)
	fmt.Println(d1)
	fmt.Printf("%T\n", d1)
	// 通过len获取切片容量
	fmt.Println(len(d1))
	fmt.Println(cap(d1))

	// nil
	var a2 []int // 声明
	fmt.Println(a2 == nil)
	fmt.Println(a2, len(a2), cap(a2))
	var b2 = []int{} // 声明同时初始化
	fmt.Println(b2 == nil)
	fmt.Println(b2, len(b2), cap(b2))
	var c2 = make([]int, 0)
	fmt.Println(c2 == nil)
	fmt.Println(c2, len(c2), cap(c2))

	// 切片赋值拷贝
	a4 := make([]int, 3) // [0 0 0 ]
	b4 := a4
	b4[0] = 100
	fmt.Println(a4)
	fmt.Println(b4)

	// 切片遍历
	// fori
	a5 := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(a5); i++ {
		fmt.Println(i, a5[i])
	}
	// for range
	for i, v := range a5 {
		fmt.Println(i, v)
	}

	// 切片的扩容 append
	// 切片必须初始化之后才能使用
	var a6 []int
	//a6 = append(a6, 100)
	//01fmt.Println(a6)
	// 一次追加一个
	//for i := 0; i < 10; i++ {
	//	a6 = append(a6, i)
	//	01fmt.Printf("%v len: %d cap: %d ptr: %p\n", a6, len(a6), cap(a6), a6)
	//}
	// 一次追加多个参数
	a6 = append(a6, 1, 2, 3)
	fmt.Println(a6)
	b6 := []int{11, 12, 13}
	a6 = append(a6, b6...)
	fmt.Println(a6)

	// copy() 赋值切片
	a7 := []int{1, 2, 3}
	b7 := make([]int, 3, 3)
	copy(b7, a7)
	b7[0] = 100
	fmt.Println(a7)
	fmt.Println(b7)

	// 切片删除元素
	a8 := []int{1, 2, 3, 4, 5}
	a8 = append(a8[:3], a8[4:]...)
	fmt.Println(a8)

	// 练习题
	// 1. 输出结果
	var a9 = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a9 = append(a9, fmt.Sprintf("%v", i))
	}
	fmt.Println(a9)
	// 2.
	// 请使用内置的sort包对数组var a = [...]int{3, 7, 8, 9, 1}进行排序（附加题，自行查资料解答）
	var a10 = [...]int{3, 7, 8, 9, 1}
	sort.Ints(a10[:])
	//sort.Reverse(sort.Ints(a10[:]))
	fmt.Println(a10)
}
