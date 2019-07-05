package main

import (
	"fmt"
	"time"
)

func main050101() {
	// 	创建定时器的三种方式
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	// 第一种方式 通过休眠的方式
	time.Sleep(2 * time.Second)
	fmt.Println(time.Now())

	/*
		第二种方式 创建定时器
			time.NewTimer(时间间隔n):在指定的时间n后
		系统自动将当前时间写入到time结构体中
	*/
	time1 := time.NewTimer(2 * time.Second) // 2秒以后将时间写入通道中
	fmt.Println(<-time1.C)                  // 读取通道中的时间

	/*
		第三种方式
			time.After(时间间隔)
		time.After() 等同于 time.NewTimer() 可以跟踪代码查阅
	*/
	time2 := time.After(2 * time.Second)
	fmt.Println(<-time2)

	// 没有出现死锁是因为 在结构中声明了通道的长度是1
}
