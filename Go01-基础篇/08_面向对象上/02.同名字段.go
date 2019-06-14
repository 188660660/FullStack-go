package main

import "fmt"

type person0802 struct {
	name string
	age  int
	sex  string
}

type student0802 struct {
	person0802
	id   int
	name string
}

func main080201() {
	var stu = student0802{person0802{"武器大师", 39, "男"}, 888, "爆破鬼才"}

	//采用就近原则 使用子类的信息
	stu.name = "幽灵冥主"

	//结构体变量.匿名字段.结构体成员
	stu.person0802.name = "天神剑灵"

	fmt.Println(stu)
}
