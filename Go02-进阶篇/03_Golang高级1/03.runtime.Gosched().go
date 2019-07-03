package main

import (
	"fmt"
	"os"
	"runtime"
)

func testGo03(str string) {
	for i := 0; i < 2; i++ {
		/*
				runtime.Gosched()用于让出CPU时间片。
				这就像跑接力赛，A跑了一会碰到代码runtime.Gosched()
			    就把接力棒交给B了，A歇着了，B继续跑。
		*/
		fmt.Println(str)
	}
}
func main030301() {
	go testGo03("Hello")
	testGo03("World")
	/*
		World   World
		Hello   Hello
		Hello   World
		World   Hello
	*/
}

func testGo05(i int) {
	fmt.Println(i)
}
func testGo04(str string) {
	for i := 0; i < 4; i++ {
		/*
			没加runtime.Gosched()之前 没有机会输出i
		*/
		runtime.Gosched() //可以看到goroutine获得了运行机会，打印出了数字。
		go testGo05(i)
	}
}
func main030302() {
	testGo04("goroutine")
	//time.Sleep(1*time.Second)
	fmt.Println("main function")
}

func main030303() {
	go func() {
		for i := 0; i <= 5; i++ {
			fmt.Println(i)
		}
	}()
	for i := 0; i < 2; i++ {
		//让出时间片 想让别的协议执行 它执行完 再来执行主程
		runtime.Gosched()
		fmt.Println("hello")
	}
}
func test() {
	defer fmt.Println("c")
	//return
	runtime.Goexit() //终止所在的协程
	fmt.Println("d")
}
func main030304() {
	//创建新的协程
	go func() {
		fmt.Println("a")
		test() //调用其他函数
		fmt.Println("b")
	}()
	for {
		//死循环不然主程结束
	}
}
func testGo06() {
	defer fmt.Println("a")
	runtime.Goexit()
	fmt.Println("f")
}
func main030305() {
	go func() {
		defer fmt.Println("c")
		go testGo06()
		fmt.Println("d") //d c a 	a d c
	}()
	for {

	}
}

func test05() {
	defer fmt.Println("a")
	os.Exit(0) //退出进程，销毁空间，原来加载的空间全部销毁
	fmt.Println("b")
}
func main() {
	go func() {
		defer fmt.Println("c")
		test05()
		fmt.Println("d")
	}()
	for {

	}
}
