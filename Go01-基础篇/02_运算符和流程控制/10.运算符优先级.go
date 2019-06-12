package main

import "fmt"

/*
	!总结：
		运算优先级：
			A.括号() 结构体成员.  数组下标[]

			B.单目运算符
			C.逻辑非! 取地址& 取值* 自增++ 自减--

			D.双目运算符
			E.乘除 * / %
			F.加减 + -
			G.关系 == != > >= < <=
			H.逻辑 || &&
			I.赋值 = += -= *= /= %=
*/
func main021001() {
	a := 10
	b := 20
	c := 30

	var d int
	//d = a + b * c
	//d = (a + b) * c

	fmt.Println(d)
	fmt.Println(a+b >= c && !(b > c))
}

func main021002() {
	/*
		判断闰年小案例
		闰年的判断条件：
			1、能够被400整除
			2、能够被4整除 不能被100整除
	*/
	var year int
	fmt.Println("请输入需要判断的年份！")
	fmt.Scan(&year)
	result := year%400 == 0 || (year%4 == 0 && !(year%100 == 0))
	fmt.Println(result) //返回的运算结果值为布尔值 true || false
}
