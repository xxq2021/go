package main

import (
	"fmt"
	"strings"
)

func main() {
	// 需要初始化使用，不定义是nil

	// make 创建map类型的数据
	//var userinfo = make(map[string]string)
	//userinfo["username"] = "张三"
	//userinfo["age"] = "20"
	//userinfo["sex"] = "男"

	//fmt.Printf("%v-%T", userinfo, userinfo)
	//fmt.Println(userinfo)
	//fmt.Println(userinfo["username"])

	// 2. map 也支持在声明的时候填充元素
	//var userinfo = map[string]string{
	//	"username": "张三",
	//	"age":      "20",
	//	"sex":      "male",
	//}
	//fmt.Println(userinfo)

	// 3. 类型推导
	//userinfo := map[string]string{
	//	"username": "zhangsan",
	//	"age":      "20",
	//	"sex":      "male",
	//}
	//fmt.Println(userinfo)

	// map 的循环遍历 for range

	//for k, v := range userinfo {
	//	fmt.Printf("key: %v value: %v\n", k, v)
	//}

	// map curd
	// 1. 创建修改 map 类型的数据
	//var userinfo = make(map[string]string)
	//userinfo["username"] = "zhangsan"
	//userinfo["age"] = "20"
	//fmt.Println(userinfo)

	// 2. 创建修改 map 类型的数据
	//var userinfo = map[string]string{
	//	"username": "zhangsan",
	//	"age":      "20",
	//}
	//userinfo["username"] = "lisi"
	//fmt.Println(userinfo)

	// 3. 获取 查找 map 类型的数据
	//var userinfo = map[string]string{
	//	"username": "zhangsan",
	//	"age":      "20",
	//}
	//fmt.Println(userinfo["username"])

	// 如果key存在ok为true, v为对应的值，不存在ok为false，v的值类型为零值
	//v, ok := userinfo["username"]
	//fmt.Println(v, ok) // zhangsan true

	//v, ok := userinfo["male"]
	//fmt.Println(v, ok) // 空 false

	// 4. 删除 map 类型的键值对
	// delete()函数 - delete(map 对象 key)

	//var userinfo = map[string]string{
	//	"username": "zhangsan",
	//	"age":      "20",
	//	"sex":      "male",
	//	"height":   "180",
	//}
	//fmt.Println(userinfo)
	//
	//delete(userinfo, "sex")
	//fmt.Println(userinfo)

	// 定义元素为 map 类型的切片
	//var a = make([]int, 3, 3) // 定义切片

	//var userinfo = make([]map[string]string, 3, 3)
	//fmt.Println(userinfo[0]) // map[] map不初始化的默认值为 nil
	//fmt.Println(userinfo[0] == nil) // true
	//
	//if userinfo[0] == nil {
	//	userinfo[0] = make(map[string]string)
	//	userinfo[0]["username"] = "zhangsan"
	//	userinfo[0]["age"] = "20"
	//	userinfo[0]["height"] = "180"
	//	// map[age:20 height:180 username:zhangsan]
	//}
	//if userinfo[1] == nil {
	//	userinfo[1] = make(map[string]string)
	//	userinfo[1]["username"] = "lisi"
	//	userinfo[1]["age"] = "22"
	//	userinfo[1]["height"] = "183"
	//	// map[age:22 height:183 username:lisi]
	//}
	//if userinfo[2] == nil {
	//	userinfo[2] = make(map[string]string)
	//	userinfo[2]["username"] = "wangwu"
	//	userinfo[2]["age"] = "23"
	//	userinfo[2]["height"] = "177"
	//	// map[age:23 height:177 username:wangwu]
	//}
	//fmt.Println(userinfo)

	// 循环遍历
	//for _, v1 := range userinfo {
	//	fmt.Println(v1)
	//map[age:20 height:180 username:zhangsan]
	//map[age:22 height:183 username:lisi]
	//map[age:23 height:177 username:wangwu]
	//}

	//for _, v1 := range userinfo {
	//	for k, v2 := range v1 {
	//		fmt.Printf("%v %v\n", k, v2)
	//	}
	//}

	// 值为切片类型的 map
	// 在 map 中存放一系列属性，可以把 map 类型的值定义为切片
	//var userinfo = make(map[string]string)

	//var userinfo = make(map[string][]string)
	//userinfo["hobby"] = []string{
	//	"Eating",
	//	"Coding",
	//	"Math",
	//}
	//fmt.Println(userinfo) // map[hobby:[Eating Coding Math]] // hobby 为 key, value为切片

	//userinfo["working"] = []string{
	//	"python",
	//	"golang",
	//	"java",
	//}
	//fmt.Println(userinfo) // working:[python golang java]

	// 循环遍历
	//for k, v := range userinfo {
	//	fmt.Println(k, v) // hobby [Eating Coding Math]
	//}

	//for _, v := range userinfo {
	//	for _, v1 := range v {
	//		fmt.Println(v1)
	//	}
	//}

	// map 是引用类型
	//var userinfo1 = make(map[string]string)
	//userinfo1["username"] = "zhangsan"
	//userinfo1["age"] = "20"
	//userinfo2 := userinfo1
	//fmt.Println(userinfo1)
	//fmt.Println(userinfo2)
	//userinfo2["username"] = "lisi"
	//fmt.Printf("更改后userinfo1: %v\n", userinfo1)
	//fmt.Printf("更改后userinfo2: %v\n", userinfo2)

	// map 的排序

	//map1 := make(map[int]int, 10)
	//map1[10] = 100
	//map1[1] = 13
	//map1[4] = 56
	//map1[8] = 90
	//fmt.Println(map1) // map[1:13 4:56 8:90 10:100]

	//for k, v := range map1 {
	//	fmt.Println(k, v)
	/*  每次会变化
	10 100
	4 56
	1 13
	8 90
	*/
	//}

	// 需求：按照 key 的升序输出 map 的 key 的value
	// 1. 把map的可以放到切片里

	//var keySlice []int
	//for key, _ := range map1 {
	//	keySlice = append(keySlice, key)
	//}
	//fmt.Println(keySlice) // [10 1 8 4] 每次会变化

	// 2. 让 key 升序排序
	//sort.Ints(keySlice)
	//fmt.Println(keySlice) // [1 4 8 10]

	// 3. 循环遍历 key 输出 map 的值
	//for _, v := range keySlice {
	//	fmt.Printf("key=%v value=%v\n", v, map1[v])
	//}

	// 写一个程序，统计一个字符串中每个单词出现的次数
	// 如 "how do you do" 中 how=1 do =2 you=1

	var str = "how do you do"
	var strSlice = strings.Split(str, " ")
	//fmt.Println(strSlice) // [how do you do]

	var wordCount = make(map[string]int)
	for _, v := range strSlice {
		wordCount[v]++
	}
	fmt.Println(wordCount)
}
