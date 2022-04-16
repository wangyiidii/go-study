package main

import "fmt"

// 空接口
// 接口中没有定义任何方法时，该接口就是一个空接口
// 任意类型都实现了空接口 --> 空接口变量可以存储任意值

// 空接口的应用
// 1. 空接口类型作为函数的参数
// 2. 空接口可以作为map的value

func main() {
	var x xxx
	x = "hello"
	fmt.Printf("x : %#v\n", x)
	x = 18
	fmt.Printf("x : %#v\n", x)

	//m := make(map[string]xxx, 8)
	m := make(map[string]interface{}, 8)
	m["name"] = "string"
	m["age"] = 18
	m["hobby"] = []string{"篮球", "足球", "乒乓球"}
	fmt.Printf("m : %#v\n", m)

	// 类型断言
	_, ok := x.(bool)
	if ok {
		fmt.Println("是bool")
	} else {
		fmt.Println("不是bool")
	}

	// 使用switch进行类型断言
	switch v := x.(type) {
	case string:
		fmt.Printf("是string类型, value: %#v\n", v)
	case int:
		fmt.Printf("是int类型, value: %#v\n", v)
	default:
		fmt.Printf("未知类型, value: %#v\n", v)
	}
}

// 空接口一般不需要提前定义
type xxx interface {
}
