package main

import "fmt"

func main020301() {
	//常量的定义和使用
	//在程序运行过程中 其值不能发生改变的量 称为常量
	//变量存储在栈区 而常量一在内存中的数据区
	//栈区 系统为每一个应用程序分配1M空间来存储变量 在程序结束运行后系统会自动释放
	var s1 = 10
	var s2 = 30
	//！常量的存储位置在数据区 不能通过& 取地址来访问
	const a int = 10

	fmt.Printf("%d\n", &s1)
	fmt.Printf("%d\n", &s2)

	fmt.Println(&s1)
	fmt.Println(&s2)

	//a = 20 常量的值不能修改 修改会报错
	fmt.Println(a)
}

func main020302() {
	//常量一般用于大写字母表示
	const MAX = 999
	b := 20
	c := MAX + b
	fmt.Println(c)

	//字面常量 不入住内存区
	fmt.Println(123)
	fmt.Println("hello world")

	//硬常量
	d := c + 32
	e := "hello"
	e += "world"
	fmt.Println(d)
	fmt.Println(e)

	//！再次理解下什么是 硬常量
	//所谓字面常量（literal），是指程序中硬编码的常量，
}
