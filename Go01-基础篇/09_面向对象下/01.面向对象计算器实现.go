package main

import "fmt"

//父类
type Base struct {
	num1 int
	num2 int
}

//加法子类
type Sum struct {
	Base
}

//实现加法类的结果实现
func (sum *Sub) SumOperate(a int, b int) int {
	sum.num1 = a
	sum.num2 = b
	return sum.num1 + sum.num2
}

//减法子类
type Sub struct {
	Base
}

//实现减法类的结果实现
func (sub *Sub) SubOperate(a int, b int) int {
	sub.num1 = a
	sub.num2 = b
	return sub.num1 - sub.num2
}

func main() {
	var s Sub
	value := s.SubOperate(10, 20)
	fmt.Println(value)
}
