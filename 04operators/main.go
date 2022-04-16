package main

import "fmt"

// 运算符
func main() {
	// 1. 算数运算符
	a := 10
	b := 20
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(b / a)
	fmt.Println(5 % 2)
	a++
	b--
	fmt.Println(a)
	fmt.Println(b)

	// 2. 关系运算符 (==, !=, >, >=, <, <= )
	fmt.Println(10 > 2)
	fmt.Println(10 < 2)

	// 3. 逻辑运算符 (&&, ||, !)
	fmt.Println(10 > 5 && 1 == 1)
	fmt.Println(!(10 > 5))
	fmt.Println(1 > 5 || 1 == 1)

	// 4. 位运算符 (&, |, ^, <<, >>)
	a1 := 1 // 0001
	b1 := 5 // 0101
	println(a1 & b1)
	println(a1 | b1)
	println(a1 ^ b1)
	println(1 << 2)
	println(4 >> 2)
	println(1 << 10)

	// 5. 赋值运算符
	var a2 int
	a2 = 5
	a2 += 5
	println(a2)

}
