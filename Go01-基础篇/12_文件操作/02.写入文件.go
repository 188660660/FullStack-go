package main

import (
	"fmt"
	"os"
)

func main100601() {
	//os.create创建文件时 如果文件不存在 会直接创建一个新文件 如果文件已存在 则会将其覆盖掉
	fp, err := os.Create("E:\\Go_WorkSpace\\FullStack-go\\Go01-基础篇\\10_异常处理和文件读写\\test.txt")
	if err != nil {
		fmt.Println("文件创建失败！")
		return
	}

	defer fp.Close()

	/*
		写入文件：
		\n 在文件中是换行 window文本文件换行是 \r\n
		在go语言中 一个汉字占三个字节
	*/
	n, _ := fp.WriteString("林志玲")
	fmt.Println(n)
	_, _ = fp.WriteString("结婚了！")

	fmt.Println("文件创建成功！")
}

func main100602() {
	//创建文件
	fp, err := os.Create("E:\\Go_WorkSpace\\FullStack-go\\Go01-基础篇\\10_异常处理和文件读写\\test2.txt")
	if err != nil {
		fmt.Println("文件创建失败！")
		return
	}
	//关闭文件
	defer fp.Close()

	//1.将字符切片写入到文件中
	a := []byte{'a', 'b', 'e'}
	//如果是对文件对象.write(字符切片)
	n, _ := fp.Write(a)
	fmt.Println(n)

	//2.将字符串转成字符切片写入到文件中
	str := "锄禾日当午"
	b := []byte(str)
	_, _ = fp.Write(b)
}

func main100603() {
	/*
		打开文件
			os.open(文件名) 只读方式打开
			os.open(文件路径)
			openfile不能创建新文件
			os.openfile(文件名,打开模式,打开权限)
	*/
	//os.Create("E:\\Go_WorkSpace\\FullStack-go\\Go01-基础篇\\10_异常处理和文件读写\\test3.txt")
	fp, err := os.OpenFile("E:\\Go_WorkSpace\\FullStack-go\\Go01-基础篇\\10_异常处理和文件读\\test3.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("文件创建失败！")
		return
	}
	defer fp.Close() //延迟释放内存

	fmt.Println("文件创建成功！")

	//通过字符串获取字符切片
	//b := []byte("你瞅啥")
	//当使用writeat进行指定位置插入数据时会依次覆盖源内容
	//fp.WriteAt(b,0)
}
