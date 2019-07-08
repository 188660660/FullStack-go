package main

import "fmt"

// 一个通道在一个GO程 读写
func main1501() {
	/*
			死锁是指两个或两个以上协程在执行过程中 由于竞争资源或由于彼此通信而造成的一种
		阻塞的现象 若无外力作用 他们将无法进行下去
	*/
	c := make(chan bool)
	go func() {
		c <- true
	}()
	fmt.Println(<-c)
}

// GO程在通道开启之前使用
func main1502() {
	c := make(chan int)
	c <- 15
	go func() {

		fmt.Println(<-c)
	}()
}

// 通道1中调用了通道2 通道2中调用了通道1
