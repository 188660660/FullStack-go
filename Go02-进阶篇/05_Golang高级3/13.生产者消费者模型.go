package main

import (
	"fmt"
)

func Producer(c chan<- int) {
	// 生产者
	for i := 0; i < 5; i++ {
		c <- i
		fmt.Println("生产：", i)
	}
	close(c)
}

func Consumer(c <-chan int) {
	// 消费者
	for v := range c {
		fmt.Println("消费：", v)
	}
}

func main1301() {
	// 生产者和消费者模型
	c := make(chan int, 2) // 容量越大 并发性越好
	go Producer(c)
	Consumer(c)
}
