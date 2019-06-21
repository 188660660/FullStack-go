package main

import (
	"fmt"
	"io"
	"os"
)

func main120601() {
	/*
		文件拷贝，将已有的文件复制一份，同时重新命名。
		基本的思路：
		(1)让用户输入要拷贝的文件的名称(源文件)以及目的文件的名称
		(2)创建目的文件
		(3)打开源文件，并且读取该文件中的内容
		(4)将从源文件中读取的内容写到目的文件中。
	*/
	//定义变量
	var srcFileName, dstFilename string

	//获取用户输入信息
	fmt.Println("请输入源文件名称地址")
	_, _ = fmt.Scan(&srcFileName)
	fmt.Println("请输目标文件名称地址")
	_, _ = fmt.Scan(&dstFilename)

	if srcFileName == dstFilename {
		fmt.Println("源文件和目标文件地址名称不能相同！")
		return
	}

	//只读方式打开源文件
	sfp, err1 := os.Open(srcFileName)   //只读方式 读取源文件
	dfp, err2 := os.Create(dstFilename) //创建目标文件

	if err1 != nil || err2 != nil {
		fmt.Println("源文件地址读取或新文件创建失败！")
		return
	}

	//操作完成关闭文件
	defer dfp.Close()
	defer sfp.Close()

	//核心处理 从源文件读取内容往目标文件进行复制
	buf := make([]byte, 500*1024) //500K大小临时缓冲区

	for {
		n, err := sfp.Read(buf) //将源文件读取的内容 放到缓冲区
		if err != nil {
			fmt.Println(err)
			return
		}
		if err == io.EOF { //文件读取完毕 作为结束条件
			break
		}
		_, err = dfp.Write(buf[:n]) //往目标文件写入 读多少写入多少
		if err == nil {
			fmt.Println("恭喜您，文件拷贝操作成功！")
		} else {
			fmt.Println("拷贝错误：", err)
		}
		return
	}
}

/*
	请输入源文件名称地址
	C:\Users\VULCAN\Pictures\1.jpg
	请输目标文件名称地址
	E:\Go_WorkSpace\FullStack-go\Go01-基础篇\12_文件操作\2.jpg
*/
