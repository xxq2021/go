package main

import "fmt"

func deferFn() {
	/*
		defer 语句会将其后面跟随的语句进行延迟处理 调用顺序类似于 栈
		先被 defer 的语句最后执行，后被 defer 的语句先执行
	*/
	fmt.Println("start")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("end")
}

func deferFn2() {
	fmt.Println("start")
	defer func() {
		fmt.Println("a")
		fmt.Println("b")
	}() // 必须执行
	fmt.Println("end")
}

// defer 在命名返回值和匿名函数中表现不一样

// 调用方法返回的值是 0
func fn() int {
	var a int // 0
	defer func() {
		a++
	}()
	return a
}

// 匿名函数
// 调用方法返回的值是 1
func fn2() (a int) {
	defer func() {
		a++
	}()
	return a
}

func main() {

	// deferFn()  // start end 3 2	1
	// deferFn2() // start end a b
	// fmt.Println(fn()) // 0
	fmt.Println(fn2()) // 1

}
