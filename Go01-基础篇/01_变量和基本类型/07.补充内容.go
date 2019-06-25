package main

import (
	"fmt"
	"os"
)

//main函数不支持任何返回值
func main0107() {
	//获取命令行参数
	if len(os.Args) > 1 {
		fmt.Println(os.Args)
	}
	fmt.Println("Golang ")
	//退出返回值
	os.Exit(1)
}
