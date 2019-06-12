package main

import "fmt"

//Println和Println的区别：
func main0401() {
	//fmt.Println("阿萨")
	//fmt.Println(10)
	//fmt.Println(2.654)
	fmt.Print("莎莎")
	fmt.Print(10)
	fmt.Print(2.654)
	/*
	   【★】
	   Println和Println的区别：
	      Println打印数据自动换行
	      Print打印数据不带换行
	      Printf保证数据对齐 可以一次性有结构的输出多个变量的值
	*/
}

//golang的数据占位符
func main0402() {
	//a := 10
	//b := 3.1415926
	//fmt.Println("%d",a)
	//fmt.Println("%f",b)
	var a bool
	//fmt.Println(a)
	fmt.Printf("a=%t\n", a)
	fmt.Printf("%t\n", a)
	var b string
	b = "O(∩_∩)O哈哈~"
	fmt.Printf("输出：%s\n", b)
	var c byte
	c = 'a'
	//字符型变量对应的ASCII码值
	//fmt.Printf("%c\n",c)
	//fmt.Printf("%c",c)
	fmt.Printf("%c\n", c)
	fmt.Printf("%d\n", c)
	//%v会打印实际类型的值。
	fmt.Printf("%v\n", c)
	fmt.Println(c)
	//[★]！1. 此处的疑惑 %c 的占位符不能通过printf输出吗？
}

/*
bool布尔 byte字符 string字符串
1.布尔类型默认为false
总结：
	1.%d 表示整型
	2.%t 表示布尔型
	3.%s 表示字符串
	4.%c 表示字符型
*/
//“&”符号，表示获取内存单元的地址
