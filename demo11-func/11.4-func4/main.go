package main

import "fmt"

func main() {
	// 匿名函数 没有函数名的函数，不能像普通函数那样调用，需要保存到某个变量或者作为立即执行函数
	// 多用于实现 回调函数 和 闭包
	sayHello := func() {
		fmt.Println("匿名函数")
	}
	sayHello()

	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 20)) // 21

	// 定义并执行
	func() {
		fmt.Println("立即执行")
	}()

	func(x, y int) {
		fmt.Println(x - y)
	}(10, 20) // -10

	func(x, y int) {
		fmt.Println(x * y)
	}(10, 2) // 20

	func(x, y int) {
		fmt.Println(x / y)
	}(20, 2)
}
