package main

import (
	"errors"
	"fmt"
)

// 模拟了一个读取文件的方法
func readFile(fileName string) error {
	// 返回 error 类型
	if fileName == "main.go" {
		return nil // 表示存在
	} else {
		return errors.New("读取文件失败")
	}
}

func myFn() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("给管理员发送邮件")
		}
	}()
	err := readFile("xxx.go")
	if err != nil {
		panic(err)
	}
}

func main() {
	myFn()
	fmt.Println("Go on...")
}
