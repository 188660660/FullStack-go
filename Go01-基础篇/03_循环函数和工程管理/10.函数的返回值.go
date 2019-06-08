package main

import "fmt"

func sub1(a int, b int) (sub int) {
	sub = a - b
	return
}

//func sub2(a int, b int) (sub int) { A 不报错
//func sub2(a int, b int) (int) {	B 不报错
func sub2(a int, b int) int {
	//return 可以写一个表达式 将表达式的结果作为返回值
	return b - a
}

func sub3(a int, b int) int {
	/*
		!此处注意下 和sub1的区别
		sub1中没有使用推导类型
		而此处使用了 推导类型就需要返回类型结果 否则报错
	*/
	sum := a + b
	return sum
	/*
		return 表示函数的结束 return后面的代码不会执行
		同时return 也会将函数的返回值传递给主调函数
	*/
	fmt.Println("hello world")
	fmt.Println("hello world")
	fmt.Println("hello world")
	return 111
}

func main1001() {
	var value1, value2, value3 int

	value1 = sub1(10, 20)
	value2 = sub2(10, 20)
	value3 = sub3(10, 20)

	fmt.Println(value1)
	fmt.Println(value2)
	fmt.Println(value3)
}

//函数返回多个值
func main1002() {
	//多重赋值
	a, b, c := test9()

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//如果函数有多个返回值 但是不使用返回值 此处可以使用匿名变量接收数据
	d, _, f := test9()

	fmt.Println(d)
	//fmt.Println(e)
	fmt.Println(f)
}

//函数没有参数 在函数调用时()也必须要写 test9()的这个小括号 语法需要
func test9() (a int, b int, c int) {
	a, b, c = 10, 20, 30
	return
}
