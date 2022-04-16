package main

import "fmt"

// panic 和 recover
func main() {
	a()
	b()
	c()
}

func a() {
	fmt.Println("a")
}

func b() {
	defer func() {
		// recover()必须搭配defer使用。
		// defer一定要在可能引发panic的语句之前定义。
		err := recover()
		if err != nil {
			println("func b err")
		}
	}()
	panic("panic b")
}

func c() {
	fmt.Println("c")
}
