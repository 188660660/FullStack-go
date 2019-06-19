package main

import (
	"fmt"
	"os"
	"strings"
)

func main110601() {
	//递归遍历文件夹
	path := `E:\Go_WorkSpace\FullStack-go\Go01-基础篇\11_目录操作`
	GetSubFile01(path, 0)
}

func GetSubFile01(p string, n int) {
	fp, err := os.OpenFile(p, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	f, err := fp.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, f := range f {
		fmt.Print(strings.Repeat("\t", n))
		if f.IsDir() {
			fmt.Println(f.Name())
			path := p + "\\" + f.Name()
			GetSubFile01(path, n+1)
		} else {
			fmt.Println(f.Name())
		}
	}
}

func main100602() {
	path := `E:\AP-REC\App`
	GetSubFile02(path, 0)
}

//递归函数
func GetSubFile02(p string, n int) {
	//打开文件
	fp, err := os.OpenFile(p, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()        //关闭文件
	f, err := fp.Readdir(1) //获取文件夹目录层级
	for _, v := range f {
		fmt.Printf(strings.Repeat("\t", n))
		if v.IsDir() {
			fmt.Println(v.Name())
			GetSubFile02(p+"\\"+v.Name(), n+1)
		} else {
			fmt.Println(v.Name())
		}
	}
}
