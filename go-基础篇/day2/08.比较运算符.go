package main

import "fmt"

func main0801() {
	//比较运算符又作关系运算符 返回值为布尔型数据
	a := 10
	b := 20

	fmt.Println(a == b) //返回false == 相等于 相同于
	fmt.Println(a > b)  //大于 返回false
	fmt.Println(a < b)  //大于 返回true

	//>= 大于等于运算符 <=运算符
	fmt.Println(a >= b) //大于等于 返回false
	fmt.Println(a <= b) //小于等于 返回true

	// !=
	fmt.Println(a != b) //返回true
	/*
		!总结：
			在golang语言中有!== 补全等于吗?
			经查询 go 中没有不全等于比较运算符
			只有 == != > < >= <=
	*/
	//c := 'a'
	//d := "A"
	//fmt.Println(c!==d)
}
