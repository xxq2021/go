package main

import (
	"fmt"
)

func main() {
	// 字符 组成字符串
	// byte => unit8的别名 ASCII码
	// rune => int32的别名 中文、日文、韩文...

	// var c1 byte = 'a'
	// var c2 rune = 'c'
	// fmt.Println(c1, c2)                      // 97 99
	// fmt.Printf("c1的值是%v, c2的值是%v", c1, c2)   // 97 99
	// fmt.Printf("c1的类型是%T, c2的类型是%T", c1, c2) // uint8 int32
	// fmt.Printf("c1=%c, c2=%c", c1, c2)

	// s := "hello, 你好"
	// for i := 0; i < len(s); i++ {
	// 	fmt.Printf("%c", s[i]) // hello, ä½ å¥½
	// }

	// s := "hello, 你好"
	// for _, r := range s {
	// 	fmt.Printf("%c", r) //hello, 你好
	// }

	// s := "你好"
	// str := "this"
	// fmt.Println(len(s), len(str)) // 6 4
	// //unsafe.Sizeof()无法查看
	// fmt.Println(unsafe.Sizeof(s), unsafe.Sizeof(str)) // 16 16

	/* x := '你'
	y := '好'
	fmt.Printf("x的值是%v, y的值是%v", x, y)   // 20320 22909
	fmt.Printf("x的类型是%T, y的类型是%T", x, y) // int32 int32 */

	// 字符串修改
	// 无法直接修改

	s1 := "big"
	// 强制类型转化
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1)) // pig
	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2)) // 红萝卜

}
