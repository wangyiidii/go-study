package main

import "fmt"

func main() {
	c := Cat{}
	c.say()
	d := Dog{}
	d.say()

	c1 := Cat{}
	c2 := Dog{}
	hit(c1)
	hit(c2)

	p1 := person{
		name: "tom",
	}
	hit(p1)

	var s1 sayer
	s1 = person{"s1"}
	fmt.Printf("%#v\n", s1)

}

func hit(arg sayer) {
	// 不管是传进来的是什么，我都要打Ta，打TaTa就会叫，就要执行Ta的say方法
	arg.say()
}

// sayer 定义一个抽象的类型，只要实现了say()这个方法都可以成为sayer类型
type sayer interface {
	say()
}

type Cat struct{}

func (c Cat) say() {
	fmt.Println("喵喵喵~")
}

type Dog struct{}

func (d Dog) say() {
	fmt.Println("汪汪汪~")
}

type person struct {
	name string
}

func (p person) say() {
	fmt.Println("啊啊啊~")
}
