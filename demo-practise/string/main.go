package main

import (
	"fmt"
	"strings"
)

func main() {
	// s := "Hello World"
	// fmt.Println(s) // Hello World
	// 字符串常用方法

	// 查看字符串的长度 len()
	// fmt.Println(len(s)) // 11

	// 拼接字符串 + /Sprintf
	// s1 := "hello"
	// s2 := "go"

	// + 拼接
	// fmt.Println(s1 + " " + s2) // hello go

	// Sprintf
	// s3 := fmt.Sprintf("%s %s", s1, s2)
	// fmt.Println(s3) // hello go

	// 分割字符串
	// s4 := strings.Split(s3, " ")
	// fmt.Println(s4)    // [hello go]
	// fmt.Println(s4[0]) // hello
	// fmt.Println(s4[1]) // go

	// 判断是否包含
	// s5 := "how do you do?"
	// fmt.Println(strings.Contains(s5, "do"))    // true
	// fmt.Println(strings.Contains(s5, "where")) // false

	// 前缀
	// fmt.Println(strings.HasPrefix(s5, "h")) // true
	// fmt.Println(strings.HasPrefix(s5, "w")) // false

	// 后缀
	// fmt.Println(strings.HasSuffix(s5, "?")) // true
	// fmt.Println(strings.HasSuffix(s5, "s")) // false

	// 第一次出现的位置
	// fmt.Println(strings.Index(s5, "o")) // 1
	// fmt.Println(strings.Index(s5, "s")) // -1

	// 最后一次出现的位置
	// fmt.Println(strings.LastIndex(s5, "o")) // 12
	// fmt.Println(strings.LastIndex(s5, "s")) // -1

	// Join拼接
	// s6 := []string{"how", "do", "you", "do", "?"}
	// fmt.Println(s6)                    //[how do you do ?]
	// fmt.Println(strings.Join(s6, " ")) // how do you do ?
	// fmt.Printf("%T", strings.Join(s6, " ")) // string

	// 大写
	// s7 := "how dO yOu Do"
	// fmt.Println(s7)                  // how dO yOu Do
	// fmt.Println(strings.ToUpper(s7)) // HOW DO YOU DO

	// 首字母大写
	// fmt.Println(strings.Title(s7)) // How DO YOu Do

	// 小写
	// fmt.Println(strings.ToLower(s7)) // how do you do

	// 去除字符串
	// 去除字符串前后端的符号

	s8 := " Hello Wo rld  "
	// fmt.Println(s8)
	// fmt.Println(len(s8)) // 15
	// fmt.Println(strings.Trim(s8, " "))
	// fmt.Println(len(strings.Trim(s8, " "))) // 12

	// fmt.Println(strings.TrimSpace(s8))
	// fmt.Println(len(strings.TrimSpace(s8))) // 12

	// 去除左端
	fmt.Println(strings.TrimLeft(s8, " "))
	fmt.Println(len(strings.TrimLeft(s8, " "))) // 14

	// 去除右端
	fmt.Println(strings.TrimRight(s8, " "))
	fmt.Println(len(strings.TrimRight(s8, " "))) //13

}
