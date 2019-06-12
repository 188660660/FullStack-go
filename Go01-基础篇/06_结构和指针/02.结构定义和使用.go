package main

import "fmt"

/*
	在函数外部定义结构体 作用域是全局的
	语法：
		type 结构体名 struct{
			结构成员列表
		}
*/
type Student060201 struct {
	id   int
	name string
	sex  string
	age  int
	addr string
}

func main060201() {
	//通过结构体名 定义结构变量
	var s1 Student060201
	//结构体变量名.成员名
	s1.id = 125
	s1.name = "叶凡"
	s1.sex = "男"
	s1.age = 225
	s1.addr = "华夏古龙区"

	fmt.Println(s1)        //{125 叶凡 男 225 华夏古龙区}
	fmt.Printf("%T\n", s1) //main.Student

	//两种结构体的定义方式
	var s2 Student060201 = Student060201{101, "赵四", "女", 180, "华夏大区"}
	//s := Student060201{101,"赵四","女",180,"华夏大区"}

	fmt.Println(s2.id)
	fmt.Println(s2.name)
	fmt.Println(s2.sex)
	fmt.Println(s2.age)
	fmt.Println(s2.addr)
}
