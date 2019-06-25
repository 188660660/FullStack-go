package main

import "fmt"

//父类
type person090201 struct {
	name string
	sex  string
	age  int
}

//学生子类
type student090201 struct {
	person090201
	score int
}

//老师子类
type teacher090201 struct {
	person090201
	subject string
}

/*
	接口的定义
	接口定义了规则 方法实现了规则
	接口是虚的 方法是实的
格式 type 接口名 interface{方法列表}
方法名(参数列表)(返回值列表)
*/
type Humaner090201 interface {
	/*
		 方法 函数声明 没有具体实现
			具体的实现根据对象不同 实现的方式也不一样
			★接口中的方法必须要有具体的实现
	*/
	SayHello090201()
}

func (s *student090201) SayHello090201() {
	fmt.Printf("2大家好，我是%s，我是%s生，我今年%d岁了，我的成绩是%d分\n", s.name, s.sex, s.age, s.score)
}

func (t *teacher090201) SayHello090201() {
	fmt.Printf("1大家好，我是%s，我是%s生，我今年%d岁了，我叫的学科是%s\n", t.name, t.sex, t.age, t.subject)
}

func main090201() {
	var stu = student090201{person090201{"陈浩南", "男", 18}, 23}
	var tea = teacher090201{person090201{"杨永信", "男", 48}, "电磁学"}
	/*
		定义接口类型
			1.接口做了统一的处理 先实现接口 在根据接口实现对应的方法
			2.在需求改变时 不需要改变接口 只需要修改方法
	*/
	var H Humaner090201

	fmt.Printf("%T\n", H) //nil
	H = &stu
	H.SayHello090201()
	fmt.Printf("%T\n", H) //*main.student090201

	H = &tea
	H.SayHello090201()
	fmt.Printf("%T\n", H) //*main.teacher090201

	fmt.Printf("%p\n", H) //0xc0000600c0
	//fmt.Println(H.SayHello090201)

	//直接通过方法名调取方法
	stu.SayHello090201()
	tea.SayHello090201()
}
