package main

import (
	"fmt"
	"os"
)

// 学员信息管理系统
// 1. 添加学员信息
// 2. 编辑学员信息
// 3. 展示所有学员信息

func main() {
	mgr := newStudentMgr()
	for {
		// 1. 打印系统菜单
		showMenu()
		// 2. 等待用户选择要执行的选项
		var input int
		fmt.Printf("请输入你要操作的序号:")
		fmt.Scanf("%d\n", &input)
		fmt.Println("用户输入的是：", input)
		// 3. 执行用户选择的动作
		switch input {
		case 1:
			stu := getInput()
			mgr.addStudent(stu)
		case 2:
			stu := getInput()
			mgr.addStudent(stu)
			mgr.modifyStudent(stu)
		case 3:
			mgr.listStudent()
		case 4:
			fmt.Println("系统退出")
			os.Exit(0)
		default:
			fmt.Println("指令输入错误，系统退出")
			os.Exit(0)
		}
	}

}

func showMenu() {
	fmt.Println("欢迎来到学员信息管理系统")
	fmt.Println("1. 添加学员")
	fmt.Println("2. 编辑学员信息")
	fmt.Println("3. 展示所有学员信息")
	fmt.Println("4. 退出")
}

// 获取用户输入的学员信息
func getInput() *student {
	var (
		id    int
		name  string
		class string
	)
	fmt.Println("请按要求输入学员信息")
	fmt.Println("请输入学员学号")
	fmt.Scanf("%d\n", &id)
	fmt.Println("请输入学员姓名")
	fmt.Scanf("%s\n", &name)
	fmt.Println("请输入学员班级")
	fmt.Scanf("%s\n", &class)

	stu := newStudent(id, name, class)
	return stu
}
