package main

import (
	"fmt"
	"os"
)

func main12010201() {
	_, err := os.Create("E:\\Go_WorkSpace\\FullStack-go\\Go01-基础篇\\12_文件操作\\test.text")
	if err != nil {
		fmt.Println(err)
		return
	}
	//漏写了关闭文件
	fmt.Println("文件创建成功！")
}

func main12010202() {
	fp, err := os.Create("E:\\Go_WorkSpace\\FullStack-go\\Go01-基础篇\\12_文件操作\\test.text")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	fmt.Println("文件创建成功！")
}
