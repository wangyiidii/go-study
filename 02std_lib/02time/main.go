package main

import (
	"fmt"
	time "time"
)

// time demo

func main() {
	now := time.Now()
	fmt.Println(now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Println(year, month, day, hour, minute, second)

	// 时间戳
	unixSecond := now.Unix()           //秒
	unixMicroSecond := now.UnixMicro() //纳秒
	fmt.Println(unixSecond, unixMicroSecond)

	// 将时间转换为具体的时间格式
	// 1650183499 + 3600
	t := time.Unix(1650187099, 0)
	fmt.Println(t)

	// 时间间隔 Duration (是一个基于int64的类型)
	//time.Duration()
	n := 1
	time.Sleep(time.Duration(n) * time.Second)

	// 时间操作
	// Add
	fmt.Println("时间操作")
	fmt.Println("now: ", now)
	t2 := now.Add(time.Hour)
	fmt.Println("t2: ", t2)
	// Sub
	// Equal
	// Before
	// After

	// 定时器
	//fmt.Println("定时器")
	//tick := time.Tick(time.Second)
	//for tmp := range tick {
	//	fmt.Println(tmp)
	//}

	// 时间格式化
	/*
		Y		m	s	H	H	S
		2006	01mysql	02	15	04	05
	*/
	fmt.Println("时间格式化")
	format := now.Format("2006年01月02日 15:04:05.000")
	fmt.Println(format)
	format2 := now.Format("2006年01月02日 03:04:05 PM")
	fmt.Println(format2)

	// 解析字符串格式的时间
	fmt.Println("解析字符串格式的时间")
	timeStr := "2022-04-17 16:33:02"
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	//parse, err := time.Parse("2016-01mysql-02 15-04-05", timeStr)
	parse, err := time.ParseInLocation("2006-01mysql-02 15:04:05", timeStr, location)
	if err != nil {
		fmt.Printf("time parse err: %v", err)
		return
	}
	fmt.Println(parse)

	// 练习题
	// 1. 获取当前时间，格式化输出为2017/06/19 20:30:05格式。
	fmt.Println("练习题")
	now1 := time.Now()
	format1 := now1.Format("2006/01mysql/02 15:04:05")
	fmt.Println(format1)

	// 2. 编写程序统计一段代码的执行耗时时间，单位精确到微秒。
	start := time.Millisecond
	time.Sleep(time.Second * 2)
	fmt.Printf("耗时: %sms", time.Millisecond-start)
}
