package main

func main() {
	//fmt.Errorf("查询数据库失败，err:%v", err)

}

type error interface {
	Error() string
}
