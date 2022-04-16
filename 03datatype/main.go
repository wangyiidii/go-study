package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// 整型
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a) // 10
	fmt.Printf("%b \n", a) // 1010  占位符%b表示二进制

	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b) // 77

	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) // ff
	fmt.Printf("%X \n", c) // FF

	// 浮点型
	fmt.Printf("pi： %f\n", math.Pi)
	fmt.Printf("pi .2：%.2f\n", math.Pi)

	// 复数 complex
	var c1 complex64 = 1 + 2i
	var c2 complex128 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)

	// 布尔

	// 字符串
	s1 := "s1"
	fmt.Printf("s1: %s\n", s1)
	// 反引号间换行将被作为字符串中的换行，但是所有的转义字符均无效，文本将会原样输出。
	s2 := `hello
my name is tom
			i am 18 years \told`
	fmt.Printf("s1:\n %s\n", s2)

	/*
		len(str)								求长度
		+或fmt.Sprintf							拼接字符串
		strings.Split							分割
		strings.contains						判断是否包含
		strings.HasPrefix,strings.HasSuffix		前缀/后缀判断
		strings.Index(),strings.LastIndex()		子串出现的位置
		strings.Join(a[]string, sep string)		join操作
	*/
	s3 := "java,vue,go,python"
	fmt.Println("=========")
	fmt.Printf("len: %d\n", len(s3))
	s4 := fmt.Sprintf("%s,%s,%s", s3, "shell", "c++")
	fmt.Printf("s4: : %s\n", s4)
	fmt.Println(strings.Split(s3, ","))
	fmt.Println(strings.Contains(s3, ","))
	fmt.Printf("index: %d\n", strings.Index(s3, "java"))
	fmt.Println(strings.Join(strings.Split(s3, ","), "~"))

	// byte和rune类型
	/*
		Go 语言的字符有以下两种：

		uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
		rune类型，代表一个 UTF-8字符。
		当需要处理中文、日文或者其他复合字符时，则需要用到rune类型。rune类型实际是一个int32。

		Go 使用了特殊的 rune 类型来处理 Unicode，让基于 Unicode 的文本处理更为方便，也可以使用 byte 型进行默认字符串处理，性能和扩展性都有照顾。
	*/
	b1 := '中'
	fmt.Println(b1)

	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()

	// 修改字符串
	s11 := "big"
	// 强制类型转换
	byteS1 := []byte(s11)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s22 := "白萝卜"
	runeS2 := []rune(s22)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))

	// 类型转换
	/*
		Go语言中只有强制类型转换，没有隐式类型转换。该语法只能在两个类型之间支持相互转换的时候使用：T(表达式)
	*/
	var aa, bb = 3, 4
	var cc int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	cc = int(math.Sqrt(float64(aa*aa + bb*bb)))
	fmt.Println(cc)

	/*
		1. 编写代码分别定义一个整型、浮点型、布尔型、字符串型变量，使用fmt.Printf()搭配%T分别打印出上述变量的值和类型。
		2. 编写代码统计出字符串"hello沙河小王子"中汉字的数量。
	*/
	// 1. %T
	n1 := 123
	fmt.Printf("%T\n", n1)

	// 2.
	test2 := "hello沙河小王子"
	count := 0
	for _, t := range test2 {
		if len(string(t)) >= 3 {
			count++
		}
	}
	fmt.Println("count: ", count)

}
