package main

import (
	"fmt"
	"math"
	"time"
)

func main020201() {
	var num1 int = 25
	var num2 int = 2
	fmt.Println(num1 / num2)
	/*
		1.不同的数据类型之间不能进行运算
		2.两个整形之间相除 还是一个整型
		3.字符串想加就是字符串相连
		4.在go中没有前自增 只有后自增
		5.自增不能用于表达式中
		6.GO语言中没有三元运算符
	*/

}
func main020204() {
	a, b := 10, 20
	fmt.Printf("%d+%d=%d\n", a, b, a+b)
	fmt.Printf("%d-%d=%d\n", a, b, a-b)
	fmt.Printf("%d*%d=%d\n", a, b, a*b)
	fmt.Printf("%d/%d=%d\n", a, b, a/b) //整型和整型运算结果还是整型
	a++                                 //go只支持后自增，不能将自增应用于表达式中
	fmt.Printf("a++=%d", a)
	b--
	fmt.Printf("b--=%d", b)

	a, b = 10, 0
	//如果是短路或 第一个为true 第二个不执行
	//如果是短路与 第一个为好了 第二个不执行
	if a > 0 && a/b > 0 {
		fmt.Println("OK")
	}
}

func main() {
	fmt.Println("请输入您的考试成绩！")
	var score int
	fmt.Scan("%d", &score)
	switch {
	case score > 90:
		fmt.Println("考试成绩优秀")
	case score < 80:
		fmt.Println("考的忒差了！")
	}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(math.Pi)
	fmt.Println(time.Now().Year(), "-", time.Now().Month())

}
