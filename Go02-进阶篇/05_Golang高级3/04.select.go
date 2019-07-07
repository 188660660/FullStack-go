package main

import (
	"fmt"
	"time"
)

func main050401_1() {
	/*
		通过select监听通道
	*/
	// 例题：如果监听到数据通道有值就打印，如果监听到退出开关有值就退出
	c := make(chan int)     // 数据通道
	quit := make(chan bool) // 退出开关
	// 子GO程 写数据
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
			time.Sleep(150 * time.Millisecond)
		}
		close(c)
		quit <- true
	}()
	// 主GO程 读数据
	for {
		select {
		case num := <-c:
			fmt.Println("读取到：", num)
		case <-quit:
			return
		}
		time.Sleep(time.Second) // select只是监听 通道 不会阻塞
	}
}

func main050401_2() {
	// 例题：如果监听到数据通道有值就打印，如果监听到退出开关有值就退出
	c := make(chan int)
	quit := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
			time.Sleep(150 * time.Millisecond)
		}
		quit <- true
	}()

	for {
		select {
		case num := <-c:
			fmt.Println(num)
		case <-quit:
			return
		}
	}
}

func main050401_3() {
	// 例题：如果监听到数据通道有值就打印，如果监听到退出开关有值就退出
	c := make(chan int) // 开启通道 保存值
	quit := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(300 * time.Millisecond)
			c <- i
		}
		quit <- true
	}()

	for {
		select {
		case num := <-c:
			fmt.Println(num)
		case <-quit:
			return
		}
	}
}
func main() {
	// 例题1：将斐波那契数列写入通道并读取	描述：在通道中输入数列，并读取。

}