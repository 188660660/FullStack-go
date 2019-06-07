package main

import "fmt"

func main0401() {
	//iota是什么？ iota常量生成器初始化
	//iota 会从0 开始 然后在每一个有常量声明的行加一
	const (
		a = iota //状态A
		b = iota //状态B
		c = iota //状态C
		d = iota //状态D
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	//定义常量为状态
	value := a
	fmt.Println(value)
	fmt.Printf("%T\n", value)
	value = b
	fmt.Println(value)
}

/*
总结：
枚举一般用作于程序的流程控制
const(
  a=iota //0
  b,c=iota,iota//1
  d=20//2
)
*/
