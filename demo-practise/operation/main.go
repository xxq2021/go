package main

import (
	"fmt"
)

func main() {
	// 运算符

	// 算术运算符
	// s1 := 5
	// s2 := 2
	// fmt.Println(s1 + s2) // 7
	// fmt.Println(s1 - s2) // 3
	// fmt.Println(s1 * s2) // 10
	// fmt.Println(s1 / s2) // 2
	// fmt.Println(s1 % s2) // 1

	// 关系运算符
	// fmt.Println(s1 == s2) // false
	// fmt.Println(s1 != s2) // true
	// fmt.Println(s1 > s2)  // true
	// fmt.Println(s1 >= s2) // true
	// fmt.Println(s1 < s2)  // false
	// fmt.Println(s1 <= s2) // false

	// 逻辑运算符
	// fmt.Println((s1 == s2) && (s1 > s2)) // false
	// fmt.Println((s1 == s2) || (s1 > s2)) // true
	// fmt.Println(!(s1 == s2))             // true

	// 位运算符
	// s1 := 5 => 101
	// s2 := 2 => 010
	// fmt.Println(s1 & s2) // 0
	// fmt.Println(s1 | s2) // 111 => 7
	// fmt.Println(s1 ^ s2) // 111 => 7

	// 左移 << => 乘以2^n
	// fmt.Println(s1 << s2) // s1 = 5 => 101 << s2(2) => 10100 => 20
	// fmt.Println(1 << 10)  // 1 => 100 0000 0000 => 1024

	// 右移 >> => 除以2^n
	// fmt.Println(s1 >> s2) // s1 = 5 => 101 >> s2(2) => 1
	// fmt.Println(s1 >> 1)  // s1 = 5 => 101 >> 10 => 2\

	// 赋值运算符
	var a int
	a = 10
	a += 5         // a = a + 5
	fmt.Println(a) // 15
	a -= 5
	fmt.Println(a) // 10
	a *= 2
	fmt.Println(a) // 20
	a /= 4
	fmt.Println(a) // 5
	a %= 2
	fmt.Println(a) // 1
	a <<= 10
	fmt.Println(a) // 100 0000 0000 => 1024
	a >>= 5
	fmt.Println(a) // 10 0000 => 32

	b := 4         // 100
	b &= 2         // 010
	fmt.Println(b) // 0

	c := 4         // 100
	c |= 2         // 010
	fmt.Println(c) // 110 => 6

	d := c         // 6
	d ^= b         // 6 ^= 0
	fmt.Println(d) // 110 => 6
}
