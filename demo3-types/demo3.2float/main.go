package main

import (
	"fmt"
	"unsafe"
	// "math"
)

func main() {
	// fmt.Println(math.MaxFloat32) // 3.4028234663852886e+38
	//fmt.Println(math.MaxFloat64) // 1.7976931348623157e+308

	// %f 输出float类型，%.xf  输出数据时保留 x 位小数。

	var num = 3.1415926
	fmt.Printf("num = %v, 类型为%T", num, num)        // num = 3.1415926, 类型为float64
	fmt.Printf("\nnum保留三位小数为%.3f", num)            // num保留三位小数为3.142
	fmt.Printf("\nnum保留四位小数为%.4f", num)            // num保留四位小数为3.1416
	fmt.Println("\nnum占", unsafe.Sizeof(num), "位") //num占 8 位

}
