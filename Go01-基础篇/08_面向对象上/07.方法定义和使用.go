package main

import "fmt"

//定义结构体
type stu080701 struct {
	name string
	age  int
	sex  string
}

func (a stu080701) printinfo() {
	fmt.Println(a.name)
	fmt.Println(a.sex)
	fmt.Println(a.age)
}

func (a *stu080701) editinfo(name string, age int, sex string) {
	//(*s).name = name
	//结构体指针可以直接调用结构体成员
	a.name = name
	a.age = age
	a.sex = sex
}

func main080701() {
	var s = stu080701{"小明", 14, "男"}
	fmt.Println(s)

	//对象.方法
	s.name = "李小飞"
	s.printinfo() //可以不传参数

	var stu080702 = stu080701{"小红", 16, "女"}
	stu080702.printinfo()

	//结构体变量 可以直接使用结构体指针对应的方法
	(&stu080702).editinfo("小红帽", 14, "女")
	//stu080702.editinfo("小红",14,"女")
	stu080702.printinfo()
}

func main080702() {
	var s1 *stu080701
	s1 = new(stu080701)

	s1.editinfo("刘能", 14, "女")
	s1.printinfo()
}
