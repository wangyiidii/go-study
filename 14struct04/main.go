package main

import "fmt"

// 嵌套结构体的字段冲突

func main() {
	p1 := Person{
		Name:   "p1",
		Gender: "男",
		Age:    18,
		Address: Address{
			Province:   "北京",
			city:       "东城",
			UpdateTime: "2021",
		},
		Email: Email{
			Addr:       "email_addr",
			UpdateTime: "2022",
		},
	}
	fmt.Printf("p1: %#v\n", p1)
	//fmt.Println(p1.UpdateTime) // 不能这么写
	fmt.Println(p1.Address.UpdateTime)
	fmt.Println(p1.Email.UpdateTime)
}

type Person struct {
	Name   string
	Gender string
	Age    int
	//Address Address // 嵌套另外一个结构体
	Address // 嵌套另外一个结构体
	Email
}

type Address struct {
	Province   string
	city       string
	UpdateTime string
}

type Email struct {
	Addr       string
	UpdateTime string
}
