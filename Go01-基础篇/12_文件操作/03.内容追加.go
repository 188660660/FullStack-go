package main

import (
	"fmt"
	"io"
	"os"
)

const path1203 = `E:\\Go_WorkSpace\\FullStack-go\\Go01-基础篇\\12_文件操作\\test.txt`

func main120301() {
	//在指定的位置插入数据
	fp, err := os.OpenFile(path1203, os.O_WRONLY, 6)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	//方法一
	_, _ = fp.WriteString("abcdefg野火烧不尽") //每次执行都是从头往后写
	//方法二
	_, _ = fp.Write([]byte("一岁一枯荣"))
	_, _ = fp.WriteAt([]byte("追龙"), 6) //从开始往后 6个字节的偏移量追加
	n, _ := fp.Seek(6, io.SeekEnd)
	_, _ = fp.WriteAt([]byte("追龙"), n)
	//方法三
	//n, _ := fp.Seek(6, io.SeekEnd)
	//fmt.Println(n)
	//_, _ = fp.WriteAt([]byte("美国"), n)
	fmt.Println("文件操作执行成功！")
}

func main120302() {
	//_, err := os.Create(path1203)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	fp, err := os.OpenFile(path1203, os.O_APPEND, 6)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	_, _ = fp.WriteString("追加内容！")
	fmt.Println("追加成功！")
}

func main120303() {
	fp, err := os.OpenFile(path1203, os.O_RDWR, 6)
	if err != nil {
		fmt.Println(err)
		return
	}
	//_, err = fp.WriteString("上山打老虎")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	_, err = fp.WriteAt([]byte("★"), 6)
	_, err = fp.Seek(6, io.SeekEnd)
	fmt.Println("文件追加成功！")
}
