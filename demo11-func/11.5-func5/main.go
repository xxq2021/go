package main

import "fmt"

// 递归函数
// func fn1() {
// 	fmt.Println("方法1")
// }

// func fn2() {
// 	fn1()
// 	fmt.Println("方法2")
// }

func f3(x int) {
	// 2. 传入 1 个整数，递归打印出 1-这个数之内的所有整数
	if x > 0 {
		fmt.Println(x)
		x--
		f3(x)
	}
}

var sum = 0

func f4(x int) int {
	// 递归实现 1 - 100 的和
	if x > 0 { // 100
		sum += x // 100
		x--      // 99
		f4(x)    // 99
	}
	return sum
}

func f5(x int) int {
	if x > 1 {
		return x + f5(x-1)
	} else {
		return 1
	}

}

// 递归实现 5 的阶乘
func f6(x int) int {
	if x > 1 {
		return x * f6(x-1)
	} else {
		return 1
	}
}

func main() {
	// fn2()

	// 1. for 循环实现 1-100 的和
	// sum := 0
	// for i := 0; i <= 100; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	// 2. 传入 1 个整数，递归打印出 1-这个数之内的所有整数
	// f3(5)

	// 递归实现 1 - 100 的和
	// fmt.Println(f4(100))
	// fmt.Println(f5(100))
	fmt.Println(f6(5))

}
