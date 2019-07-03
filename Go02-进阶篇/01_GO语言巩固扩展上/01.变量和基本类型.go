package main

import (
	"fmt"
	"math"
)

func main020101() {
	/*
	   一：Go语言的架构
	   	GO语言简洁：GO语言是一个全新的 静态类型编程语言
	   	    1.静态类型和动态类型的区别：
	   		    静态类型： 编译时即知道每一个变量的类型，因此，若存在类型错误编译是无法通过的。
	   		    动态类型： 编译时不知道每一个变量的类型，因此，若存在类型错误会在运行时发生错误。
	   		2.编译型和解释性的区别
	   			编译型： 翻译发生在程序运行之前，将高级语言翻译成机器语言。再次运行时，可直接使用上一次翻译好的机器码，不需要重新编译。
	   			解释型： 翻译发生在程序运行时，即边翻译边运行。再次运行时，需要重新进行翻译。

	       GO语言的特点：
	   		1.编译速度很快，因为它本身是一个静态类型的语言
	   		2.静态类型的语言 但是又有着动态类型的特点 第三方包
	   		3.自身支持垃圾回收机制
	   		4.丰富的标准库
	   		5.高效的垃圾回收机制
	*/

	/*
	   二：Go的注释
	   	//
	   	/**\/
	*/

	/*
	   三：转义字符：
	   	\n \r\n \t
	*/

	/*
	   四:变量

	*/
	var name string
	name = "张三"
	var age = 25
	add := "徐家汇"
	fmt.Println(name, age, add)

	//一次声明多个变量
	var (
		a = 255
		b = "sss"
	)
	fmt.Printf("数据类型：%T\n", b)
	fmt.Println(a, b)

	fmt.Println(math.MaxInt32)
	fmt.Println(math.MinInt8)
	fmt.Println(math.MinInt16)
}

func main020102() {
	num1, num2, num3, num4, num5 := 3222, 3.14213123, true, "搜索", 'a'
	fmt.Printf("整型：%d\n", num1)
	fmt.Printf("浮点型：%f\n", num2)
	fmt.Printf("布尔型：%t\n", num3)
	fmt.Printf("字符串型：%s\n", num4)
	fmt.Printf("字节型:%v\n", num5)
	fmt.Printf("100%%\n")
	//%v : 按照f数据打印 自动匹配数据格式
	//%T ： 输出数据类型
	//%p 输出数据指针
	fmt.Printf("输出字节型：%c\n", num5)
	fmt.Printf("输出指针地址：%p\n", &num5)
	fmt.Printf("输出：%q\n", num5)

	fmt.Printf("%5d\n", num1)
	fmt.Printf("%.5d\n", num1)
	fmt.Printf("%.2f\n", num2)
}

func main020203() {
	var (
		name string
		age  int
	)
	//fmt.Scanf("姓名：%s,年龄：%d", &name, &age)
	//fmt.Scanf("姓名：%s,年龄：%d", &name, &age)
	//fmt.Scan(&name, &age)
	fmt.Println(name, age)
	type myint int64
	var namex myint = 254
	fmt.Printf("数据类型：%T,数据的值是：%v\n", namex, namex)
	type (
		myset int
		mygod float64
	)
	var namew myset = 648
	var my mygod = 258
	fmt.Printf("数据类型：%T,数据的值是：%v\n", namew, namew)
	fmt.Printf("数据类型：%T,数据的值是：%v\n", my, my)
}
func main020() {
	const (
		a, b = iota, iota //初始值是0，每一行自动加1
		c, d = iota, iota
		e, f  //不写iota的行的格式必须与上一行格式一致
	)
	fmt.Println(a, b, c, d, e, f) //0 0 1 1 2 2
}
