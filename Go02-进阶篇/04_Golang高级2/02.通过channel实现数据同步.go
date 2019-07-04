package main

import (
	"fmt"
	"time"
)

var c = make(chan bool)  //声明通道 用来控制用户的执行顺序
var ec = make(chan bool) //声明通道 用来控制用户的执行顺序
func printer(str string) {
	for _, v := range str {
		fmt.Printf("%c", v)
		time.Sleep(time.Millisecond * 300)
	}
}

func user1() {
	printer("hello")
	c <- true //写阻塞
}
func user2() {
	<-c //读阻塞 等到另外一个go程往c中写值 等到user1写完才接触阻塞
	printer("world")
	ec <- true //写阻塞
}

func main040201() {
	go user1()
	go user2()
	<-ec //一致阻塞 一直到user2执行完毕
	//思考 ： 如何去掉for循环？
	//for  {
	//	;
	//}
}
