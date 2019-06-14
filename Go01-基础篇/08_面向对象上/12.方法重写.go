package main

import "fmt"

type person081201 struct {
	name string
	age  int
	sex  string
}

type student081201 struct {
	person081201
	score int
}

func (s person081201) PrintInfo081201() {
	fmt.Printf("姓名：%s | 年龄：%d | 性别：%s | 成绩：%d\n", s.name, s.age, s.sex)
}

/*
	方法重写 在一个对象中不能出现相同的方法名
	方法的接收者 带* 和不带* 表示一个相同的对象(本文指 student081201)
*/
func (s student081201) PrintInfo081201() {
	fmt.Printf("姓名：%s | 年龄：%d | 性别：%s | 成绩：%d\n", s.name, s.age, s.sex, s.score)
}

//方法和函数可以重名 但方法不可以同时出现两个重名的
func main081201() {
	t := student081201{person081201{"神王", 999, "未知"}, 888}
	//默认使用子类的方法 采用就近原则
	//调用子类的方法
	t.PrintInfo081201()                   //姓名：神王 | 年龄：999 | 性别：未知 | 成绩：888
	fmt.Printf("%p\n", t.PrintInfo081201) //0x493230
	//调用父类的方法
	t.person081201.PrintInfo081201()                   //姓名：神王 | 年龄：999 | 性别：未知 | 成绩：%!d(MISSING)
	fmt.Printf("%p\n", t.person081201.PrintInfo081201) //0x4932b0
}
