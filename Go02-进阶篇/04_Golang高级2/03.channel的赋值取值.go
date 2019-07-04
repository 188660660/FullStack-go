package main

import (
	"fmt"
)

func main() {
	var c = make(chan int)
	var ec = make(chan bool)
	go func() {
		for i := 0; i <= 5; i++ {
			c <- i //创建阻塞
			fmt.Println("写入 ", i)
			ec <- true
		}
	}()
	for i := 0; i <= 5; i++ {
		num := <-c
		fmt.Println("读取：", num)
		<-ec //解除阻塞
	}
	//time.Sleep(3*time.Second)
}
