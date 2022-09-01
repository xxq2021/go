package main

import (
	"fmt"
	"sort"
)

func main() {
	// int
	var a = []int{4, 1, 3, 5, 2}
	sort.Ints(a)
	fmt.Println(a)

	// string
	var b = []string{"e", "c", "b", "d", "a"}
	sort.Strings(b)
	fmt.Println(b)

	// float64
	var c = []float64{5.2, -1.3, 0.7, -4.6, 2.6}
	sort.Float64s(c)
	fmt.Println(c)

	// 逆序 int
	var d = []int{4, 1, 3, 5, 2}
	sort.Sort(sort.Reverse(sort.IntSlice(d)))
	fmt.Println(d)

	// 逆序 string
	var e = []string{"e", "c", "b", "d", "a"}
	sort.Sort(sort.Reverse(sort.StringSlice(e)))
	fmt.Println(e)
}
