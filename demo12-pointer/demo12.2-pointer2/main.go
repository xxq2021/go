package main

import "fmt"

func modify1(x int) {
	x = 100
}
func modify2(y *int) {
	*y = 100
}

func main() {
	a := 1
	modify1(a)
	fmt.Println(a) // 1
	//x := 20
	//modify1(x)
	//fmt.Println(x) // 20
	modify2(&a)
	fmt.Println(a) // 100
}
