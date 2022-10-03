package main

import "fmt"

func main() {
	num := []int{3, 1, 2, 5, 4, 9, 7, 2}
	//fmt.Println(num)

	for i := 0; i < len(num); i++ {
		for j := i + 1; j < len(num); j++ {
			if num[i] == num[j] {
				fmt.Println(num[i], i, j)
			}

		}
	}
}
