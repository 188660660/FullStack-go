package main

import "fmt"

func test100201() {
	fmt.Println("测试函数A")
}

func test100202() {
	fmt.Println("测试函数B")
	//可以在程序中 直接调用panic 调用后程序会终止执行
	//panic("测试函数B")
	/*
		print() println() 和 fmt.print fmt.println 的区别
		print 和 println 是对错误信息进行处理的 不建议在程序中打印数据使用
	*/
}

func test100203() {
	fmt.Println("测试函数C")
}

func main100201() {
	test100201()
	test100202()
	test100203()
}
