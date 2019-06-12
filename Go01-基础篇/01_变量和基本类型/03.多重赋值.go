package main

import "fmt"

func main0301() {
	//go的多重赋值
	a, b, c, d := 10, 2.1, false, 'b'
	fmt.Println(a, b, c, d)
	//!定义了多少个变量 必须和赋值统一
}

func main0302() {
	//交换 A B的值
	a, b := 10, 20
	//a = b
	//b = a
	//!思考 变量之间可以相互赋值
	//c := a
	//a = b
	//b = c
	a, b = b, a
	fmt.Println(a, b)
}

//匿名变量
func main0303() {
	a, b, _, c := 1, 2, 3, 4
	fmt.Println(a, b, c)
}
