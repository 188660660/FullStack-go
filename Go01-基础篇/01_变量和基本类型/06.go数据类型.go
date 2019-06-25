package main

import "fmt"

func main0601() {
	//布尔值数据
	var a bool
	a = true
	fmt.Println(a)
	b := false
	fmt.Printf("%T", b)
	//布尔数据类型的默认值是false

}

func main0602() {
	//浮点型数据类型
	var a float32
	var b float64

	a = 1.111111111111111111111111111 //输出结果总结：float32只能保存小数点后七位 且进行+1操作
	b = 2.222222222222222222222222222 //输出结果总结：float64可以保存小数点后十五位 此处有点类似于PHP的单精度和多精度浮点数

	fmt.Println(a)
	fmt.Println(b)
}

func main0603() {
	//字符数据类型
	var a byte = '0'
	var b byte = 'a'
	//使用fmt.Println打印字符型会显示出来 向对应的ASCII码值
	fmt.Println(a)
	fmt.Println(b)

	fmt.Printf("%v=%v", a, b)
}

func main0604() {
	//字符串数据类型
	var a = "夏天"
	var b = "炎热"
	c := a + b
	fmt.Println(a + b)
	fmt.Printf("%s\n", c)
	fmt.Printf("%T\n", c)
	//==运算符 可以判断两个字符串是否相等
	fmt.Println(a == b)
}
