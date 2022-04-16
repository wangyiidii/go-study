package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func main() {
	// 声明但是不初始化，map就是零值nil
	var a map[string]int
	fmt.Println(a == nil)
	// map的初始化
	a1 := make(map[string]int, 8)
	fmt.Println(a1 == nil)
	// map添加键值对
	a1["a"] = 1
	a1["b"] = 2
	fmt.Println(a1)
	fmt.Printf("a1: %#v\n", a1)
	fmt.Printf("type: %T\n", a1)
	// 声明map的同时完成初始化
	a2 := map[string]int{
		"a": 2,
		"b": 3,
	}
	fmt.Println(a2)
	fmt.Printf("a2: %#v\n", a2)
	fmt.Printf("type: %T\n", a2)
	// map必须初始化才能使用
	//var a3 map[int]int
	//a3[1] = 1
	//fmt.Println(a3)
	// 判断某个键是否存在
	a4 := make(map[string]string, 20)
	a4["a"] = "11"
	a4["b"] = "12"
	a4["c"] = "13"
	v, ok := a4["a"]
	fmt.Println(v)
	fmt.Printf("value: %s, exist: %t\n", v, ok)

	// for range map是无序的
	for k, v := range a4 {
		fmt.Printf("k: %s, v: %s\n", k, v)
	}
	// 只遍历key
	for k := range a4 {
		fmt.Printf("k: %s\n", k)
	}
	// 只遍历value
	for _, v := range a4 {
		fmt.Printf("v: %s\n", v)
	}

	// map 删除
	a5 := make(map[string]string, 20)
	a5["a"] = "11"
	a5["b"] = "12"
	a5["c"] = "13"
	a5["d"] = "14"
	fmt.Println(a5)
	delete(a5, "b")
	fmt.Println(a5)

	// 按照某个顺序遍历map
	var a6 = make(map[string]int, 100)
	for i := 0; i < 50; i++ {
		k := fmt.Sprintf("%02d", i)
		v := rand.Intn(100) // 0-99的随机数
		a6[k] = v
	}
	for k, v := range a6 {
		fmt.Printf("k: %s, v: %d\n", k, v)
	}
	// 按照key从小到大的顺序排序
	keys := make([]string, 0, 100)
	for k, _ := range a6 {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("sort k: %s, v: %d\n", k, a6[k])
	}

	// 元素类型为map的切片
	mapSlice := make([]map[string]string, 8, 10) // 只完成了slice的初始化
	mapSlice[0] = make(map[string]string, 5)
	mapSlice[0]["a"] = "1"
	fmt.Println(mapSlice)

	// 值为切片的map
	sliceMap := make(map[string][]string, 8) // 只完成了map的初始化
	sliceMap["b"] = make([]string, 8, 8)
	sliceMap["b"][0] = "b1"
	sliceMap["b"][1] = "b2"
	fmt.Println(sliceMap)

	// 练习题
	// 写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。
	s1 := "how do you do"
	m1 := make(map[string]int, 10)
	for _, v := range strings.Split(s1, " ") {
		m1[v] = m1[v] + 1
	}
	fmt.Println(m1)

	//
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}
