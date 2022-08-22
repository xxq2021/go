package main

import "fmt"

func main() {
	// 数组的遍历
	// a := [...]string{"北京", "上海", "广州", "深圳", "苏州"}
	// for循环

	// for i := 0; i <= len(a)-1; i++ {
	// 	fmt.Println(a[i])
	// }

	// for range

	// for k, v := range a {
	// 	fmt.Printf("key: %v value: %v\n", k, v)
	// }

	// 求出一个数组里面元素的和以及这些元素的平均值，分别用 for 和 for range

	// var a = [...]int{12, 23, 45, 67, 2, 5}
	// sum := 0
	// for i := 0; i < len(a); i++ {
	// 	sum += a[i]
	// }
	// fmt.Println("Sum:", sum)
	// fmt.Printf("Average:%.2f", float64(sum)/float64(len(a)))

	// 求出一个数组中最大值，并得到下标
	// var intArr = [...]int{1, -1, 12, 65, 11}
	// 假设第一个是最大
	// var intMax = intArr[0]
	// 定义下标
	// var index = 0
	// for i := 0; i < len(intArr); i++ {
	// 	if intArr[i] > intMax {
	// 		intMax = intArr[i]
	// 		index = i
	// 	}
	// }
	// fmt.Printf("最大值为:%v, 对于索引为: %v", intMax, index)

	// 从数组[1,3,5,7,8]中找到和为8的两个元素的下标分别为(0,3)和(1,2)

	// 定义数组

	var array = [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i]+array[j] == 8 {
				fmt.Printf("(%v,%v)", i, j)
			}
		}
	}
}
