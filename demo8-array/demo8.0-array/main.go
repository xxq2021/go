package main

import "fmt"

func main() {
	//var a [4]int
	//var b [5]int
	//fmt.Println(a,b)

	// 数组的初始化
	// 1. 定义时使用初始值列表的方式初始化
	//var cityArray = [4]string{"北京", "上海", "广州", "深圳"}
	//fmt.Println(cityArray)
	//fmt.Println(cityArray[0])
	//fmt.Println(cityArray[3])

	// 2. 编译器推导数组的长度
	//var boolArray = [...]bool{true, false, false, true, true}
	//fmt.Println(boolArray)
	//fmt.Println(len(boolArray))

	// 3. 使用索引值方式初始化
	//var longArray = [...]string{1: "Golang", 3: "Python", 7: "java"}
	//fmt.Println(longArray)
	//fmt.Println(len(longArray))
	//fmt.Printf("%T", longArray)

	// 数组的遍历
	//var cityArray = [4]string{"北京", "上海", "广州", "深圳"}

	// 1. for循环
	//for i := 0; i < len(cityArray); i++ {
	//	fmt.Println(cityArray[i])
	//}

	// 2. for range 遍历
	//for index, value := range cityArray {
	//	fmt.Println(index, value)
	//}
	//
	//for _, value := range cityArray {
	//	fmt.Println(value)
	//}

	// 二维数组
	cityArray := [3][2]string{
		{"北京", "西安"},
		{"上海", "杭州"},
		{"重庆", "成都"},
	}
	//fmt.Println(cityArray)
	//fmt.Println(len(cityArray))
	//fmt.Printf("%T", cityArray)
	//fmt.Println(cityArray[2][0])

	// 二维数组遍历
	for _, v1 := range cityArray {
		//fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	// 多维数组只有最外层可以使用[...]，最内层不可以

	// 数组是值类型，赋值的时候把值做了拷贝
	x := [3]int{1, 2, 3}
	fmt.Println(x)
	f1(x)
	fmt.Println(x)

	y := [3][2]int{
		{1, 2},
		{2, 3},
		{3, 4},
	}
	fmt.Println(y)
	f2(y)
	fmt.Println(y)
	z := y
	z[0][0] = 10
	fmt.Println(z)
}
func f1(a [3]int) {
	a[0] = 100
}
func f2(b [3][2]int) {
	b[0][0] = 100
}
