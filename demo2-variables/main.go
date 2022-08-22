package main

import (
	"fmt"
	// "math"
)

var gloableVariable = 10

func main() {
	// 变量是为存储特定类型的值的内存位置所指定的名称。
	// 变量声明

	// 单变量声明格式： var 变量名 变量类型
	// 这种命名方式没有为变量分配值 => Go默认会为变量分配初始值即变量的初始化
	// 字符串的默认值为空字符串，整形的默认值为0
	// var name string
	// var age int

	// fmt.Println(name, age) // 打印空字符串，0

	// 变量赋值
	// age = 18
	// fmt.Println("my age is", age) // my age is 18

	// 声明一个有初始值的变量 var 变量名 变量类型 = 初始值
	// var age int = 18
	// fmt.Println("My age is", age)

	// 类型推导 var 变量名 = 初始值
	// 如果变量有初始值，Go能根据变量的值自动推断出变量的类型

	// var age = 18
	// fmt.Println("my age is", age)

	// 多变量声明
	// 单个语句声明多个类型相同变量 var 变量1, 变量2 变量类型 = 初始值, 初始值
	// var width, height int = 100, 50
	// fmt.Println("width is", width, "height is", height)
	// 省略变量类型 var 变量1, 变量2 = 初始值, 初始值

	// var width, height = 100, 50
	// fmt.Println("width is", width, "height is", height)

	// 单个语句声明多个类型不相同的变量
	// var (
	// 	name   = "anne"
	// 	age    = 18
	// 	height = 80
	// )
	// fmt.Println("my name is", name, "my age is", age, "my height is", height)

	// 短变量声明 必须赋值
	// count := 10
	// fmt.Println("count is", count)

	// name, age, height := "anne", 18, 80
	// fmt.Println("my name is", name, "my age is", age, "my height is", height)

	// 仅当 := 左侧的至少一个变量是新声明的时，才能使用简写语法。

	// a, b := 20, 30
	// fmt.Println("a is", a, "b is", b)
	// b, c := 40, 50
	// fmt.Println("b is", b, "c is", c)
	// b, c = 60, 70
	// fmt.Println("now b is", b, "now c is", c)

	// no new variables on left side of :=
	// a, b := 20, 30
	// fmt.Println("a is", a, "b is", b)
	// a, b := 40, 50
	// fmt.Println("a is", a, "b is", b)

	// a, b := 145.8, 543.8
	// c := math.Min(a, b)
	// fmt.Println("Minimum value is", c)

	// 短变量声明 只能在函数内声明（局部变量），而var可以在任何位置创建变量（全局变量）

	localVariable := 20
	fmt.Println("gloableVariable is", gloableVariable, "localVariable is", localVariable) // 10 20

}
