package main

import (
	"fmt"
	"os"
)

/*
	语法：
		Remove():删除叶子节点目录(目录为空)
		RemoveAll():删除目录和目录下所有内容
*/
func main100201() {
	err := os.Remove("E:\\Go_WorkSpace\\FullStack-go\\Go01-基础篇\\11_目录操作\\test")
	if err != nil {
		fmt.Println(err)
		return
	}
	//如果文件不为空 会报一个非空错误
	fmt.Println("目录删除成功！")
}

func main100202() {
	err := os.RemoveAll("E:\\Go_WorkSpace\\FullStack-go\\Go01-基础篇\\11_目录操作\\111")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("目录删除成功！")
}
