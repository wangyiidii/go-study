package main

// 当代码都放在$GOPATH目录下的话
// 导入包就要从$GOPATH/src后面继续写起
// 但如宝不允许导入包而不使用！！！
// 不允许循环引用包
import (
	"fmt"
	c "go-study/16package/calc"
)

func main() {
	sum := c.Add(1, 2)
	fmt.Println(sum)
	fmt.Println(c.Name)
}

func init() {
	fmt.Println("main init")
}
