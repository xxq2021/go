package main

import "fmt"

func main() {
	// 布尔值 默认值 false
	var a bool
	fmt.Println(a) // false
	a = true
	fmt.Println(a) // true

	// 无法进行数值运算，无法与其他类型进行转换

	// var b bool
	// c := b + 1
	// fmt.Println(c) // mismatched types bool and untyped int

}
