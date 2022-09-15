package main

import "fmt"

// Go 闭包
/*
全局变量特点:
	1. 常驻内存
	2. 污染全局
局部变量特点:
	1. 不常驻内存
	2. 不污染全局
闭包：
	1. 可以让一个变量常驻内存
	2. 可以让一个变量不污染全局
*/

var a = 12 // 全局变量 只能用 var 定义
func test() {
	fmt.Println(a)
}

func main() {
	test()         // 12
	fmt.Println(a) // 12
}
