package main

import (
	"fmt"
	"time"
)

func test1() {
	for i := 1; i < 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Println(i)
	}
}

func test2() {
	fmt.Println("Hello World!")
}

func main030201() {
	go test1() //创建第一个goroutine
	go test2() //创建第二个goroutine
	time.Sleep(1 * time.Second)
	fmt.Println("main function")
}

/*
思考：子程一里面有sleep 会导致子程二 堵塞或等待吗？
	解析:不会
		当程序执行go FUNC()的时候，只是简单的调用然后就立即返回了，
		并不关心函数里头发生的故事情节，所以不同的goroutine直接不影响，
		main会继续按顺序执行语句。
*/
