package main

import "fmt"

type person struct {
	id    int
	name  string
	score int
	sex   string
}

func test060601(s person) {
	s.name = "刘备"
	fmt.Println(s)             //{101 刘备 9 男}
	fmt.Printf("函数内:%T\n", s)  //main.person
	fmt.Printf("函数内:%p\n", &s) //0xc00007a330
}

func main060601() {
	stu := person{101, "宋江", 9, "男"}

	/*
		结构体作为函数参数 值传递
		函数内部改变不会影响外部的值
	*/
	test060601(stu)

	fmt.Println(stu)             //{101 宋江 9 男}
	fmt.Printf("函数外:%T\n", stu)  //main.person
	fmt.Printf("函数外:%p\n", &stu) //0xc00007a300
}

//所有的切片都属于地址传递
func test060602(stu []person) {
	stu[0].name = "晁盖"
}

func main060602() {
	//结构体切片
	stus := []person{
		{101, "宋江", 9, "男"},
		{101, "宋江", 9, "男"},
		{101, "宋江", 9, "男"},
	}

	//为切片添加信息
	stus = append(stus, person{102, "张飞", 10, "男"})
	fmt.Println(stus) //[{101 宋江 9 男} {101 宋江 9 男} {101 宋江 9 男} {102 张飞 10 男}]
	//结构体切片作为函数参数是地址传递
	test060602(stus)
	//结构体数组作为函数参数是值传递
	fmt.Println(stus) //[{101 晁盖 9 男} {101 宋江 9 男} {101 宋江 9 男} {102 张飞 10 男}]
}
