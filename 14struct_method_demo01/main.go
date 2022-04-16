package main

import "fmt"

// 方法的定义实例 (java里面的对象方法)

func main() {
	p1 := newPerson("tom", 18)
	(*p1).Dream()
	p1.Dream()

	// 调用修改年龄的方法
	fmt.Printf("p1: %#v\n", p1)
	p1.SetAge(19)
	fmt.Printf("p1: %#v\n", p1)
	p1.SetAge2(17)
	fmt.Printf("p1: %#v\n", p1)

}

// Person 结构体
type Person struct {
	name string
	age  int8
}

// newPerson person类型的构造函数
func newPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// Dream 定义方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是go语言\n", p.name)
}

// SetAge 是一个修改年龄的方法（指针类型接收）
// 指针接收者指的就是接收者的类型是指针
func (p *Person) SetAge(age int8) {
	p.age = age
}

// SetAge2 是一个修改年龄的方法（值类型接收）
func (p Person) SetAge2(age int8) {
	p.age = age
}
