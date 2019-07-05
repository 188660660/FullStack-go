package main

import (
	"fmt"
	"time"
)

func main050301_1() {
	// 每隔一段时间 执行一个事件
	// 例题 每隔一秒打印一次:"锄禾日当午" 一共打印三次
	myTaker := time.NewTicker(time.Second)
	i := 0
	quit := make(chan bool) // 退出程序开关
	go func() {
		for {
			<-myTaker.C
			fmt.Println("锄禾日当午")
			i++
			if i == 3 {
				quit <- true
			}
		}
	}()
	<-quit
}

// 例题 每隔一秒打印一次:"锄禾日当午" 一共打印三次
func main050201_2() {
	myTicker := time.NewTicker(time.Second)
	i := 0
	quit := make(chan bool)
	go func() {
		for {
			<-myTicker.C
			fmt.Println("锄禾日当午！")
			i++
			if i == 3 {
				quit <- true
			}
		}
	}()
	<-quit
}

func main050302() {
	// 通过定时器实现
	i := 0
	for {
		myTimer := time.NewTimer(time.Second)
		<-myTimer.C
		fmt.Println("锄禾日当午")
		i++
		if i == 3 {
			break
		}
	}
}

// 练习：在11点35分打印"锄禾日当午"
func main050303() {
	myTicker := time.NewTicker(time.Second)
	for {
		nowTime := <-myTicker.C
		if nowTime.Hour() == 4 && nowTime.Minute() == 19 {
			fmt.Println("锄禾日当午")
			break
		}
	}
}

// 练习：在11点35分打印"锄禾日当午"
func main050304() {
	myTicker := time.NewTicker(time.Second * 3)
	for {
		time2 := <-myTicker.C
		if time2.Hour() == 4 && time2.Minute() == 24 {
			fmt.Println("汗滴禾下土")
			break
		}
	}
}
