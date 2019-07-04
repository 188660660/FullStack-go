package main

import "fmt"

func main030501() {
	c := make(chan bool)
	go func() {
		//x := c<-true
		c <- true
		c <- false
		c <- true
	}()
	fmt.Println(<-c)
}

func main() {
	//实现子go程执行完毕，主线程退出
	c := make(chan bool)
	go func() {
		fmt.Println("子程实现完毕")
		c <- true //写入阻塞
	}()
	//<-c
	fmt.Println("main function")
	//fmt.Println(<-c)
	<-c
}
