package main

import "fmt"

func main() {
	// 数组的定义
	// 数组长度是类型的一部分。

	// var a [3]string
	// var b [3]int
	// var c [4]bool
	// fmt.Printf("a:%T,b:%T,c:%T", a, b, c) // a:[3]string,b:[3]int,c:[4]bool

	// 数组的初始化
	// 第一种方法
	// var a [3]int
	// fmt.Println(a) // [0 0 0]
	// a[0] = 1
	// a[1] = 2
	// a[2] = 3
	// fmt.Println(a) // [1 2 3]

	// var b [2]string
	// fmt.Println(b) // [ ]
	// b[0] = "hello"
	// b[1] = "world"
	// b[2] = "golang" //  index 2 out of bounds [0:2]
	// fmt.Println(b) // [hello world]

	// 第二种方法

	// var a = [3]int{1, 2, 3}
	// fmt.Println(a) // [1 2 3]

	// var b = [3]string{"Hello", "Golang", "!"}
	// fmt.Println(b) // [Hello Golang !]

	// 第三种方法
	// c := [4]bool{false, true, true, false}
	// fmt.Println(c) // [false true true false]

	// 第四种方法
	// 让编译器根据初始值的个数进行自行判断数组的长度

	// var a = [...]string{"北京", "上海", "广州", "深圳"}
	// fmt.Printf("a: %T", a) // a: [4]string
	// fmt.Println(a)         // [北京 上海 广州 深圳]

	// 查看数组长度

	// fmt.Println(len(a)) // 4

	// 根据下标改变数组里的值

	// a[0] = "苏州"
	// fmt.Println(a) // [苏州 上海 广州 深圳]

	// 第四种初始化方式
	// 指定索引值

	var b = [...]int{1: 1, 2: 2, 5: 3}
	fmt.Println(b)         // [0 1 2 0 0 3]
	fmt.Printf("b: %T", b) // b: [6]int

}
