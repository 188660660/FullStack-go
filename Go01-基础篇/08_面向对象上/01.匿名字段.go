package main

import "fmt"

type person0801 struct {
	name string
	sex  string
	age  int
}

//结构体嵌套结构体
type student0801 struct {
	//通过匿名字段实现继承操作
	person0801 //结构体名称作为结构体成员
	id    int
	score int
}

func main080101() {
	var stu = student0801{person0801{"小白", "男", 45}, 102, 666}

	//定义结构体成员信息
	stu.name = "李白"
	stu.sex = "男"
	stu.score = 775
	stu.id = 102
	stu.age = 888

	//结构体名称.分类成员信息
	stu.person0801.name = "李清照"
	stu.person0801.sex = "女"
	stu.person0801.age = 23
	stu.score = 666
	stu.id = 101
	fmt.Println(stu) //{{李清照 女 23} 101 666}
}
