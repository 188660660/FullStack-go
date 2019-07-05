package main

import (
	"fmt"
	"time"
)

func main050201() {
	/*
		Reset():重置定时器
		Stop(): 停止定时器
	*/
	// 哈哈
	mtime := time.NewTimer(2 * time.Second)
	mtime.Reset(time.Second)
	go func() {
		<-mtime.C
		fmt.Println("锄禾日当午！")
	}()
	// mtime.Stop() // 关闭定时器
	for {

	}
}

func main050202() {
	mtimer := time.NewTimer(2 * time.Second)
	mtimer.Reset(time.Second * 15)
	go func() {
		<-mtimer.C // 读取重新设置的时间
		fmt.Println("风雨彩虹~！")
	}()
	// mtimer.Stop() //终止定时器任务
	for {
		;
	}
}
