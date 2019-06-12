package main

import "fmt"

/*
	!函数：
		语法：
			func 函数名(参数类型 参数){
				代码体 程序体 函数体
			}
		A: 如果有多个参数 使用,号进行隔开
		B: 函数需要先定义后调用
			函数的定义：定义过程中 定义的参数称为形参 传过来的参数叫做实参
		C：函数可以重复调用
*/
func test1(a int, b int) {
	sum := a + b
	fmt.Println(sum)
}

func main030701() {
	a := 10
	b := 20
	test1(a, b)
	test1(35, 82)
}

func main030702() {
	/*
		!重点：形参和实参是不同的存储单元 在函数的调用过程中 不会互相产生影响
		函数的调用使用的是实参 函数的定义中使用的是形参
	*/
	a, b := 85, 117
	test2(a, b)
	//此处的值会发生改变吗？不会：因为是两个不同的存储单元
	fmt.Println(a, b)
}

func test2(a int, b int) {
	//案例:交换两个数的值
	a, b = b, a
	fmt.Println(a, b)
}
