package main

import (
	"fmt"
	"os"
)

const path = `E:\\Go_WorkSpace\\FullStack-go\\Go01-基础篇\\12_文件操作\\test.txt`

func main10020201() {
	//创建一个文件
	fp, err := os.Create("E:\\Go_WorkSpace\\FullStack-go\\Go01-基础篇\\12_文件操作\\test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	f, err := os.OpenFile("E:\\Go_WorkSpace\\FullStack-go\\Go01-基础篇\\12_文件操作\\test.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, _ = f.WriteString("传说中的暴龙兽！\t")
	n, _ := f.WriteString("传说中的暴龙兽!") //备注中文下的标点符号 也同样占用3个字节
	fmt.Println(n)
	fmt.Println("操作执行成功！")
}

func main10020202() {
	//创建文件
	fp, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	f1, err := os.OpenFile(path, os.O_RDWR, 6)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = f1.WriteString("传说中的暴龙兽！\r\t")
	_, err = f1.WriteString("传说中的基尔兽！\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("文件操作成功！")
}

/*
	1、WriteString()将字符串写入文本，返回写入的字节数
	2、Write()将字节写入文件，返回写入的字节数
	在UTF8编码（国际通用编码）下，一个中文占用3个字节
	在GBK（简体中文，GB2312也是简体中文）编码下，一个中文占2个字节
*/
func main10020203() {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	f, err = os.OpenFile(path, os.O_RDWR, 6)
	if err != nil {
		fmt.Println(err)
		return
	}
	var b []byte = []byte{1, 2, 3}   //字符型切片 占位符%c
	n, _ := f.Write(b)               //中文返回3个字节
	_, _ = f.Write([]byte("\t日历圆上")) //将字符串转换成字符型写入
	fmt.Println(n)
	fmt.Println("文件写入操作成功！")
}
