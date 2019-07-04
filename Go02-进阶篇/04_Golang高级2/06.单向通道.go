package main

import "fmt"

/*
	隐式转换和显示转换的区别
	显式转换是由程序员自己主动完成，隐式转换是由编译器完成的
*/
func main040601() {
	/*
		1.默认情况下 通道是双向的
		2.双向通道 可以隐式转换成单向通道
		3.通道的传递是地址传递
			语法：
				双向通道 chan int
				写通道 chan<- int
				读通道<-chan int
	*/
	c := make(chan int)
	go func() {
		var c_w chan<- int = c
		c_w <- 10
	}()
	var c_r <-chan int = c
	num := <-c_r
	fmt.Println(num)
}

func main040602() {
	c := make(chan int)
	go func() {
		c_w := make(chan<- int) //自动推导声明写通道
		c_w = c                 //给单向通道赋值
		c_w <- 20               //将数字写入到写通道中
	}()
	c_r := make(<-chan int)
	c_r = c
	num := <-c_r
	fmt.Println(num) //20
}

//定义单向通道-方法一练习
func main040603() {
	c := make(chan int)
	go func() {
		var c_w chan<- int = c
		c_w <- 10
	}()
	var c_r <-chan int = c
	num := <-c_r
	fmt.Println(num)
}

//定义单向通道-方法一练习
func main040604() {
	c := make(chan int)
	go func() {
		c_w := make(chan<- int)
		c_w = c
		c_w <- 20
	}()
	c_r := make(<-chan int)
	c_r = c
	fmt.Println(<-c_r)
}

//单向通道作为函数参数
func test_w(c chan<- int) { //写通道作参数
	c <- 205
}

func test_r(c <-chan int) { //读通道作参数
	fmt.Println(<-c)
}

func main040605() {
	c := make(chan int) //双向通道
	go test_w(c)
	test_r(c)
}
