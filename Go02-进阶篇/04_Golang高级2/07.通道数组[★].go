package main

import (
	"fmt"
	"math/rand"
	"time"
)

//设置值
func setValue10(c chan int) {
	c <- rand.Intn(100)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var num []chan int = make([]chan int, 0)
	for i := 1; i <= 5; i++ {
		c := make(chan int)
		num = append(num, c) //将通道添加到切片中

	}
	go func() {
		for _, v := range num {
			v <- rand.Intn(100)
		}
	}()
	//	取值
	for _, v := range num {
		fmt.Println(<-v)
	}
}
