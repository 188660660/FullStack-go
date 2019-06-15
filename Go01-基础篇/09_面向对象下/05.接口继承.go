package main

import "fmt"

//定义接口
type Humaner10905 interface {
	//子集
	SayHi()
}

type Person0905 interface {
	//超集 ： 一组子集的集合
	Humaner10905
	Sing(string)
}

//创建对象
type student0905 struct {
	name string
	age  int
	sex  string
}

func (s *student0905) SayHi() {
	fmt.Printf("大家好，我是%s，我今年%d岁了，我是%s生\n", s.name, s.age, s.sex)
}

func (s *student0905) Sing(name string) {
	fmt.Printf("大家好，我叫%s,我给大家唱一首：%s\n", s.name, name)
}

func main090501() {
	//子集的使用
	var T Humaner10905
	T = &student0905{"韩红", 40, "女"}
	T.SayHi()

	//超集的使用
	var p Person0905
	p = &student0905{"王菲", 18, "女"}
	p.SayHi()
	p.Sing("传奇")

	//将超集转成子集
	T = p
	//p = T //err
	p.Sing("搜索")
	p.SayHi()
}
