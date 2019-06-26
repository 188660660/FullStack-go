package main

import "fmt"

func main020201() {
	//var a int64 = 10
	a := 10 //int
	fmt.Printf("%d\n", a)
	fmt.Printf("%T\n", a)

	b := 10.0
	fmt.Printf("%f\n", b)
	fmt.Printf("%T\n", b)

	var c = true //布尔值默认为false
	fmt.Printf("%t\n", c)
	fmt.Printf("%T\n", c)

	var d1 = "张三丰"
	d2 := "asas"
	fmt.Printf("%s\n", d1)
	fmt.Printf("%s\n", d2)

	var e1 byte = 'a'
	e2 := 'A'
	//%v可以输出推导类型的值
	fmt.Printf("%v\n", e1)
	fmt.Printf("%v\n", e2)

	var f string
	fmt.Println("请您输入任意值")
	fmt.Scanf("%p\n", &f)
	fmt.Printf("%s\n", f)
	//[★]此处存疑 为什么我的输入值无法正常获取到

	//%%会打印出一个%
	fmt.Printf("35%%")
}

func main020202() {
	/*问题：计算机都能够识别哪些进制？
	答：二进制 八进制 十进制 十六进制
	*/
	a := 123   //十进制
	b := 0123  //八进制 标识：以0开头的
	c := 0xabc //十六进制 标识：w
	//go语言中 二进制不能直接标识

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//%b占位符 表示输出一个二进制数据
	fmt.Printf("%b\n", a)
	fmt.Printf("%b\n", b)
	fmt.Printf("%b\n", c)

	//%o占位符 表示输出一个八进制数据
	fmt.Printf("%o\n", a)
	fmt.Printf("%o\n", b)
	fmt.Printf("%o\n", c)

	//%x %X 占位符表示输出一个十六进制数据
	fmt.Printf("%x\n", a)
	fmt.Printf("%x\n", b)
	fmt.Printf("%x\n", c)

	fmt.Printf("%X\n", a)
	fmt.Printf("%X\n", b)
	fmt.Printf("%X\n", c)
	//%x和%X的区别：大写和小写的区别

	s := ' ' //字节型的默认值为int32
	fmt.Printf("%T", s)
}
