package main

import (
	"fmt"
)

// 函数进阶

// 函数作用域 - 全局变量 局部变量

// 定义全局变量
// var num = 10

// func testGlobal() {

// 可以在函数中使用变量
// 1. 先在自己的函数中查找，找到了就用自己函数中的，
// 2. 找不到就往外层函数中找全局变量
// 	num := 100
// 	fmt.Println("函数内变量", num)

// 	name := "fan"
// 	fmt.Println(name)

// }

// 函数作为参数传递到另一个函数中
func add(x, y int) int {
	return x + y
}

// 自定义一个方法类型
type calcType func(int, int) int

// func calc(x, y int, op func(int, int) int) int {
func calc(x, y int, op calcType) int {
	// op 是一个 func 函数类型的变量接收两个 int 类型的参数，返回值也是 int
	return op(x, y)
}

func sub(x, y int) int {
	return x - y
}

// 函数作为返回值
func do(o string) calcType {
	switch o {
	case "+":
		return add
	case "-":
		return sub
	case "*":
		return func(i1, i2 int) int { return i1 * i2 } // 返回匿名函数
	default:
		return nil
	}

}

func main() {
	// testGlobal()
	// fmt.Println(num)

	// 外层不能访问到函数内部的变量 （局部变量）
	// fmt.Println(name) // undefined: name

	// 外层不能访问到语句块的变量
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }
	// fmt.Println(i) //  undefined: i

	// 函数作为一等公民可以作为变量
	// abc := testGlobal

	// fmt.Printf("%T\n", abc) // func()
	// abc()                   // fan

	// 函数作为参数
	r := calc(10, 12, add)
	fmt.Println(r) // 22

	r2 := calc(10, 30, sub)
	fmt.Println(r2) // -20

	// 匿名函数 求两数的积
	r3 := calc(3, 5, func(i1, i2 int) int { return i1 * i2 })

	/*func calc(x, y int, op calcType) int {
		x = 3 y = 5 op =  func(i1, i2 int) int { return i1 * i2 }
		return op(x, y) x = 3 y = 5 => i1 = 3 i2 = 5 => return i1 * i2 = 15
	}*/
	fmt.Println(r3)

	b := do("+")
	// do("+") => return add => b := add()
	fmt.Println(b(10, 22)) // 32

	c := do("*")
	fmt.Println(c(22, 3)) // 66

}
