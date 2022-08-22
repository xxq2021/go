package main

import (
	"fmt"
)

func main() {
	// 常量声明 编译前知道的值
	// const a = 50
	// fmt.Println(a)

	// 声明一组常量
	// const (
	// 	name    = "anne"
	// 	age     = 50
	// 	country = "Canada"
	// )
	// age = 13 // 不允许
	// fmt.Println(name, age, country)

	// var a = math.Sqrt(4)
	// const b = math.Sqrt(4) // not allowed
	// fmt.Println(a, b)

	// string

	// var hello = "Hello World"
	// fmt.Println(hello)

	// const n = "Sam"
	// var name = n
	// fmt.Printf("type: %T value: %v", name, name)
	// fmt.Printf("\nsize: %T", unsafe.Sizeof(name))

	// const typedhello string = "Hello World"

	// fmt.Printf("type: %T value: %v", typedhello, typedhello)

	// fmt.Printf("\nsize: %T", unsafe.Sizeof(typedhello))

	// boolean

	// const trueConst = true
	// type myBool bool
	// var defaultBool = trueConst
	// var customBool myBool = trueConst
	// fmt.Printf("defaultBool:%T,customBool:%T", defaultBool, customBool)

	// const a = 5
	// fmt.Printf("a:%T,a:%T", a, unsafe.Sizeof(a))
	// len() 函数来获取切片、字符串、通道、map等的长度
	var name = "anne"
	fmt.Println("len of the name is", len(name))

	// 报错
	// var x = 10
	// fmt.Println("len of the name is", len(x)) // invalid argument

}
