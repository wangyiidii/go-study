package main

import (
	"fmt"
	"reflect"
)

/*
反射是把双刃剑
反射是一个强大并富有表现力的工具，能让我们写出更灵活的代码。但是反射不应该被滥用，原因有以下三个。

基于反射的代码是极其脆弱的，反射中的类型错误会在真正运行的时候才会引发panic，那很可能是在代码写完的很长时间之后。
大量使用反射的代码通常难以理解。
反射的性能低下，基于反射实现的代码通常比正常代码运行速度慢一到两个数量级。
*/
func main() {
	stu1 := student{
		Name:  "stu1",
		Score: 99,
	}
	// 通过反射去获取结构体的所有字段信息
	t := reflect.TypeOf(stu1)
	fmt.Printf("name: %v, kind: %v\n\n", t.Name(), t.Kind())

	for i := 0; i < t.NumField(); i++ {
		// 根据结构体字段的索引去取字段
		fieldObj := t.Field(i)
		fmt.Printf("name: %v, type: %v, tag: %v\n", fieldObj.Name, fieldObj.Type, fieldObj.Tag)
		fmt.Printf("json: %v\n", fieldObj.Tag.Get("json"))
		fmt.Printf("ini: %v\n", fieldObj.Tag.Get("ini"))
		fmt.Printf("66: %v\n", fieldObj.Tag.Get("66"))
	}

	// 根据名字去取结构体中的字段
	name, ok := t.FieldByName("Name")
	if ok {
		fmt.Println("name: ", name)
	}

	printMethod(&stu1)
}

type student struct {
	Name  string `json:"name" ini:"ini_name"`
	Score int    `json:"score" ini:"ini_score"`
}

// Study 给student添加两个方法 Study和Sleep(注意首字母大写)
func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

// 根据反射获取结构体中方法的函数
func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println("method count: ", t.NumMethod())

	// 因为下面需要拿到具体的方法取调用，所以使用的是值信息
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args []reflect.Value
		v.Method(i).Call(args)
	}

	// 通过方法名获取结构体的方法
	name, b := t.MethodByName("Sleep")
	fmt.Printf("t.MethodByName : name: %v, exist: %t\n", name, b)
	methodByName := v.MethodByName("Sleep")
	fmt.Printf("v.MethodByName : name: %v", methodByName)
}
