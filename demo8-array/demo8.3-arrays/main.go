package main

import "fmt"

func main() {
	// 数组是值类型，赋值和传参会复制整个数组，因此改变副本的值，不会改变本身的值
	// 基本数据类型 和 数组都是值类型
	// 值类型 （改变变量副本值的时候，不会改变变量本身的值）
	// 引用类型（改变变量副本值的时候，会改变变量本身值）

	// var a = 10
	// b := a
	// fmt.Println(a, b) // 10 10
	// a = 20
	// fmt.Println(a, b) // 20 10

	// var arr1 = [...]int{1, 2, 3}
	// arr2 := arr1
	// fmt.Println(arr1) // [1 2 3]
	// fmt.Println(arr2) // [1 2 3]

	// arr1[0] = 11
	// fmt.Println(arr1) // [11 2 3]
	// fmt.Println(arr2) // [1 2 3]

	// 切片 引用类型
	// var arr3 = []int{1, 2, 3}
	// arr4 := arr3
	// arr3[0] = 11
	// fmt.Println(arr3) // [11 2 3]
	// fmt.Println(arr4) // [11 2 3]
	// arr4[1] = 22
	// fmt.Println(arr3) // [11 22 3]
	// fmt.Println(arr4) // [11 22 3]

	// 多维数组

	// var a = [...]int{1, 2, 3, 4}

	// var a = [3][3]string{
	// 	{"北京", "上海", "广州"},
	// 	{"深圳", "浙江", "厦门"},
	// 	{"苏州", "无锡", "成都"},
	// }
	// fmt.Println(a)
	// fmt.Println(a[0][2]) // 广州
	// fmt.Println(a[2][2]) // 成都

	// 多维数组遍历
	var a = [3][3]string{
		{"北京", "上海", "广州"},
		{"深圳", "浙江", "厦门"},
		{"苏州", "无锡", "成都"},
	}

	// for range
	// for _, v1 := range a {
	// 	for _, v2 := range v1 {
	// 		fmt.Println(v2)
	// 	}
	// }

	// for

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Println(a[i][j])

		}
	}

	var b = [...][2]string{
		{"苏州", "厦门"},
		{"成都", "深圳"},
	}
	fmt.Println(b) // [[苏州 厦门] [成都 深圳]]
}
