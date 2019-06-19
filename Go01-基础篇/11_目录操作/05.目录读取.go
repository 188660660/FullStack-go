package main

import (
	"fmt"
	"os"
)

func main110501() {
	//打开目录路径 反引号是字符串常量
	path := `E:\Go_WorkSpace\FullStack-go\Go01-基础篇\11_目录操作\test1001`
	fp, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir) //O_RDONLY 表示只读模式 ModeDir表示打开目录
	if err != nil {
		fmt.Println(err)
		return
	}

	defer fp.Close() //关闭文件夹
	/*
		读取目录：
		n表示读取文件或文件夹的数目
		<=0 表示全部读取
		返回文件切片数据
	*/

	fileinfo, _ := fp.Readdir(0)
	for _, v := range fileinfo {
		if v.IsDir() {
			fmt.Println("文件夹名：", v.Name(), "文件夹大小:：", v.Size())
		} else {
			fmt.Println("文件名：", v.Name(), "文件大小:", v.Size())
		}
	}
}

func main100502() {
	path := `E:\Go_WorkSpace\FullStack-go\Go01-基础篇\11_目录操作\test1001`
	fp, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err := fp.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, k := range f {
		if k.IsDir() {
			fmt.Println("文件夹名：", k.Name(), "文件夹大小：", k.Size())
		} else {
			fmt.Println("文件名：", k.Name(), "文件大小：", k.Size())
		}
	}
}

func main100503() {
	//反引号定义常量 不用转义
	path := `E:\Go_WorkSpace\FullStack-go\Go01-基础篇\11_目录操作`
	fp, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	f, err := fp.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("--%v--\n", f) //[0xc00009a000 0xc00009a070 0xc00009a0e0 0xc00009a150 0xc00009a1c0 0xc00009a230 0xc00009a2a0]
	for _, fp := range f {
		if fp.IsDir() {
			fmt.Println("文件夹名：", fp.Name(), "文件夹大小：", fp.Size())
		} else {
			fmt.Println("\t文件名：", fp.Name(), "文件大小：", fp.Size()) //\t表示缩进
		}
	}
}
