package main

import "fmt"

type person081301 struct {
	name string
	age  int
	sex  string
}

type student081301 struct {
	person081301
	score int
}

func (p *person081301) PrintInfo081301() {
	fmt.Printf("大家好，我是%s，我今年%d岁，我是%s生\n", p.name, p.age, p.sex)
}

func (s *student081301) PrintInfo081301(name string) {
	s.name = name
	fmt.Printf("大家好，我是%s，我今年%d岁，我是%s生,我的分数是%d分\n", s.name, s.age, s.sex, s.score)
}

func hello() {
	fmt.Println("hello")
}

type funcdemo func()

func main081301() {
	s := student081301{person081301{"张三", 11, "男"}, 19}

	//存储在代码区
	fmt.Println(s.PrintInfo081301)              //0x4932c0
	fmt.Println(s.person081301.PrintInfo081301) //0x4932c0

	fmt.Println(hello) //0x492e20

	var f1 funcdemo
	//f1 = s.PrintInfo081301
	//fmt.Println(f1) //0x493470
	fmt.Printf("%T\n", f1) //main.funcdemo

	f := hello
	fmt.Printf("%T\n", f) //func()
	f()                   //hello

	a := 10
	fmt.Printf("%p\n", &a) //0xc000062088

	//方法类型(函数类型) 变量
	f3 := s.PrintInfo081301

	f3("神王")
	fmt.Println(s)
	fmt.Printf("%T\n", f3) //func(string)
}
