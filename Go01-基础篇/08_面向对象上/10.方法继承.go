package main

import "fmt"

type person081001 struct {
	id   int
	name string
	age  int
	sex  string
}

type student081001 struct {
	p person081001 //结构体变量 结构体类型

	score int
}

func (p *person081001) SayHello081001() {
	//父类不能继承子类信息
	fmt.Print("大家好 我叫", p.name, "今年", p.age, "岁了", "我是", p.sex, "性 我的编号是", p.id, "\n")
}

func main081001() {
	var stu student081001 = student081001{person081001{1, "铸星龙王", 18, "男"}, 600}
	var p person081001 = person081001{2, "黑暗暴君", 36, "男"}

	//★ 父类不能继承子类信息
	p.SayHello081001()

	stu.p.id = 2
	stu.p.name = "光明神君"
	stu.p.sex = "男"
	stu.p.age = 999

	//子类结构体继承父类结构体 允许使用父类结构体成员

	stu.p.SayHello081001()

	fmt.Println(stu)
}
