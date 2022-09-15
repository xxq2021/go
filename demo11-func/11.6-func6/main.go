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
/*
闭包是指有权访问另一个函数作用域中的变量的函数
创建闭包的常见方式是在一个函数内部创建另一个函数，通过另一个函数访问这个函数的局部变量
注意：由于闭包里作用域返回的局部变量资源不会被立刻销毁回收，所以可能会占用更多的内存
过度使用 闭包会导致性能下降，建议在非常有必要的时候才使用闭包
*/

//var a = 12 // 全局变量 只能用 var 定义
//func test() {
//	fmt.Println(a)
//}

// 闭包写法：函数里面嵌套一个函数，最后返回里面的函数
func adder() func() int {
	var i = 10
	return func() int {
		return i + 1
	}
}

func adder2() func(y int) int {
	var i = 10
	return func(y int) int {
		i += y
		return i
	}
}

func main() {
	//test()         // 12
	//fmt.Println(a) // 12

	var fn = adder()  // 表示执行方法
	fmt.Println(fn()) // 11
	fmt.Println(fn()) // 11
	fmt.Println(fn()) // 11

	var fn2 = adder2()
	fmt.Println(fn2(1)) // 11
	fmt.Println(fn2(1)) // 12
	fmt.Println(fn2(1)) // 13
}
