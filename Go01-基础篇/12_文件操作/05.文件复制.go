package main

import (
	"fmt"
	"io"
	"os"
)

/*
	文件复制原理：
		打开一个文件A 再创建一个文件B 将A的内容拷贝到B文件上
补充知识：
	文本是字符流，字符流可以读取到某个字符结束,比如换行
	图片是二级制字节流，只能全部读出来以后一起执行
*/
const path1206 = "E:\\Go_WorkSpace\\FullStack-go\\Go01-基础篇\\12_文件操作\\"

func main120601() {
	fp1, err1 := os.Open(path1206 + "20714428.jpg")
	fp2, err2 := os.Create(path1206 + "20714429.jpg")
	defer fp1.Close()
	defer fp2.Close()
	if err1 != nil || err2 != nil { //！=nil (=) 有问题 只要一个满足就进入到这里
		fmt.Println("文件打开或创建失败！")
		return
	}
	//读取文件
	buf := make([]byte, 100)
	for {
		n, err := fp1.Read(buf)
		if err == io.EOF {
			break
		}
		_, _ = fp2.Write(buf[:n])
	}
}

//练习
func main120602() {
	fp1, err1 := os.Open(path1206 + "20714428.jpg")
	fp2, err2 := os.Create(path1206 + "207144299.jpg")
	defer func() {
		fp1.Close()
		fp2.Close()
	}()
	if err1 != nil || err2 != nil {
		fmt.Println("图片读取或创建失败")
		return
	}
	//读取目录
	buf := make([]byte, 100)
	for {
		n, err := fp1.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		_, _ = fp2.Write(buf[:n])
		if err == io.EOF {
			break
		}
	}
}

func main120603() {
	fp1, err1 := os.Open(path1206 + "20714428.jpg")
	fp2, err2 := os.Create(path1206 + "111.jpg")

	defer func() {
		fp1.Close()
		fp2.Close()
	}()

	if err1 != nil || err2 != nil {
		fmt.Println("图片读取或创建失败！")
		return
	}
	buf := make([]byte, 100)
	for {
		_, err := fp1.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		_, _ = fp2.Write(buf)
		if err == io.EOF {
			break
		}
	}
}

func main120604() {
	fp1, err1 := os.Open(path1206 + "20714428.jpg")
	fp2, err2 := os.Create(path1206 + "999.jpg")
	defer fp1.Close()
	defer fp2.Close()
	if err1 != nil || err2 != nil {
		fmt.Println("图片读取或创建失败！")
		return
	}
	buf := make([]byte, 100)
	for {
		fp, err := fp1.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		_, _ = fp2.Write(buf[:fp])
		if err == io.EOF {
			break
		}
	}
}
