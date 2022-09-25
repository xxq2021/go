package main

import "fmt"

func funcA(a, b int) int {
	return a / b
}

// 接入panic
func funcB(a, b int) int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err:", err)
		}
	}()
	return a / b
}
func main() {
	/* panic/recover 用来处理错误
	panic 可在任何地方引发，recover 只有在 defer调用的函数中有效
	*/

	// funcA(10, 0) // panic: runtime error: integer divide by zero
	funcB(10, 0) // err: runtime error: integer divide by zero
	fmt.Println(funcB(10, 2))

}
