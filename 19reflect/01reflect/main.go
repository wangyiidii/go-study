package main

import (
	"fmt"
	"reflect"
)

/*
Invalid Kind = iota  // 非法类型
    Bool                 // 布尔型
    Int                  // 有符号整型
    Int8                 // 有符号8位整型
    Int16                // 有符号16位整型
    Int32                // 有符号32位整型
    Int64                // 有符号64位整型
    Uint                 // 无符号整型
    Uint8                // 无符号8位整型
    Uint16               // 无符号16位整型
    Uint32               // 无符号32位整型
    Uint64               // 无符号64位整型
    Uintptr              // 指针
    Float32              // 单精度浮点数
    Float64              // 双精度浮点数
    Complex64            // 64位复数类型
    Complex128           // 128位复数类型
    Array                // 数组
    Chan                 // 通道
    Func                 // 函数
    Interface            // 接口
    Map                  // 映射
    Ptr                  // 指针
    Slice                // 切片
    String               // 字符串
    Struct               // 结构体
    UnsafePointer        // 底层指针
*/

func main() {
	var a float32 = 3.14
	reflectType(a) // type:float32
	var b int64 = 100
	reflectType(b) // type:int64

	// 结构体类型
	var c Cat
	reflectType(c)
	var d Dog
	reflectType(d)

	// slice
	var slice []int
	reflectType(slice)
	// string
	var str string
	reflectType(str)

	var a1 int32 = 100
	reflectValue(a1)

	var a2 int32 = 99
	fmt.Printf("reflectSetValue pre v: %v, type: %T\n\n", a2, a2)
	reflectSetValue(&a2)
	fmt.Printf("reflectSetValue post v: %v, type: %T\n\n", a2, a2)

	// isNil(): IsNil()报告v持有的值是否为nil。v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一；否则IsNil函数会导致panic。

	// isValid(): IsValid()返回v是否持有一个值。如果v是Value零值会返回假，此时v除了IsValid、String、Kind之外的方法都会导致panic。

}

func reflectType(x interface{}) {
	// 不知道调用该方法的参数类型
	// 1. 通过类型断言
	// 2. 借助反射
	v := reflect.TypeOf(x)
	fmt.Println("obj: ", v)
	fmt.Println("Name: ", v.Name())
	fmt.Println("Kind: ", v.Kind())
	fmt.Printf("type: %v\n\n", v)
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("value: %v, type: %T\n\n", v, v)
	// 如何得到一个传入类型的变量
	switch v.Kind() {
	case reflect.Float32:
		ret := float32(v.Float())
		fmt.Printf("float32: %v, type: %T\n\n", ret, ret)
	case reflect.Int32:
		ret := int32(v.Int())
		fmt.Printf("Int32: %v, type: %T\n\n", ret, ret)
	}

}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	// Elem用来根据指针取对应的值
	k := v.Elem().Kind()
	switch k {
	case reflect.Int32:
		v.Elem().SetInt(101)
	}
}

type Cat struct {
}
type Dog struct {
}
