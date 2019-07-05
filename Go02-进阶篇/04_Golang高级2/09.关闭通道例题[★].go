package main

import (
	"fmt"
	"time"
)

func say1(str string, c chan int) {

	for i := 1; i <= 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(str, "次数：", i, "\t", i*100)
	}
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)
}

func say2(str string) {
	for i := 1; i <= 5; i++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Println(str, "次数：", i, "\t", i*150)
	}
}

func main_X() {
	c := make(chan int, 5)
	go say1("hello", c)
	say2("world")
	for v := range c {
		fmt.Println(v)
	}
}

/*
	我们引入一个信道，默认的，信道的存消息和取消息都是阻塞的，
	在 goroutine 中执行完成后给信道一个值 0，则主函数会一直等待信道中的值，
	一旦信道有值，主函数才会结束。
*/
//结果验证和预期不一样 思考解决分析
//https://www.runoob.com/go/go-concurrent.html
