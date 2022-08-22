package main

import "fmt"

func main() {
	// 变量类型
	// bool true or false

	// a := true
	// b := false
	// c := a && b
	// d := a || b
	// e := !a

	// fmt.Println("a:", a, "b:", b, "c:", c, "d:", d, "e:", e)

	// int int8 int16 int32 int64
	// 通常使用int

	// var a int = 89
	// b := 95
	// fmt.Println("a is", a, "b is", b)

	// fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a))
	// fmt.Printf("\ntype of b is %T, size of b is %d", a, unsafe.Sizeof(b))

	// 十进制

	var a int = 10
	fmt.Printf("a=%d\n", a) // 10
	fmt.Printf("a=%b\n", a) // 二进制输出 1010

	// 八进制 以0开头
	var b int = 077
	fmt.Printf("a=%o\n", b) // 八进制输出 77
	fmt.Printf("b=%b\n", b) // 二进制输出 111111
	fmt.Printf("a=%d\n", a) // 10

	// 十六进制 以0x开头

	var c int = 0xff
	fmt.Printf("c=%x\n", c) //十六进制输出 ff
	fmt.Printf("c=%X\n", c) //十六进制输出 FF

	// uint uint8 uint16 uint32 uint64

	//float float32 float64(默认类型)
	// a, b := 5.67, 8.97
	// fmt.Printf("type of a %T b %T\n", a, b)
	// sum := a + b
	// diff := a - b
	// fmt.Println("sum", sum, "diff", diff)

	// no1, no2 := 56, 89
	// fmt.Println("sum", no1+no2, "diff", no1-no2)

	// complex complex64 (float32) complex128(float64)

	// c1 := complex(5, 7)
	// c2 := 8 + 27i
	// cadd := c1 + c2
	// fmt.Println("sum:", cadd)
	// cmul := c1 * c2
	// fmt.Println("product:", cmul)

	// string

	// first, last := "Anne", "Naveen"
	// name := first + " " + last
	// fmt.Println("my name is", name)

	// 类型转换
	x := 2
	y := 2.6
	// sumkl := k + l
	sumxy := x + int(y) // 取整（2）不会四舍五入

	fmt.Println("sumxy is", sumxy)

}
