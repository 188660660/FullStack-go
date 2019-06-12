package main

import "fmt"

//定义全局结构体
type Student060301 struct {
	id   int
	name string
	sex  string
	age  int
}

func main060301() {
	//结构体赋值
	s1 := Student060301{1, "老板", "男", 28}

	//结构体变量赋值
	s2 := s1
	s2.name = "吕不韦"

	/*
		结构体赋值 重新开辟了空间 指向的不是一个内存地址
	*/
	fmt.Println(s1)         //{1 老板 男 28}
	fmt.Printf("%T\n", &s1) //*main.Student060301
	fmt.Printf("%p\n", &s1) //0xc00007a300

	fmt.Println(s2)         //{1 吕不韦 男 28}
	fmt.Printf("%T\n", &s2) //*main.Student060301
	fmt.Printf("%p\n", &s2) //0xc00007a330

	//在结构体中使用 == ！= 可以对结构体成员进行比较操作

	if s1 == s2 {
		fmt.Println("相同")
	} else {
		fmt.Println("不相同")
	}

	//大于小于 可以比较结构体成员
	if s1.age > s2.age {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
