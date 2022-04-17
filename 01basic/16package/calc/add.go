package calc

import (
	"fmt"
	_ "go-study/01basic/16package/snow"
)

// 标识符首字母大写表示对外可见
// 通常不对外的标识符都是要首字母小写

var Name = "calc"

func Add(a, b int) (sum int) {
	sum = a + b
	return
}

// init函数在包导入的时候自动执行
// init函数没有参数也没有返回值
func init() {
	fmt.Println("calc init...")
	fmt.Println("init Name:", Name)
}
