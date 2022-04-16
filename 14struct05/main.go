package main

import "fmt"

// 结构体的继承

func main() {
	d1 := Dog{
		Feet: 4,
		Animal: &Animal{
			name: "d1",
		},
	}
	fmt.Printf("d1: %#v\n", d1)
	d1.wang()
	d1.move()
}

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s move\n", a.name)
}

type Dog struct {
	Feet    int8
	*Animal // 匿名嵌套指针
}

func (d *Dog) wang() {
	fmt.Printf("%s wang\n", d.name)
}
