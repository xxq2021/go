package main

import "fmt"

func main() {

	// byte (uint8) rune (int32)
	// 直接输出byte时，输出的是这个字符对于的ASCII码值

	// var c1 byte = 'a'
	// var c2 rune = 'c'
	// fmt.Println(c1, c2)         // 97 99
	// fmt.Printf("%v %v", c1, c2) // 97 99
	// fmt.Printf("%T %T", c1, c2) // uint8 int32

	// 原样输出 - %c格式化输出
	// fmt.Printf("%c %c", c1, c2) // a c

	// 汉字占3个字节 打印输出的是Unicode编码10进制
	// s0 := "你好"
	// fmt.Println(len(s0)) // 6
	// s1 := '你'
	// s2 := '好'
	// fmt.Println(s1, s2)         // 20320 22909
	// fmt.Printf("%T %T", s1, s2) // int32 int32

	// 循环输出字符串中的字符 for range

	// byte会乱码
	// s3 := "Hello, 你好"

	// for i := 0; i < len(s3); i++ {
	// 	fmt.Printf("%c", s3[i]) // Hello, ä½ å¥½
	// }

	// 使用 rune for range
	// for _, r := range s3 {
	// 	fmt.Printf("%c", r)
	// }

	// 修改字符串
	// 强制类型转化
	// byte类型
	s4 := "big"
	byteS4 := []byte(s4)
	byteS4[0] = 'p'
	fmt.Println(byteS4)         // [112 105 103]
	fmt.Println(string(byteS4)) // pig

	// rune类型
	s5 := "白萝卜"
	runeS5 := []rune(s5)
	runeS5[0] = '红'
	fmt.Println(runeS5) // [32418 33821 21340]

	fmt.Println(string(runeS5)) // 红萝卜

}
