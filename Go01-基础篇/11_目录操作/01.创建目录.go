package main

import (
	"fmt"
	"os"
)

/*
	1、Mkdir(地址，权限)：只能创建一级目录
	2、MkdirAll(地址，权限)：创建多级目录
特别注意：
	Golang的相对路径是相对于执行命令时的目录
*/
/*
获取运行文件目录
	_, filename, _, _ := runtime.Caller(1)
	fmt.Println(path.Join(path.Dir(filename)))
*/
func main110101() {
	//err := os.Mkdir("./test", 777) 相对路径
	err := os.Mkdir("E:\\Go_WorkSpace\\FullStack-go\\Go01-基础篇\\11_目录操作\\test1001", 777) //绝对路径
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("目录创建成功！")
}

func GetCurrentPat() {

}

func main100102() {
	err := os.MkdirAll("./test1002/test1002", 777)
	//err := os.MkdirAll("E:\\Go_WorkSpace\\FullStack-go\\Go01-基础篇\\11_目录操作\\test1001\\test1002",777)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("多级目录创建成功")
}
