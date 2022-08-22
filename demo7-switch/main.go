package main

import "fmt"

func main() {
	// 1. case 判断

	//figure := 3
	//switch figure {
	//case 1:
	//	fmt.Println("大拇指")
	//case 2:
	//	fmt.Println("食指")
	//case 3:
	//	fmt.Println("中指")
	//case 4:
	//	fmt.Println("无名指")
	//case 5:
	//	fmt.Println("小拇指")
	//default:
	//	fmt.Println("无效输入")
	//}

	// 2. case 一次判断多个值
	//num := 5
	//switch num {
	//case 1, 3, 5, 7, 9:
	//	fmt.Println("奇数")
	//case 2, 4, 6, 8, 10:
	//	fmt.Println("偶数")
	//}

	// 3. case中使用表达式
	age := 18
	switch {
	case age >= 18:
		fmt.Println("可以上网")
	case age < 18:
		fmt.Println("不可以上网")
	default:
		fmt.Println("嘿嘿嘿")
	}
}
