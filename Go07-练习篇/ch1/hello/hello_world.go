package main

import (
	"fmt"
	"os"
)

func mainch1() {
	//通过os.Args获取命令行参数
	if len(os.Args)-1 > 0 {
		fmt.Println(os.Args)
	}
	fmt.Println("Hello World！")
	//运行结束时返回
	os.Exit(-1)
}
