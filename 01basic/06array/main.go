package main

import "fmt"

// 数组
func main() {
	// 1.定义数组
	var a [3]int
	var b [4]int
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	// 数组初始化
	// 1.1 定义时使用初始值列表的方式初始化
	cityArr := [4]string{"a", "b", "c", "d"}
	fmt.Println(cityArr)
	// 1.2 编译器推导数组的长度
	boolArr := [...]bool{false, true, false, true, true}
	fmt.Println(boolArr)
	// 1.3 使用索引值方法初始化
	longArr := [...]string{1: "go", 3: "java", 7: "python"}
	fmt.Println(longArr)
	fmt.Println(len(longArr))
	fmt.Printf("%T\n", longArr)

	// 2. 数组遍历
	// 2.1 fori
	for i := 0; i < len(cityArr); i++ {
		fmt.Println(cityArr[i])
	}
	// 2.2 for range
	for _, v := range cityArr {
		fmt.Println(v)
	}

	// 3. 二维数组
	// 3.1 外层的可以用...
	//arr2 := [3][2]string{
	arr2 := [...][2]string{
		{"11", "12"},
		{"21", "22"},
		{"31", "32"},
	}
	fmt.Println(arr2)
	fmt.Println(arr2[1][1])
	// 3.2 遍历
	for _, v1 := range arr2 {
		//01fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	// 3.3 数组时值类型
	x := [3]int{1, 2, 3}
	fmt.Println(x)
	f1(a)
	fmt.Println(x)

	// 练习题
	// 求数组[1, 3, 5, 7, 8]所有元素的和
	// 找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。

	arr := [...]int{1, 3, 5, 7, 8}
	count := 0
	for _, v := range arr {
		count += v
	}
	fmt.Println(count)

	for i, v1 := range arr {
		for j := i; j < len(arr); j++ {

			if v1+arr[j] == 8 {
				fmt.Printf("(%d, %d)\t", i, j)
			}
		}
	}
}

func f1(a [3]int) {
	a[0] = 100
}
