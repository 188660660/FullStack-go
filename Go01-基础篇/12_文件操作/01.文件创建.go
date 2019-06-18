package _2_文件操作

import (
	"fmt"
	"os"
)

func main100501() {
	fp, err := os.Create("E:\\Go_WorkSpace\\FullStack-go\\Go01-基础篇\\10_异常处理和文件读写\\test.txt")
	if err != nil {
		fmt.Println("文件创建失败！")
		return //如果return出现在主函数中 表示程序的结束
	}

	/*
		延迟调用 关闭文件
		如果不关闭:
			1.会占用内存和缓冲区
			2.文件打开有上限 65535
	*/

	defer fp.Close()
	fmt.Println("文件创建成功！")
}
