/* 7-3 逆序的三位数
程序每次读入一个正3位数，然后输出按位逆序的数字。注意：当输入的数字含有结尾的0时，输出不应带有前导的0。比如输入700，输出应该是7。

输入格式：每个测试是一个3位的正整数。
输出格式：输出按位逆序的数。

输入样例：123
输出样例：321
*/

package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)                           // 123
	baiWei := num / 100                      // 123 / 100 = 1
	shiWei := (num / 10) % 10                // (123 / 10) = 23 % 10 = 2
	geWei := num % 10                        // 123 % 10 = 3
	newNum := geWei*100 + shiWei*10 + baiWei // 相加起来而不是简单的互换位置 3 * 100 +2 * 100 + 1
	fmt.Println(newNum)
}
