package main

import (
	"fmt"
	"time"
)

func main04050101() {
	//通道关闭以后就不在阻塞
	var c = make(chan int)
	go func() {
		c <- 10
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}

//通道关闭以后不能再写入 但是可以继续进行读取
func main040501() {
	var c = make(chan int, 3)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	time.Sleep(10 * time.Second)
	fmt.Println(len(c))
	for v := range c {
		fmt.Println(v)
	}
	/*
		close()不是必须的，如果读取端循环读取数据必须要调用close()，否则会出现死锁
		只有写端才能关闭
	*/
}
