package main

import "fmt"

func demo1(a string) {
	a = "传说中的暴龙兽"
	fmt.Println(a)
}
func main1201() {
	/*
		A:局部变量： 在函数内部定义的变量 作用域限定于本函数的内部 从变量定义到本函数运行结束时有效
			注：在同一作用域范围内 变量名范围是唯一的
	*/
	//demo1()
	var func1 func(string)
	func1 = demo1
	func1("传说中的加鲁鲁")

	a := 201
	a = 10
	{
		//匿名内部函数
		a := 20
		fmt.Println(a)
	}
	fmt.Println(a)

	/*
		!总结：
			1.如果程序中出现了相同的变量名 如果本函数中的变量有值 就使用本函数中的 如果没有就会向上层进行查找
			2.名字相同会采用就近原则 使用自己的变量 重复赋值会覆盖上次的值
	*/

	i := 0
	for i := 0; i <= 10; i++ {
		//此处下面的值输出为0 这个地方是不是可以把for也作一个函数体来理解
	}
	fmt.Println(i)

	for ; i < 10; i++ {
		//此处下面的值输出为 10
	}
	fmt.Println(i)

	for ; i <= 10; i++ {
		//此处下面的值输出为 11
	}
	fmt.Println(i)
}

var a int = 100   //全局变量
const A int = 200 //全局常量

func main1202() {
	/*
		作用域特性：
			A:采用就近原则使用变量
			B:在局部变量作用域范围内 全局变量不起作用 ->值相同的情况下
	*/
	demo2()
	a := 234
	fmt.Println(a)

	const A int = 300 //！常量原则同上
	fmt.Println(A)
}

func demo2() {
	a := 123
	fmt.Println(a)
}
