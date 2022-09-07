package main

import "fmt"

func main() {
	/*
	                       身体质量指数BMI
	               计算公式：BMI=体重（千克）除以身高（米）的平方。
	   *BMI分类        *中国参考标准      *相关疾病发病的危险性
	   体重过低         BMI<18.5         低（但其它疾病危险性增加)
	   正常范围         18.5≤BMI<24      平均水平
	   超重             BMI≥24           增加
	   肥胖前期         24<BMI<28        增加
	   Ⅰ度肥胖         28≤BMI<30        中度增加
	   Ⅱ度肥胖         30≤BMl<40        严重增加
	   Ⅲ度肥胖         BMI≥40.0         非常严重增加

	*/

	var (
		height         int // 定义身高变量
		weight         int // 定义体重变量
		lower          float64
		meter          float64 // 定义BMI变量
		classIfication = "BMI分类: "
		serious        = "危险程度: "
	)
	fmt.Println("请输入您的身高(厘米)和体重(千克&公斤):")
	fmt.Scan(&height, &weight) // 获取用户输入
	fmt.Printf("您的身高为: %v 厘米, 体重为: %v 千克\n", height, weight)
	lower = float64(height) / 100             // 身高厘米转换成米
	meter = float64(weight) / (lower * lower) // 计算BMI指数
	fmt.Printf("您的BMI指数为: %v\n", meter)

	// 进行BMI指数判断
	switch {
	case meter >= 40:
		{
			fmt.Println(classIfication, "Ⅲ度肥胖")
			fmt.Println(serious + "非常严重增加")
		}
	case meter >= 30 && meter < 40:
		{
			fmt.Println(classIfication, "Ⅱ度肥胖")
			fmt.Println(serious + "严重增加")
		}
	case meter >= 28 && meter < 30:
		{
			fmt.Println(classIfication, "Ⅰ度肥胖")
			fmt.Println(serious, "中度增加")
		}
	case meter > 24 && meter < 28:
		{
			fmt.Println(classIfication, "肥胖前期")
			fmt.Println(serious, "增加")
		}
	case meter >= 18.5 && meter < 24:
		{
			fmt.Println(classIfication, "正常范围")
			fmt.Println(serious, "平均水平")
		}
	default:
		fmt.Println(classIfication, "体重过低")
		fmt.Println(serious, "低")
	}
}
