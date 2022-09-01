package main

import "fmt"

func main() {
	var a = []int{4, 1, 2, 5, 3}
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			// 从小到大
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
			/*
				[1 4 2 5 3]
				[1 2 4 5 3]
				[1 2 3 5 4]
				[1 2 3 4 5]
			*/

			// 从大到小
			/* if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}*/
		}
		//fmt.Println(a)
	}
	fmt.Println(a)
}
