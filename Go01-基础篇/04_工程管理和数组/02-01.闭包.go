package main

import "fmt"

func test1(a int) {
	a++
	fmt.Println(a)
}

func main040201_1() {
	/*
		!总结： 什么是闭包？
	*/
	a := 0
	for i := 0; i <= 5; i++ {
		//函数会在调用结束后从内存中销毁
		test1(a)
	}
	fmt.Println(a)
}

func test2() func() int {
	//可以通过匿名函数和闭包 实现函数在栈区的本地化
	var a int
	return func() int {
		a++
		return a
	}
}

func main040202_1() {
	////将test2函数类型赋值给f
	//f := test2
	f := test2() //函数调用 将test2的返回值赋值给test2
	for i := 0; i < 5; i++ {
		fmt.Println(f())
	}
	fmt.Printf("%T", f) //func() int
}
