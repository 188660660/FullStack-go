package main

import "fmt"

func main040201_2() {
	/*
		1.闭包是函数中的匿名函数
		2.形成闭包的条件必须内部函数要返回出来
		3.每个闭包都独立保存自己的信息
	*/
	fun := fun1() //函数的调用
	fun()
}

func fun1() func() {
	var add = "上海"
	fun2 := func() { //匿名函数就是闭包
		fmt.Println(add)
	}
	return fun2 //闭包形成的条件将匿名函数返回出来
}

//再次练习
func main040202_2() {
	f := fun2()
	f()
}

func fun2() func() {
	a := "中国"
	fun := func() {
		fmt.Println(a)
	}
	return fun
}

func main040203_2() {
	fmt.Println(fun3()) //1
	fmt.Println(fun3()) //1
	fmt.Println(fun3()) //1
}

func fun3() int {
	x := 0
	x++
	//函数执行完毕后 局部变量即销毁 所以此处每次的存储空间都是不同的
	fmt.Printf("%p\n", &x) //表示为十六进制，并加上前导的0x &符为取址
	return x * x
}

func main040204_2() {
	fmt.Println(fun4()) //1
	fmt.Println(fun4()) //1
	fmt.Println(fun4()) //1

	fun := fun4
	fmt.Println(fun())
	fmt.Println(fun())
	fmt.Println(fun())
}

func fun4() func() int {
	//验证和视频出的结果是否一致：

	//return func() int {
	//	x := 0
	//	x++
	//	//fmt.Printf("%p\n", &x)
	//	return x * x
	//}

	var x = 0
	return func() int {
		x++
		return x * x
	}

	//闭包每次指向的内存空间都是一样的
}

func main040205_2() {
	var fn [3]func() //函数数组
	for i := 0; i < 3; i++ {
		fn[i] = func() {
			fmt.Println(i)
		}
	}
	//每个调用都是打印3 原因是i变量只有一个 被三个函数共同访问
	fn[0]() //3
	fn[1]() //3
	fn[2]() //3

	for i := 0; i < 3; i++ {
		fn[i] = make_fun(i)
	}

	fn[0]() //0
	fn[1]() //1
	fn[2]() //2
	//闭包可以保存自己独有的信息
	//此处的运行结果不是很理解
}

func make_fun(i int) func() {
	return func() {
		fmt.Println(i)
	}
	/*
		！总结：
			闭包就像气球 没调用一次就产生一个气球
			每个气球中的气体可以不一样
	*/
}
