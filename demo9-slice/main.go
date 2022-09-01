package main

import "fmt"

func main() {
	// 切片是一个拥有相同类型元素的可变长度的序列
	// 它是基于数组类型做的一层封装，支持自动扩容
	// 切片是一个引用类型，它的内部结构包括 地址、长度和容量
	// 一般用于快速地操作一块数据集合

	// 切片的声明 初始化
	//var a []int
	//fmt.Printf("%v - %T - 长度:%v", a, a, len(a))
	//fmt.Println(a == nil) // true

	//var b []string
	//fmt.Println(b == nil)
	//
	//var c []bool
	//fmt.Println(c == nil)

	// 无法扩容
	//a[0] = 1
	//fmt.Println(a) // panic: runtime error: index out of range [0] with length 0

	//var a = []string{"北京", "上海", "广州", "深圳"}
	//fmt.Printf("%v - %T - 长度:%v", a, a, len(a))

	//var a = []int{1, 2, 3, 4, 5}
	//fmt.Printf("%v - %T - 长度:%v", a, a, len(a))

	//var a = []int{0: 1, 1: 3, 3: 5, 5: 7}
	//fmt.Printf("%v - %T - 长度:%v", a, a, len(a))
	//var a []string
	//var b = []int{}
	//var c = []bool{false, true}
	//var d = []bool{false, true}
	//fmt.Println(a) // []
	//fmt.Println(b) // []
	//fmt.Println(c) // [false true]
	//fmt.Println(d) // [false true]
	//fmt.Println(a == nil) // true
	//fmt.Println(b == nil) // false
	//fmt.Println(c == nil) // false
	//fmt.Println(d == nil) // false

	// 切片的循环遍历
	//var a = []string{"北京", "上海", "广州", "深圳", "成都"}
	// for 循环
	//for i := 0; i < len(a); i++ {
	//	fmt.Println(a[i])
	//}
	// for range循环

	//for _, v := range a {
	//	fmt.Println(v)
	//}
	// len()函数获取切片的长度，cap()获取切片的容量

	// 基于数组的切片
	//var a = [5]int{1, 2, 3, 4, 5}
	//b := a[:]
	//fmt.Printf("%v - %T", b, b)
	//c := a[2:]
	//fmt.Println(c)
	//fmt.Printf("切片长度：%v,切片容量：%v", len(c), cap(c))
	//d := a[:3]
	//fmt.Println(d)
	//fmt.Printf("切片长度：%v,切片容量：%v", len(d), cap(d))
	//e := a[1:3] // 左包含，右不包含
	//fmt.Println(e)
	//fmt.Printf("切片长度：%v,切片容量：%v", len(e), cap(e))

	// 基于切片的切片
	//var a = []string{"北京", "上海", "广州", "深圳", "成都"}
	//var b = a[3:]
	//fmt.Println(b)
	//a = append(a, 1, 2, 3, 4, 5)
	//b := []int{11, 12, 13, 14, 15}
	//a = append(a, b...)
	//fmt.Println(a)

	// make 函数创建切片
	//a := make([]int, 5, 8)
	//fmt.Printf("%v - %T - len: %v - cap: %v", a, a, len(a), cap(a))
	// 切片赋值
	//a[0] = 1
	//a[1] = 2
	//a[2] = 3
	//a[3] = 4
	//a[4] = 5
	//fmt.Println(a) // [1 2 3 4 5]
	// 修改切片
	//var b = []string{"北京", "上海", "广州", "深圳"}
	//b[2] = "苏州"
	//fmt.Println(b)

	// 无法通过下标给切片扩容

	//var b = []string{"北京", "上海", "广州", "深圳"}
	//fmt.Printf("len: %v - cap: %v", len(b), cap(b))
	//b[4] = "苏州"
	//fmt.Println(b) // panic: runtime error: index out of range [4] with length 4

	// append() 数组扩容
	// append的使用
	//var a = make([]int, 4, 8)
	//a = append(a, 1)
	//fmt.Println(a)
	//fmt.Printf("len: %v - cap: %v", len(a), cap(a))

	// 合并数组
	//b := []int{2, 3, 4}
	//b = append(a, b...)
	//fmt.Println(b)

	// 切片的扩容
	//var a []int
	//for i := 0; i < 10; i++ {
	//	a = append(a, i)
	//	fmt.Printf("%v - len: %v - cap: %v - ptr: %p\n", a, len(a), cap(a), a)
	//}

	//a := []int{1, 2, 3, 4, 5}
	//b := make([]int, 5, 5)
	//c := b
	//fmt.Println(a)
	//fmt.Println(b)
	//copy(b, a)
	//b[0] = 11
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)

	//var a = make([]string, 5, 10)
	//for i := 0; i < 10; i++ {
	//	a = append(a, fmt.Sprintf("%v", i))
	//}
	//fmt.Println(a)
	//fmt.Println(len(a))

	//b := [...]int{3, 8, 6, 4, 9, 7}
	//
	//fmt.Println(b)
	//
	//sort.Ints(b[:])
	//fmt.Println(b)

	//var a = []int{0, 1, 2, 3, 4, 5}
	//b := a[:3]
	//fmt.Printf("切片长度：%v,切片容量：%v", len(b), cap(b))
	//c := a[2:]
	//fmt.Printf("切片长度：%v,切片容量：%v", len(c), cap(c))
	//d := a[2:4]
	//fmt.Printf("切片长度：%v,切片容量：%v", len(d), cap(d))

	//var a = []int{0, 1, 2, 3, 4, 5}
	//b := a
	//b[0] = 2
	//fmt.Println(a) // [2 1 2 3 4 5]
	//fmt.Println(b) // [2 1 2 3 4 5]

	// copy函数
	//var a = []int{0, 1, 2, 3, 4, 5}
	//b := make([]int, 6, 8)
	//copy(b, a)
	//fmt.Println(a) // [0 1 2 3 4 5]
	//fmt.Println(b) // [0 1 2 3 4 5]
	//b[0] = 11
	//fmt.Println(a) // [0 1 2 3 4 5]
	//fmt.Println(b) // [11 1 2 3 4 5]

	// 切片中删除元素
	a := []int{1, 2, 3, 4, 5}
	// 删除索引为 2 的元素 a
	a = append(a[:2], a[3:]...)
	fmt.Println(a)
}
