package main

import (
	"fmt"
	"io"
	"os"
)

const path1203 = `E:\\Go_WorkSpace\\FullStack-go\\Go01-基础篇\\12_文件操作\\test.txt`

func main() {
	//在指定的位置插入数据
	fp, err := os.OpenFile(path1203, os.O_WRONLY, 6)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	//方法一
	_, _ = fp.WriteString("野火烧不尽") //每次执行都是从头往后写
	//方法二
	_, _ = fp.Write([]byte("一岁一枯荣"))
	//方法三
	n, _ := fp.Seek(6, io.SeekEnd)
	fmt.Println(n)
	_, _ = fp.WriteAt([]byte("美国"), n)
	fmt.Println("文件操作执行成功！")
}
