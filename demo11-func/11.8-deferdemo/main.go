package main

import "fmt"

func f1() int {
	// 匿名返回值
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f2() (x int) {
	// 返回 x 后执行defer语句， 最后返回 x++
	defer func() {
		x++
	}()
	return 5 // 6
}
func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // 5
}

func f4() (x int) {
	defer func(x int) {
		// x = 0
		fmt.Println(x)
		x++ // 改变的是 defer 函数内的 x
	}(x)
	fmt.Println("a =", x)
	return 5 // 5 defer 注册要执行延迟执行的函数时，该函数所有的参数都需要确定其值
}
func main() {
	//fmt.Println(f1())
	//fmt.Println(f2())
	//fmt.Println(f3())
	fmt.Println(f4())

}
