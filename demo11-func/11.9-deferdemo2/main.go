package main

import "fmt"

//defer 注册要执行延迟执行的函数时，该函数所有的参数都需要确定其值
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}

/* 代码运行分析
注册顺序：
	1、	defer calc("AA", x, calc("A", x, y))
	2、	defer calc("BB", x, calc("B", x, y))
执行顺序：
	1、defer calc("BB", x, calc("B", x, y))
	2、defer calc("AA", x, calc("A", x, y))

*/
// 1、 calc("A", x, y)  A 1 2 3
// 2、 calc("B", x, y)  B 10 2 12
// 3、 calc("BB", x, calc("B", x, y)) BB 10 12 22
// 4、 calc("AA", x, calc("B", x, y)) AA 1 3 4
