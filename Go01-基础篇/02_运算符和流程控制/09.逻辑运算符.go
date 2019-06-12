package main

import (
	"fmt"
	"strconv"
)

func main020901() {
	/*
		!总结：
			逻辑运算符：
				1.非【!】 非真为假 非假为真
				2.或【||】同假为假 其余为真
				3.与【&&】同真为真 其余为假
		注意：逻辑运算符只针对于布尔类型的数据值 或者表达式使用
	*/
	var a bool
	a = true

	b := 20
	c := 20

	fmt.Println(!(b == c)) //false 取反
	fmt.Println(!a)        //false
}

func main020902() {
	//逻辑于 && 并且关系 表达式1 && 表达式2 同真为真 其余为假
	a := 10
	b := 20
	c := 30
	d := 40

	fmt.Println(a > b && c < d) //false
	fmt.Println(a < b && d > c) //true
}

func main020903() {
	//逻辑或 || 或者关系 表示式1 || 表达式2 同假为假 其余为真
	a := 10
	b := 20
	c := 30
	d := 40

	fmt.Println(a > b || c > d) //false
	fmt.Println(a < b || c > d) //true
}

func main020904() {
	//取地址运算符
	var a int = 10
	fmt.Println(&a) //0xc000064080 取的是栈区的内存地址

	b := 10
	//指针变量
	p := &b
	//* 取值运算符
	fmt.Println(*p) //输出值为10
}

//补充 数据类型互换
//字符串和数字类型互换
func main020905() {
	num1, num2 := 123, "321"
	fmt.Println(strconv.Itoa(num1)) //整形转成字符串
	fmt.Printf("%T", num1)

	strconv.Atoi(num2)
	num, err := strconv.Atoi(num2)
	if err == nil { //错误为nil表示没有错误
		fmt.Println("转换成功", num)
	} else {
		fmt.Println("转换错误", err)
	}
	//nil是预定义的标识符，代表变量的零值，也就是预定义好的一个变量。类似于null
}
