package main

import (
	"fmt"
)

// func sayHello() {
// 	fmt.Println("hello")
// }

// func sayHello(name string) string {
// 	return name
// }

// func sayHello(name string) {
// 	fmt.Println("hello", name)
// }

// func intSum(x int, y int) int {
// 	res := x + y
// 	return res
// }

// func intSum2(x int, y int) (res2 int) {
// 	res2 = x + y
// 	return
// }

// Go 语言中函数参数类型简写
// func intSum(x, y int) (ret int) {
// 	ret = x + y
// 	return
// }

// 函数可以接收可变参数，可变参数是指函数的参数数量不固定，在参数后面加... 表示可变参数
// 可变参数在函数体内是切片类型

// func intSum3(a ...int) {
// 	// fmt.Println(a)           // []
// 	// fmt.Printf("a: %T\n", a) // a: []int
// }

// func intSum4(a ...int) int {
// 	num := 0
// 	for _, v := range a {
// 		num += v
// 	}
// 	return num
// }

// 固定参数和可变参数同时出现时，可变参数要放在最后
// Go 语言中 没有默认参数
// func intSum5(a int, b ...int) int {
// 表示接收一个固定参数和一个可变参数

// 	ret1 := a
// 	for _, v1 := range b {
// 		ret1 += v1
// 	}
// 	return ret1
// }

// 定义多个返回值的参数
// func calc(a, b int) (int, int) {
// 	sum := a + b
// 	sub := a - b
// 	return sum, sub
// }

// 定义多个返回值的函数, 多个返回值也支持参数类型简写
func calc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

// defer语句
// 先被 defer的语句最后被执行，最后被defer的语句先执行 (栈)
// 由于 defer 语句具有延迟调用的特性，非常方便处理资源释放问题，
// 如：资源清理，文件关闭，解锁及记录时间等
func deferDemo() {
	fmt.Println("Start")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("End")

}
func main() {
	// var name string
	// fmt.Scan(&name)
	// sayHello(name)
	// fmt.Println("hello", x)
	// intSum(10, 20)
	// intSum2(102, 20)
	// intSum3()

	// fmt.Println(intSum4())           // 0
	// fmt.Println(intSum4(10))         // 10
	// fmt.Println(intSum4(10, 20))     // 20
	// fmt.Println(intSum4(10, 20, 30)) // 30

	// fmt.Println(intSum5(0))          // a = 0 b = []
	// fmt.Println(intSum5(10))         // a = 10 b = []
	// fmt.Println(intSum5(10, 20))     // a = 10 b = [20]
	// fmt.Println(intSum5(10, 20, 30)) // a = 10 b = [20 30]

	// a, b := calc(10, 20)
	// fmt.Println(a, b)

	fmt.Println(calc(10, 20)) // 30 -10

	deferDemo()

}
