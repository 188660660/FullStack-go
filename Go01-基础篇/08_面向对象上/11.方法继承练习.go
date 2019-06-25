package main

import "fmt"

/*
  记者：我是记者  我的爱好是偷拍 我的年龄是34 我是一个男狗仔
  程序员：我叫孙全 我的年龄是23 我是男生 我的工作年限是 3年
*/
//共有的信息
type person081101 struct {
	name   string
	age    int
	sex    string
	career string
}

//记者 结构体
type Rep struct {
	p     person081101
	hoppy string
}

type Dep struct {
	p        person081101
	worktime int
}

func (p person081101) SayHello081101() {
	fmt.Printf("大家好 我是%s 我是%s生我今年%d岁了 很高兴见到大家！\n", p.name, p.sex, p.age)
}

func (r Rep) SayHi() {
	r.p.SayHello081101()
	fmt.Printf("我的兴趣爱好是%s\n\n", r.hoppy)
}

func (d Dep) SayHi() {
	d.p.SayHello081101()
	fmt.Printf("我的工作年限是%d年\n", d.worktime)
}

func main081101() {
	var r = Rep{person081101{"卓伟", 38, "男", "狗仔"}, "偷拍"}
	var d = Dep{person081101{"马化腾", 44, "男", "程序员"}, 3}

	r.SayHi()
	d.SayHi()
}
