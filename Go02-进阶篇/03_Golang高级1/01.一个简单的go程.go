package main

import (
	"fmt"
	"time"
)

//创建一个并发的新goroutine

func HelloWord() {
	fmt.Println("Hello world goroutine") //go程
}
func main() {
	go HelloWord() //开启一个新的并发运行！
	time.Sleep(3 * time.Second)
	fmt.Println("主程！")
	//当main函数返回时 所有的goroutine都是暴力终结的
}

func testGo() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
}

func main02() {
	go testGo()
	time.Sleep(1 * time.Second)
	fmt.Println("主线程！") //主go程结束，即使子go程没有结束也退出
}
