package main

import "fmt"

func main0101() {
	/*
		!语法：
			for  表达式1 ; 表达式2 ; 表达式3{
				循环体
			}
	*/

	//循环1-10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	//计算1-100的和
	//sum := 0
	//for i:=0;i<=100;i++  {
	//	sum += i
	//}
	//fmt.Println(sum)

	//计算0-100之间偶数的和 - 方法一
	//sum := 0
	//for i:=0;i<=100;i+=2  {
	//	sum += i
	//}
	//fmt.Println(sum)

	//计算0-100之间偶数的和 - 方法二
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%2 == 0 { //在循环中嵌套条件判断语句
			sum += i
		}
	}
	fmt.Println(sum)
}

func main0102() {
	//for循环中变量的作用域
	/*
		!注意：
			for函数格式 i是在函数内部定义的就不能在外部进行使用
	*/
	for i := 1; i <= 100; i++ {
		fmt.Println(i) //局部变量
	}

	//fmt.Println(i) //err

	//go语言中死循环的概念
	for {
		//在程序的编写中 要尽量避免死循环 在此处 我们可以使用break跳出当前的循环
		fmt.Println("hello！")
		break
	}

	//这种写法有点类似于PHP的while循环的写法
	i := 1
	for {
		if i > 10 {
			break
		}
		fmt.Println(i)
		i++
	}
}
