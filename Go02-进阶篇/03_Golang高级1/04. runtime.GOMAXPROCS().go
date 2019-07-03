package main

import (
	"fmt"
	"runtime"
)

func main030401() {
	n := runtime.GOMAXPROCS(100) //设置CPU核数，返回先前的设置
	fmt.Println(n)
	//fmt.Println(runtime.NumCPU())
	for {
		go fmt.Println(0)
		fmt.Println(1)
	}
	//文件操作和IO密集情况下可以去设置cpu核数，至少配置到硬件线程数目的5倍以上, 最大1024。
}

func main030402() {
	fmt.Println(runtime.NumCPU())  //获取计算机CPU硬件核数
	fmt.Println(runtime.Version()) //获取当前GO语言的版本
	fmt.Println(runtime.GOROOT())  //获取GO的安装目录
	runtime.GC()                   //立即执行一次垃圾回收
}
