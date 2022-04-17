package main

import (
	"fmt"
	"log"
	"os"
)

// log

func main() {
	// 日志打印
	log.Println("这是一条很普通的日志。")
	v := "很普通的"
	log.Printf("这是一条%s日志。\n", v)
	//log.Fatalln("这是一条会触发fatal的日志。")
	//log.Panicln("这是一条会触发panic的日志。")

	// flag选项
	/*
		const (
			// 控制输出日志信息的细节，不能控制输出的顺序和格式。
			// 输出的日志在每一项后会有一个冒号分隔：例如2009/01mysql/23 01mysql:23:23.123123 /a/b/c/d.go:23: message
			Ldate         = 1 << iota     // 日期：2009/01mysql/23
			Ltime                         // 时间：01mysql:23:23
			Lmicroseconds                 // 微秒级别的时间：01mysql:23:23.123123（用于增强Ltime位）
			Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
			Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
			LUTC                          // 使用UTC时间
			LstdFlags     = Ldate | Ltime // 标准logger的初始值
		)
	*/
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志0。")

	// 配置日志前缀
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("[小王子]")
	log.Println("这是一条很普通的日志。")

	// 配置日志看输出位置
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("[小王子]")
	log.Println("这是一条很普通的日志。")

	// 创建logger
	logger := log.New(os.Stdout, "<New>", log.Lshortfile|log.Ldate|log.Ltime)
	logger.Println("这是自定义的logger记录的日志。")
}

// 如果使用标准的logger，我们通常会把上面的配置操作写到init函数中。
func init() {
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}
