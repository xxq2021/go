package main

import "fmt"

func main() {
	/*
		// panic: runtime error: invalid memory address or nil pointer dereference
		var a *int // nil pointer
		*a = 10
		fmt.Println(*a)

		var b map[string]int // panic: assignment to entry in nil map
		b["age"] = 19
		fmt.Println(b)
	*/
	// 使用 new() 进行初始化，很少用，主要是用短变量声明法
	// func new(Type) *Type 返回该类型内存地址的指针
	var a *int
	a = new(int)
	fmt.Println(*a) // 0
	*a = 10
	fmt.Println(*a) // 100

	// make() 也是用于内存分配的，区别于new()
	// make() 只用于 slice, map, channel的内存创建，返回的类型也是这三个类型本身

	b := make(map[string]int)
	b["age"] = 19
	fmt.Println(b)
	for s, i := range b {
		fmt.Printf("%s=%d", s, i)
	}

	/* make() 与 new() 区别：
	1. 二者都是用来做内存分配的
	2. new() 得到的是一个类型的指针，并且内存对应的值是该类型的零值
	3. make() 只用于 slice, map, channel的内存创建，返回的类型也是这三个类型本身
	*/

}
