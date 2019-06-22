package main

import (
	"fmt"
	"os"
)

/*
	go build：编译源码，生成可执行文件
	go run：不生成可执行文件，直接运行源码
	go version：查看版本
*/

func main150401() {
	list := os.Args //获取命令行所有的参数
	for k, v := range list {
		//fmt.Println(k,v)
		fmt.Printf("第%d个参数是%s\n", k, v)
	}
}
