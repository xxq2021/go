package main

import "fmt"

func main() {
	var a = [5]int{1, 3, 5, 7, 8}
	//fmt.Println(a)
	//var sum = 0

	// for 循环
	//for i := 0; i < len(a); i++ {
	//	sum += a[i]
	//}

	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == 8 {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
	//
	// for range
	//for index1, value1 := range a {
	//	for index2, value2 := range a {
	//		if value1+value2 == 8 {
	//			fmt.Println(index1, index2)
	//		}
	//	}
	//}

}
