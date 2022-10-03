package main

import "fmt"

// 斐波那契数列
// 1、1、2、3、5、8、13、21、34
func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 使用递归求 n!
// 5! = 5*4*3*2*1 = 5 * 4! = 5 * 4 * 3!

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

// 求 n 个数和 n = 1 + 2 + 3 + 4 + …
func sumNum(n int) (sum int) {
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum
}

// 求 m 阶乘的和 m = 6 => 1! + 2! + 3! + 4! + 5! + 6!

func sumFact(m int) (factSum int) {
	for i := 1; i <= m; i++ {
		factSum += fact(i)
	}
	return factSum
}

func main() {

	x := fibonacci(3)
	fmt.Println(x)
	y := fact(5)
	fmt.Println(y)
	fmt.Println(sumNum(10))
	fmt.Println(sumFact(5))
}
