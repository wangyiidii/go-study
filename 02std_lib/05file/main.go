package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 文件操作

func main() {
	//readFromFile()
	//readAll()
	//readByBufIO()
	//readByIOUtil()

	//write()
	//writeByBufIO()
	writeByIOUtil()
}

func readFromFile() {
	file, err := os.Open("./02std_lib/05file/1.txt") // 相对路径，相对执行的程序同目录下的1.txt
	if err != nil {
		fmt.Printf("open file err: %s", err)
		return
	}
	// 关闭文件
	defer file.Close()

	// 读取文件内容
	tmp := make([]byte, 128)
	n, err := file.Read(tmp)
	if err != nil {
		fmt.Printf("read file err: %s", err)
		return
	}
	fmt.Printf("read %d bytes from file. \n", n)
	fmt.Println(string(tmp))
}

func readAll() {
	file, err := os.Open("./02std_lib/05file/1.txt") // 相对路径，相对执行的程序同目录下的1.txt
	if err != nil {
		fmt.Printf("open file err: %s", err)
		return
	}
	// 关闭文件
	defer file.Close()

	// 读取文件内容
	for {
		tmp := make([]byte, 128)
		n, err := file.Read(tmp)
		if err == io.EOF {
			// 把当前读了多少个字节的数据打印出来，然后退出
			fmt.Println(string(tmp[:n]))
			return
		}
		if err != nil {
			fmt.Printf("read file err: %s \n", err)
			return
		}
		fmt.Printf("read %d bytes from file. \n", n)
		fmt.Println(string(tmp[:n]))
	}
}

func readByBufIO() {
	file, err := os.Open("./02std_lib/05file/1.txt") // 相对路径，相对执行的程序同目录下的1.txt
	if err != nil {
		fmt.Printf("open file err: %s", err)
		return
	}
	// 关闭文件
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Print(line)
			return
		}
		if err != nil {
			fmt.Printf("read file by bufio err: %s", err)
			return
		}
		fmt.Print(line)
	}
}

func readByIOUtil() {
	content, err := ioutil.ReadFile("./02std_lib/05file/1.txt")
	if err != nil {
		fmt.Printf(" ioutil ReadFile err: %s", err)
		return
	}
	fmt.Println(string(content))
}

func write() {
	file, err := os.OpenFile("./02std_lib/05file/2.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644)
	if err != nil {
		fmt.Printf("OpenFile err: %s\n", err)
		return
	}
	defer file.Close()

	str := "第一行"
	file.Write([]byte(str))
	file.WriteString("第二行")
}

func writeByBufIO() {
	file, err := os.OpenFile("./02std_lib/05file/3.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644)
	if err != nil {
		fmt.Printf("OpenFile err: %s\n", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	// 将内容写入缓冲区
	writer.WriteString("writeByBufIO ")
	// 将缓冲区的内容写入磁盘
	writer.Flush()
}

func writeByIOUtil() {
	str := "writeByIOUtil "
	err := ioutil.WriteFile("./02std_lib/05file/4.txt", []byte(str), 0644)
	if err != nil {
		fmt.Printf("writeByIOUtil WriteFile err: %s\n", err)
		return
	}

}
