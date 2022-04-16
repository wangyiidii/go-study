package main

import "fmt"

// 使用值接收者实现接口和使用指针接收者实现接口的区别

func main() {
	var m mover // 定义一个mover类型的变量

	//p1 := person{"p1"} // 值类型
	//m = p1             // ？ 无法保存，因为p1是是person类型的值，没有实现mover接口
	//m.move()

	p2 := &person{"p2"} // 指针类型
	m = p2
	m.move()

	var s sayer // 定义一个sayer类型的变量
	s = p2
	s.say()

	var a animal // 定义一个animal类型的变量
	a = p2
	a.move()
	a.say()
}

type animal interface {
	mover
	sayer
}

type mover interface {
	move()
}

type sayer interface {
	say()
}

type person struct {
	name string
}

// 使用值接收者实现接口：类型的值和类型的指针都能够保存到接口的变量中
//func (p person) move() {
//	fmt.Printf("%s move\n", p.name)
//}

// 使用指针接收者实现接口：只有类型指针保存到接口变量中
func (p *person) move() {
	fmt.Printf("%s move\n", p.name)
}

func (p *person) say() {
	fmt.Printf("%s say\n", p.name)
}
