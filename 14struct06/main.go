package main

import (
	"encoding/json"
	"fmt"
)

// 结构体字段的可见性与JSON序列化

// 如果一个go语言包中定义的标识符是首字母答谢的，那么就是对外可见的
func main() {
	// 创建一个班级变量c1
	c1 := Class{
		Title:    "火箭班",
		Students: make([]Student, 0, 20),
	}
	for i := 0; i < 10; i++ {
		// 添加10个学生
		temp := newStudent(i+1, fmt.Sprintf("stu%02d", i+1))
		c1.Students = append(c1.Students, temp)
	}
	fmt.Printf("原始：%#v\n\n", c1)

	// JSON序列化：go语言的数据 -> JSON格式的字符串
	c1Data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json Marshal err: ", err)
		return
	}
	fmt.Printf("c1Data type：%T\n\n", c1Data)
	fmt.Printf("c1Data：%s\n\n", c1Data)
	// JSON反序列化：JSON格式的字符串 -> go语言的数据
	// jsonStr := "{\"Title\":\"火箭班\",\"Students\":[{\"Id\":1,\"Name\":\"stu01\"},{\"Id\":2,\"Name\":\"02\"}]}"
	jsonStr := "{\"title\":\"火箭班\",\"stu_list\":[{\"Id\":1,\"Name\":\"stu01\"},{\"Id\":2,\"Name\":\"02\"}]}"
	var c2 Class
	err = json.Unmarshal([]byte(jsonStr), &c2)
	if err != nil {
		fmt.Println("json Unmarshal err: ", err)
		return
	}
	fmt.Printf("c2 type：%T\n\n", c2)
	fmt.Printf("c2：%#v\n\n", c2)

}

// Student 学生
type Student struct {
	// Id、Name对外可见
	Id   int
	Name string
	// age对外不可见
	// age int
}

func newStudent(id int, name string) Student {
	return Student{
		Id:   id,
		Name: name,
	}
}

// Class 班级
type Class struct {
	Title    string    `json:"title" db:"title" xml:"title"`
	Students []Student `json:"stu_list"`
}
