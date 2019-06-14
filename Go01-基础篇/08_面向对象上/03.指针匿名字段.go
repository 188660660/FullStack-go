package main

import "fmt"

type person0803 struct {
	name string
	age  int
	sex  string
}

type student0803 struct {
	*person0803 //指针匿名字段
	id          int
	score       int
}

func main080301() {
	var stu = student0803{&person0803{"张飞", 18, "男"}, 250, 125}
	fmt.Println(stu) //{0xc00007a300 250 125}

	stu.person0803.name = "李逵"
	fmt.Println(stu) //{0xc00007a300 250 125}

	per := person0803{"恶魔男爵", 180, "未知"}
	stu.person0803 = &per

	//将结构体变量 赋值给结构体指针类型
	//stu.name = "张三"
	stu.person0803 = new(person0803)
	stu.person0803.name = "超级大怪兽"
	fmt.Println(stu) //{0xc00007a300 250 125}

	fmt.Println(stu.name)  //张飞
	fmt.Println(stu.id)    //250
	fmt.Println(stu.score) //125
}
