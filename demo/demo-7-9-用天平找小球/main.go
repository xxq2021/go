/* 7-9 用天平找小球
三个球A、B、C，大小形状相同且其中有一个球与其他球重量不同。要求找出这个不一样的球。

输入格式：输入在一行中给出3个正整数，顺序对应球A、B、C的重量。
输出格式：在一行中输出唯一的那个不一样的球。

输入样例：1 1 2
输出样例：C
*/

// 分析: 所有情况 1 1 2    1 2 1    2 1 1
// 		   C        B        A

package main

import "fmt"

func main() {
	var (
		a int
		b int
		c int
	)
	fmt.Scan(&a, &b, &c)
	// fmt.Println(a, b, c)

	if a == b && a != c {
		fmt.Println("C")
	} else if a == c && a != b {
		fmt.Println("B")
	} else {
		fmt.Println("A")
	}
}
