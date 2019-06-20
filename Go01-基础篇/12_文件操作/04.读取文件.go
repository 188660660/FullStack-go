package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const path1204 = `E:\\Go_WorkSpace\\FullStack-go\\Go01-基础篇\\12_文件操作\\test.txt`

/*
	第二种方法的核心主要是如下几个方法：
		buf := bufio.NewReader(fp)	//创建缓冲区
		f,err := buf.ReadBytes('\n')	//界定符结束 返回字符型数据 string()转换
		f,_,err := buf.ReadLine() //返回一行数据 string()转换
		f,err := buf.ReadString('\n') //界定符结束 返回一行数据 直接输出

	读取文件失败的原因：
		1.	文件不存在
		2.  读取权限不足
		3.  超过文件打开上限
*/
func main120401_1() {
	//方法一
	fp, err := os.Open(path1204) //文件以只读方式打开
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	contend := make([]byte, 500) //什么切片
	_, err = fp.Read(contend)    //将内容读取到切片中
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(contend)) //将切片转成字符串并输出
}

func mainmain120401_2() {
	//方法一 第二遍
	fp, err := os.Open(path1204) //以只读方式打开
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	content := make([]byte, 500)
	_, err = fp.Read(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}

func main120401_3() {
	//以只读方式打开文件
	fp, err := os.Open(path1204)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	content := make([]byte, 500)
	f, err := fp.Read(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(f)) //转换成string类型
}

/*
关于string的补充：
	1. string默认值是空字符串 ""，而不是nil，并且string类型是不能和nil作比较的：
       mismatched types string and nil。
	2.用索引号如 s[i]访问得到的是字节（切片同理），而不是字符。
	3.len函数得到的是字符串字节数而不是字符数。
	4.不能通过下标来修改字符串元素，同理不能通过下标来获取元素指针。
*/
func main120401_4() {
	fp, err := os.Open(path1204)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	content := make([]byte, 500)
	f, err := fp.Read(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f) //60字节
	fmt.Println(string(content))
}

/**************使用缓冲读取文件**************/

func main120402_1() {
	fp, err := os.Open(path1204)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	//使用缓冲方式读取文件
	buf := bufio.NewReader(fp)
	for {
		//1.从缓冲中读取到定界符，返回字符切片
		//readbytes,err := buf.ReadBytes('\n')
		//fmt.Print(string(readbytes))

		//2.从缓冲中读取一行
		readbytes, _, err := buf.ReadLine()
		fmt.Println(string(readbytes))

		//3.在缓冲中读取到定界符 返回字符串
		//str,err := buf.ReadString('\n')
		//fmt.Print(str)

		if err == io.EOF { //结束条件 读取的指针到末位就结束循环
			break
		}
	}
}

func main120402_2() {
	//读取文件练习 只读模式
	fp, err := os.Open(path1204)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	f := make([]byte, 500)
	_, err = fp.Read(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(f))
}

//只读方式
func main120402_3() {
	fp, err := os.Open(path1204)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	content := make([]byte, 500)
	_, err = fp.Read(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}

func main120402_4() {
	fp, err := os.Open(path1204)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	buf := bufio.NewReader(fp)
	f, err := buf.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(f))
}

func main120402_5() {
	fp, err := os.Open(path1204)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	buf := bufio.NewReader(fp)
	for {
		//f,err := buf.ReadBytes('\n')

		//f,_,err := buf.ReadLine()

		f, err := buf.ReadString('\n')

		fmt.Printf(f)
		if err == io.EOF {
			break
		}
	}
}
