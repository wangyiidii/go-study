package main

import "fmt"

// 结构体
// 结构体是值类型
func main() {
	// 结构体
	var p1 person
	p1.name = "tom"
	p1.city = "北京"
	p1.age = 18
	fmt.Printf("%#v\n", p1)

	// 匿名结构体
	var user struct {
		name    string
		married bool
	}
	user.name = "jerry"
	user.married = false
	fmt.Printf("%#v\n", user)

	// 创建指针类型的结构体
	p2 := new(person)
	fmt.Printf("p2: type: %T\n", p2)
	//(*p2).name = "p2"
	//(*p2).city = "上海"
	//(*p2).age = 19
	p2.name = "p2"
	p2.city = "上海"
	p2.age = 19
	fmt.Printf("%#v\n", p2)

	// 取结构体的地址进行实例化
	p3 := &person{}
	fmt.Printf("%#v\n", p3)
	p3.name = "p3"
	p3.city = "city3"
	p3.age = 20
	fmt.Printf("%#v\n", p3)

	// 结构体的初始化
	// 1. 使用键值对
	p4 := person{
		name: "p4",
		city: "city4",
		age:  20,
	}
	fmt.Printf("%#v\n", p4)
	p5 := &person{
		name: "p5",
		city: "city5",
		age:  21,
	}
	fmt.Printf("%#v\n", p5)
	// 2. 值的列表进行初始化
	p6 := person{
		"p6", "city6", 22,
	}
	fmt.Printf("%#v\n", p6)
	p7 := &person{
		"p7", "city7", 23,
	}
	fmt.Printf("%#v\n", p7)

	// 结构体的内存布局
	// 结构体占用一块连续的内存
	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)

	// 构造函数
	p8 := newPerson("p8", "city8", 23)
	fmt.Printf("%#v\n", p8)
}

//type person struct {
//	name string
//	age int
//	city string
//}

type person struct {
	name, city string
	age        int8
}

type test struct {
	a int8
	b int8
	c int8
	d int8
}

func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}

}
