package main

import "fmt"

/*
总结：
	1.%d 表示整型
	2.%t 表示布尔型
	3.%s 表示字符串
	4.%c 表示字符型
	5.%p 占位符 表示输出一个数据对应的内存地址
*/
func main0501() {
	//获取用户的输入->第一步：声明变量
	var a int
	//& 运算符 取地址运算符
	fmt.Scan(&a)
	//fmt.Println(a)
	fmt.Printf("您将要取出%p元", &a) //您将要取出0xc000064080元
	//补充 0x表示十六进制的数据
}

func main0502() {
	//空格或者回车作为接收结束
	//var str string
	//fmt.Scan(&str)
	//fmt.Println(str)
	//如何接收两个数据？
	var str1, str2 string
	fmt.Scan(&str1, &str2)
	fmt.Println(str1, str2)
}

func main0503() {
	var r float64
	//PI := 3.1415926
	var PI float64
	//通过键盘获取半径
	fmt.Scan(&r, &PI)
	fmt.Printf("面积：%.2f\n", PI*r*r)
	fmt.Printf("周长：%.2f\n", 2*PI*r)
	//[★]重点记忆一下%.这种写法
}

func main0504() {
	var a int
	var b string
	fmt.Scanf("%3d", &a)
	fmt.Scanf("%s", &b)
	fmt.Println(a, b)
	//！标记 这个地方明天再熟悉一下
}

func main0505() {
	//var name string
	//var pwd string
	//fmt.Println("请您输入用户名")
	//fmt.Scanf("%s",&name)
	//fmt.Println("请您输入密码")
	//fmt.Scanf("%s",&pwd)
	//fmt.Println(name,pwd)

	var Int int
	Int = 25
	fmt.Println(Int)
}
