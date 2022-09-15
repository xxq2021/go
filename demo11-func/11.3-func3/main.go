package main

import (
	"fmt"
	"sort"
)

// 把选择排序封装成方法，实现整型切片的升序、降序排列
func sortIntAsc(slice []int) []int {
	// int 类型升序排序
	// 12, 34, 37, 35, 556, 36, 2
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice
}

// int 类型降序排序
func sortInsDesc(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] < slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice
}

// 封装一个方法，要求按照 key 升序排序，最后输出一个 key=>value格式
var m = map[string]string{
	"username": "zhangsan",
	"age":      "18",
	"sex":      "male",
	"height":   "180",
}

func sortMap(map1 map[string]string) string {

	var sliceKey []string
	for k := range map1 {
		// fmt.Println(k)
		// 1. 把 map 对象的 key 放在一个切片中
		sliceKey = append(sliceKey, k)
	}
	// 2. 对 key 进行升序排序
	sort.Strings(sliceKey)
	// 3. 返回字符串
	var str string
	for _, v := range sliceKey {
		str += fmt.Sprintf("%v=>%v", v, map1[v])
	}
	return str
}

func main() {

	var sliceA = []int{12, 34, 37, 35, 556, 36, 2}

	// 升序
	sortSlicaA := sortIntAsc(sliceA)
	fmt.Println(sortSlicaA) // [2 12 34 35 36 37 556]

	// 降序
	sortSlicaB := sortInsDesc(sliceA)
	fmt.Println(sortSlicaB) // [556 37 36 35 34 12 2]

	fmt.Println(sortMap(m))

}
