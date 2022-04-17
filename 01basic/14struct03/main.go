package main

import "fmt"

// 嵌套结构体

func main() {
	p1 := Person{
		Name:   "p1",
		Gender: "男",
		Age:    18,
		Address: Address{
			Province: "北京",
			city:     "东城",
		},
	}
	fmt.Printf("p1: %#v\n", p1)
	// 通过嵌套的匿名结构体字段访问其他内部字段
	fmt.Println(p1.Address.Province)
	fmt.Println(p1.Province)
}

type Person struct {
	Name   string
	Gender string
	Age    int
	//Address Address // 嵌套另外一个结构体
	Address // 嵌套另外一个结构体
}

type Address struct {
	Province string
	city     string
}
