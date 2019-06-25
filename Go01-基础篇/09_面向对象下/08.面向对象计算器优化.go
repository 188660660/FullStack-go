package main

import "fmt"

//定义结构体 struct不是class 但在golang中 可以当做类来理解使用
type Base struct {
	//定义父类 存放两个数值
	num1 int //数值一
	num2 int //数值二
}

//定义加法计算结构体
type Sum struct {
	Base //继承父类
}

//定义减法计算结构体
type Sub struct {
	Base //继承父类
}

//定义一个接口实现运算方法
type Api interface {
	Calculation() int //声明一个计算方法
}

func Duotai(api Api) (value int) {
	value = Api.Calculation()
	return value
}

//实现接口定义的方法 方法也是一种函数 不过要和对象进行绑定
func (sub *Sub) Calculation(num1 int, num2 int) (value int) {
	return sub.num1 - sub.num2
}

func (sum *Sum) Calculation(num1 int, num2 int) (value int) {
	return sum.num1 + sum.num2
}

type Factory struct {
	//定义工厂对象 再通过函数进行实现
}

//实现计算方法
func (cal *Factory) Factory(num1 int, num2 int, param string) (value int) {
	switch param {
	case "+": //如果为+ 执行加法运算操作
		var value = Sum{Base{num1, num2}} //创建对象
		value = Sum.Calculation()
	case "-": //如果为- 执行减法操作
		var value = Sub{Base{num1, num2}}
		value = Sub.Calculation()
	}
	return
}

func main() {
	var factory Factory
	value := factory.Factory(10, 20, "+")
	fmt.Println(value)
}
