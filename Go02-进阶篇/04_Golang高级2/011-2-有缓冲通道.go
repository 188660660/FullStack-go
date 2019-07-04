package main

import "fmt"

//1、有缓冲channel
func main0201() {
	var c = make(chan int, 3)
	fmt.Println("长度是：", len(c), "\t容量：", cap(c))
	//写入
	go func() {
		for i := 1; i <= 5; i++ {
			c <- i
			fmt.Println("写入:", i, "长度：", len(c), "容量：", cap(c))
		}
	}()
	//读取
	for i := 1; i <= 5; i++ {
		num := <-c
		fmt.Println("读取:", num, "长度：", len(c), "容量：", cap(c))
	}
}

//2、关闭channel
func main() {
	c := male
}
