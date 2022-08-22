package main

import (
	"fmt"
	"strings"
)

func main() {
	// s1 := "hello world"
	// fmt.Printf("\ts1=%v", s1)
	// fmt.Println("c:\\code\\go.exe")
	// fmt.Println("\t制表符\n换行符")
	// 	s3 := `第一行
	// 第二行 \t
	// 第三行 \n
	// 	`
	// 	fmt.Println(s3)

	// s0 := "你好"
	// s1 := "hello"
	// s2 := "world"

	// fmt.Println(len(s0), len(s1), len(s2))

	// 拼接字符串 + / fmt.Sprintf()
	// s3 := s1 + " " + s2
	// fmt.Println(s3)
	// s4 := fmt.Sprintf("%s %s", s1, s2)
	// fmt.Println(s4)

	// 分割字符串

	// var s5 = "What's-your-name?"
	// fmt.Println(strings.Split(s5, "-")) // [What's your name?]
	// fmt.Printf("%T", strings.Split(s5, "-")) //[]string 字符串切片

	// 判断是否包含

	// var s6 = "how do you do?"
	// fmt.Println(strings.Contains(s6, "do")) // true

	// 判断前缀 HashPrefix
	// fmt.Println(strings.HasPrefix(s6, "how")) // true

	// 判断后缀HasSuffix
	// fmt.Println(strings.HasSuffix(s6, "how")) // false

	// 判断字符串的位置
	// fmt.Println(strings.Index(s6, "do")) // 4

	// 判断字符串最后出现的位置
	// fmt.Println(strings.LastIndex(s6, "do")) // 11

	// join

	// s7 := []string{"how", "do", "you", "do"} // 切片
	// fmt.Println(s7)
	// fmt.Println(strings.Join(s7, "-"))
	// fmt.Printf("%T", strings.Join(s7, "-")) // string

	// 大小写

	s8 := "hELLo, World"
	// 大写
	fmt.Println(strings.ToUpper(s8))
	// 小写
	fmt.Println(strings.ToLower(s8))
	// Title
	fmt.Println(strings.Title(s8))

}
