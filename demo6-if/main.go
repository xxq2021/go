package main

func main() {
	// if 判断

	//1. 基本写法
	//var score = 65
	//if score >= 90 {
	//	fmt.Println("A")
	//} else if score > 75 {
	//	println("B")
	//} else {
	//	fmt.Println("C")
	//}

	// 2. 特殊写法
	//if score := 65; score >= 90 {
	//	fmt.Println("A")
	//} else if score > 75 {
	//	println("B")
	//} else {
	//	fmt.Println("C")
	//}
	//fmt.Println(score) //undefined: score

	// for 循环
	// 1. 基本写法
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}

	// 2. 省略初始语句但保留初始语句后面的分号
	//var i = 0
	//for ; i < 10; i++ {
	//	fmt.Println(i)
	//}

	// 3. 省略初始语句和结束语句 (while 写法)
	//var i = 10
	//for i > 0 {
	//	fmt.Println(i)
	//	i--
	//}

	// 4. 无限循环
	//for {
	//	fmt.Println("Hello")
	//}

	// 1-10 的偶数
	//for i := 1; i <= 10; i++ {
	//	if i%2 == 0 {
	//		fmt.Println(i)
	//	}
	//}

	// break 跳出循环
	//for i := 1; i < 10; i++ {
	//	if i > 5 {
	//		break // 1 2 3 4 5
	//	}
	//	fmt.Println(i)
	//}

	// continue 跳出当前for循环，进入下一次循环
	//for i := 1; i < 10; i++ {
	//	if i == 5 {
	//		continue // 1 2 3 4 6 7 8 9
	//	}
	//	fmt.Println(i)
	//}

}
