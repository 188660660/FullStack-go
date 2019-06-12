package main

import "fmt"

func main040101() {
	a := 10
	b := 20

	//匿名内部函数 - f1是匿名函数对应的变量
	f1 := func(a int, b int) {
		fmt.Println(a + b)
	}
	f1(a, b)
	f1(30, 40)

	f2 := func() {
		fmt.Println("hello world")
	}
	f2()
	f2()
	f2()
}

func main040102() {
	//带返回值的匿名函数
	a := 10
	b := 20

	f1 := func(a int, b int) int {
		//有返回值 需要定义返回的数据类型
		return a + b
	}
	//函数调用f函数类型 result接受返回值 返回类型为int
	result1 := f1(a, b)
	fmt.Println(result1)
	fmt.Printf("%T\n", result1)

	result2 := func(a int, b int) int {
		return a + b
	}(a, b)
	fmt.Println(result2)

	//方法一：
	func() {
		fmt.Println("传说中的小怪兽")
	}()
	//方法二：
	{
		fmt.Println("传说中的大怪兽")
	}
	//方法一和方法二等同

	//A
	f2 := func() string {
		return "传说中的老怪兽"
	}()
	fmt.Println(f2)
	//B
	{
		fmt.Println("传说中的大头怪")
	}
	//A和B同为 匿名函数 存疑 B如何使用返回值？写法是怎样的呢？此处存疑
}

func main040103() {
	//自己练习 再次熟悉
	a := 10
	b := 20
	//f := func(a int, b int) int {
	//	return a + b
	//}(a, b)

	f := func(a int, b int) int {
		return a + b
	}
	f(a, b)
	fmt.Println(f) //0x491120
	v := f(a, b)   //如果不在匿名函数里面定义 在外部必须还要去赋值一次 才能打印出计算结果
	fmt.Println(v) //输出结果值
}
