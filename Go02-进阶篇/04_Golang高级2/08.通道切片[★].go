package main

import (
	"fmt"
	"math/rand"
	"time"
)

//设置值
func setValue(c chan int) {
	c <- rand.Intn(100)
}

func main() {
	rand.Seed(time.Now().UnixNano()) //创建随机数种子
	//var num[5]chan int = [5]chan int{} //声明方法
	var num []chan int = make([]chan int, 5) //声明方法
	//给通道数组赋值
	for i := 0; i < 5; i++ {
		c := make(chan int)
		num[i] = c
		go setValue(c)
	}
	//取值
	for _, v := range num {
		fmt.Println(<-v)
	}
}
