package main

import "fmt"

func main0601() {
	a := 10
	b := 2.5
	//将不同的数据类型 转成相同的数据类型再进行操作
	c := float64(a) * b
	fmt.Println(c)
	/*
		!总结：
			1.类型转换格式 数据类型(变量) 数据类型(表达式)
			2.我们一般进行转换 是由低类型转成高类型的进行转换 主要是为了数据的完整性
				(如果此处我们转换b float 它只会保留整形 而舍弃掉了小数点后面的数值 不利于数据完整)
	*/
}

func main0602() {
	var a int32 = 10
	var b int64 = 20
	//c := a * b
	c := int64(a) * b
	fmt.Println(c)
	/*
		!总结:
			1.虽然int32和int64都是整型 单数不允许相互转换
			2.只有类型匹配的数据才能进行计算
			3.在go语言中习惯将低类型转成高类型  保证数据完整性
	*/
}

func main0603() {
	//小案例 A
	//编程实现107653秒是几天几小时几分钟几秒?
	time := 107653
	fmt.Println("天：", time/60/60/24%365)
	fmt.Println("时：", time/60/60%24)
	fmt.Println("分：", time/60%60)
	fmt.Println("秒：", time%60)

	//小案例 B
	//46天是第几周第几天
	day := 46
	fmt.Printf("%d周%d天", day/7, day%7)
}
