package main

import (
	"errors"
	"fmt"
)

func main() {
	e := errors.New("原始错误e")
	w := fmt.Errorf("wrap了一个错误%w", e)
	fmt.Println(w)

	var (
		name    string
		age     int
		married bool
	)
	fmt.Scan(&name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)

}
