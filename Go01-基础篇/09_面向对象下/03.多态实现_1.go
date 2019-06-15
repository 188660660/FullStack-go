package main

import "fmt"

type person090301_1 struct {
	name string
	sex  string
	age  int
}

type student090301_1 struct {
	person090301_1
	score int
}

type teacher090301_1 struct {
	person090301_1
	subject string
}

func (s *student090301_1) SayHello() {
	fmt.Printf("大家好，我是%s，我是%s生，我今年%d岁了，我的成绩是%d分\n", s.name, s.sex, s.age, s.score)
}
func (t *teacher090301_1) SayHello() {
	fmt.Printf("大家好，我是%s，我是%s生，我今年%d岁了，我叫的学科是%s\n", t.name, t.sex, t.age, t.subject)
}

//接口实现
type Personer_1 interface {
	SayHello()
}

//多态实现
//多态是将接口类型 作为函数参数 实现接口的统一处理
func SayHello(p Personer_1) {
	p.SayHello()
}

func main090301_1() {
	var P Personer_1 //声明接口变量
	P = &student090301_1{person090301_1{"小红", "女", 10}, 90}
	SayHello(P)

	P = &teacher090301_1{person090301_1{"一泓", "女", 18}, "心里辅导"}
	SayHello(P)
}
