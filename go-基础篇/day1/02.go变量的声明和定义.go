package main

import "fmt"

func main0201() {
	//go基本数据类型：布尔型 整形 浮点型 字符型 字符串型
	var a int      //声明
	a = 125        //赋值
	a += 20        //运算
	fmt.Println(a) //打印
}

func main0202() {
	//计算圆的面积和周长
	//面积 PI * r * r
	//周长 2 * PI * r
	var PI float64 = 3.1415926
	var r float64 = 2.5
	var s float64
	var l float64
	//计算面积
	s = PI * r * r
	l = 2 * PI * r
	fmt.Println("面积", s)
	fmt.Println("周长", l)
	//小节：声明以后如果不进行个定义的话会报错
}

func main0203() {
	//自动推导数据类型
	PI := 3.1415926
	r := 2.5
	//var s float64
	//var l float64
	s := PI * r * r
	l := 2 * PI * r
	fmt.Println("面积", s)
	fmt.Println("周长", l)
}

//go的类型转换 【类型转换引导】
func main0204() {
	//去市场买2斤黄瓜 价格为五元
	w := 2.8
	var t float64 = 5
	//！在go语言中 不同的数据类型 不能进行计算操作 可以通过类型转换解决问题
	fmt.Println(w * t)
	//自动推导类型： 语法 := 会根据值来进行类型推导
}

/*
总结变量定义格式：
声明 var 变量 数据类型
赋值 var 变量 数据类型 = 赋值
变量名 := 自动推导类型[最常使用]
！变量的类型不同不能进行计算 需要使用类型转换
*/

func main0205() {
	//复习：go的基本数据类型
	a := false  //布尔类型
	b := 20     //整形
	c := 3.14   //浮点型
	d := "平渊渡海" //字符串类型
	e := 'a'    //字符型
	fmt.Println(a, b, c, d, e)
}
