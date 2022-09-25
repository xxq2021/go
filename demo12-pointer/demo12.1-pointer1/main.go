package main

import "fmt"

/* 每个变量在运行时都有一个地址，这个地址代表变量在内存中的值。
使用 & 对变量进行取地址操作，使用 * 根据变量的地址进行取值
Go 语言中的值类型(int float bool string array struct)都有对应的指针类型
*/

func main() {
	a := 10 // int类型
	b := &a // *int类型 存储的是变量 a 的内存地址
	fmt.Printf("a的值为:%d, a的类型为:%T, a的地址为:%p\n", a, a, &a)
	fmt.Println(b, &b, *b) // 0xc00000a0a8 0xc000006028 10
	c := *b                // 根据内存地址去取值
	fmt.Println(c)         // 10

}
