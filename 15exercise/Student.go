package main

import "fmt"

type student struct {
	id    int // 学号唯一
	name  string
	class string
}

// newStudent 学生的构造函数
func newStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

// 学员管理
type studentMgr struct {
	allStudents []*student
}

// 学员管理 构造函数
func newStudentMgr() *studentMgr {
	return &studentMgr{
		allStudents: make([]*student, 0, 100),
	}
}

// 添加学生
func (s *studentMgr) addStudent(stu *student) {
	s.allStudents = append(s.allStudents, stu)
}

// 编辑学生
func (s *studentMgr) modifyStudent(stuForm *student) {
	for i, stu := range s.allStudents {
		if stu.id == stuForm.id {
			s.allStudents[i] = stuForm
			return
		}
	}
	// 输入的学生没有找到
	fmt.Printf("输入学生信息有无，系统中没有学好%d的学生", stuForm.id)
}

// 展示所有学生
func (s *studentMgr) listStudent() {
	for _, stu := range s.allStudents {
		fmt.Printf("学号：%d，姓名：%s，班级：%s\n", stu.id, stu.name, stu.class)
	}
}
