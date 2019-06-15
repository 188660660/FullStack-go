package main

import "fmt"

//父类
type Person_2 struct {
}

//子类
type student090301_2 struct {
	Person_2
}

type Api interface {
	show()
}

func (p *Person_2) show() {
	fmt.Println("人类")
}

func (p *student090301_2) show() {
	fmt.Println("学生类")
}

//多态是将接口类型作为函数的参数 多态实现了 函数的统一处理
func display(i Api) {
	i.show()
}

func main090302_2() {
	//多态 多种形态 相同的名字但形态不一样
	var i Api = &Person_2{}
	display(i)
	i = &student090301_2{}
	display(i)
}
