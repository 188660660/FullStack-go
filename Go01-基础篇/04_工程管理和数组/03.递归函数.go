package main

import (
	"fmt"
)

func test0401(a int) {
	//在递归中如果没有结束条件称为死递归
	fmt.Println(a)
	test0401(a)
}

func test0402(a int) {
	if a == 0 {
		return //return程序停止不继续执行
	}
	a--
	fmt.Println(a)
	test0402(a)
}

func main040301() {
	/*
		概念：在函数调用时调用函数本身 称为递归函数
		在递归函数中必须要有一个出口 不然会一直执行
	*/
	//死递归
	//test0401(10)
	test0402(10)
}

var sum1 = 1 //0*任何数都等于0 因此从1开始相乘

func test0403(n int) {
	//计算n的阶乘 n=1×2×3×...×n
	if n == 1 {
		return
	}
	test0403(n - 1)
	sum1 *= n
}

func main040302() {
	test0403(3)
	fmt.Println(sum1)
}

//再写一遍
var count = 0
var sum2 = 1 //0*任何数都等于0 因此从1开始相乘
func test0404(n int) {
	if n == 1 {
		return
	}
	test0404(n - 1)
	sum2 *= n
	count++
}

func main040303() {
	test0404(4)
	fmt.Println(sum2)
	fmt.Printf("程序执行了：%d次", count)
}

/*
	递归比较消耗资源 尽量保持少用原则
*/

func test0405_1(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	test0405_1(n - 1)
	n--
}

func test0405_2(n int, e int) {
	if e > n {
		return
	}
	fmt.Println(e)
	test0405_2(n, e+1)
	e++
}

var sum = 0

func test0406_1(num int) {
	if num == 0 {
		return
	}
	test0406_1(num - 1)
	sum += num
	fmt.Println(sum)
}
func test0406_2(num int) int {
	if num == 1 {
		return 1
	}
	return num + test0406_2(num-1)
}

func main040304() {
	//练习题1：依次打印 1,2,3,4,5...
	//test0405_1(20) //倒序打印
	test0405_2(10, 0) //正序打印

	//练习题2：从1至100 依次相加
	test0406_1(100) //方法一
	//★优化
	fmt.Println(test0406_2(100)) //方法二
}
