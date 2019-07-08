package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main1601() {
	// 例题：如果监听到数据通道有值就打印，如果监听到退出开关有值就退出
	c := make(chan int)
	ec := make(chan bool)
	go ct1(c, ec)
	ct2(c, ec)

}

func ct1(c chan<- int, ec chan bool) {
	for i := 1; i <= 5; i++ {
		c <- i
		time.Sleep(time.Second)
	}
	ec <- true
	close(c)
}

func ct2(c chan int, ec chan bool) {
	for {
		select {
		case num := <-c:
			fmt.Println("传说中的：", num)
		case <-ec:
			return
		}
	}
}

func main1602() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}

// select会循环检测条件 如果有满足则执行退出 否则一直循环检测
func main1603() {
	c := make(chan int)
	ec := make(chan bool)
	go t1(c, ec)
	select {
	case <-c:
		for v := range c {
			fmt.Println(v)
		}
	case <-ec:
		fmt.Println("执行退出")
		goto end
	}
end:
	fmt.Println("输出结束")
}

func t1(c chan int, ec chan bool) {
	for i := 0; i < 5; i++ {
		c <- i
		time.Sleep(time.Second)
	}
	close(c)
	ec <- true
}

func main1604() {
	c := make(chan int)
	ec := make(chan bool)
	go printer16(c, ec)
	for {
		select {
		case ch := <-c:
			// for v := range c {
			fmt.Println(ch)
			// }
		case <-ec:
			goto end
		}
	}
end:
}

func printer16(c chan int, ec chan bool) {
	for i := 0; i < 10; i++ {
		c <- i
		time.Sleep(time.Second)
	}
	ec <- true
}

func main1505() {
	c := make(chan int)
	ec := make(chan bool)
	go printer17(c, ec)
	for {
		select {
		case ct := <-c:
			fmt.Println("执行次数：", "输出值：", ct)
		case <-ec:
			goto end
		}
	}
end:
}

func printer17(c chan int, ec chan bool) {
	for i := 0; i < 9; i++ {
		c <- i
		time.Sleep(time.Second)
	}
	ec <- true
}

/*
	例题1：将斐波那契数列写入通道并读取
	描述：在通道中输入数列，并读取。
*/
func main1605() {
	c := make(chan int)
	ec := make(chan bool)
	go printer118(c, ec)
	num1, num2, i := 1, 1, 1 // 第一个数列 第二数列 个数
	for {
		c <- num1
		num1, num2 = num2, num2+num1
		i++
		if i == 10 {
			break
		}
	}
}
func printer118(c <-chan int, ec chan bool) {
	// 打印机
	for {
		select {
		case num := <-c:
			fmt.Print(num, " ")
		case <-ec:
			/*
				Goexit终止调用它的go程。其它go程不会受影响。
				Goexit会在终止该go程前执行所有defer的函数。
				在程序的main go程调用本函数，会终结该go程，
				而不会让main返回。因为main函数没有返回，程序会继续执行其它的go程。
				如果所有其它go程都退出了，程序就会崩溃。
			*/
			runtime.Goexit()
		}
	}
}

func main1606() {
	c := make(chan int)
	ec := make(chan bool)
	go pr16(c, ec)
	num1, num2, i := 1, 1, 1
	for {
		num1, num2 = num2, num1+num2
		select {
		case <-c:
			i++
			fmt.Print(" ", <-c)
		case <-ec:
			runtime.Goexit()
		}
		if i == 10 {

		}
	}
}
func pr16(c chan int, ec chan bool) {
	for i := 1; i <= 10; i++ {
		c <- i
	}
	<-ec
}

func main1607() {
	c := make(chan int)
	ec := make(chan bool)
	go lc(c, ec)
	num1, num2, i := 1, 1, 1
	for {
		c <- num1
		num1, num2 = num2, num1+num2
		i++
		if i == 10 {
			break
		}
	}
}
func lc(c <-chan int, ec chan bool) {
	for {
		select {
		case num := <-c:
			fmt.Println(" ", num)
		case <-ec:
			runtime.Goexit()
		}
	}
}

func main1608() {
	c := make(chan int)   // 创建通道
	ec := make(chan bool) // 创建开关
	go tc(c, ec)          // 监听通道和开关

	num1, num2, i := 1, 1, 1 // 初始化
	for {
		c <- num1
		num1, num2 = num2, num1+num2
		i++
		if i == 10 {
			break
		}
	}
	ec <- true
}

func tc(c <-chan int, ec chan bool) {
	for {
		select {
		case c := <-c:
			fmt.Println(" ", c)
		case <-ec:
			runtime.Goexit() // 监听到退出按钮 则退出当前GO程
		}
	}
}

func main1609() {
	// 例题2：从多个通道中读取数据
	c1, c2 := make(chan int), make(chan string)
	go func() {
		for {
			select {
			case num1, ok := <-c1:
				if ok {
					fmt.Println(num1)
				}
			case num2, ok := <-c2:
				if ok {
					fmt.Println(num2)
				}
			}
		}
	}()
	c1 <- 17
	c2 <- "你好啊"
	close(c1)
	close(c2)
}

func main1610() {
	// 多个通道
	c1, c2 := make(chan int), make(chan string)
	go func() {
		for {
			select {
			case v, ok := <-c1:
				if ok {
					fmt.Println(v)
					close(c1)
				}
			case v, ok := <-c2:
				if ok {
					fmt.Println(v)
					close(c2)
				}
			}
		}
	}()
	c1 <- 88
	c2 <- "传说中的加鲁鲁"
}

func main1611() {
	// 场景：有时候会出现go程阻塞的情况，可以利用select来设置超时，避免整个程序进入阻塞
	c := make(chan int)
	ec := make(chan bool)

	go func() {
		for {
			select {
			case num := <-c:
				fmt.Println(num)
			case <-time.After(time.Second * 2):
				fmt.Println("超时退出")
				ec <- true
			}
		}
	}()
	c <- 10
	<-ec
}

// 空的select可以阻塞 main 函数
func main1612() {
	go func() {
		i := 0
		for {
			i++
			fmt.Println("锄禾日当午", i)
		}
	}()
	select {}
}

func main1613() {
	c := make(chan int)
	ec := make(chan bool)
	// c<-10
	go func() {
		for {
			select {
			case num, ok := <-c:
				if ok {
					fmt.Println(num)
				}
			case <-time.After(time.Second * 2):
				fmt.Println("操作超时")
				ec <- true
			}
		}
	}()
	c <- 10
	<-ec
}

// 死锁 产生的原因 ： 竞争资源 通信不畅
// 第一种情况 一个通道 一个GO程读写
func main1614() {
	c := make(chan int)
	// c<-10

	// 解决方法 两个GO程
	go func() {
		c <- 10
	}()

	fmt.Println(<-c)
}

// 第二种情况 GO程开启之前就使用了通道
func main1615() {
	c := make(chan int)
	// c<-10
	go func() {
		c <- 10
		<-c
		fmt.Println(<-c)
	}()
}

func main1616() {
	// 通道1中调用了 通道2 通道2中调用通道1
	c1, c2 := make(chan int), make(chan int)
	go func() {
		for {
			select {
			case c1 <- 1:
				c2 <- 10
			}
		}
	}()
	for {
		select {
		case c2 <- 10:
			<-c1
		}
	}
}

// 问题： 如何解决死锁 回答：加锁
var metex sync.Mutex // 互斥锁
// 练习 互斥锁
func main1617() {

	go user1601("Hello")
	go user1602("World")
	for {
		;
	}
}

func user1601(str string) {
	printer1605(str)
}

func user1602(str string) {
	printer1605(str)
}

func printer1605(str string) {
	metex.Lock()         // 加锁
	defer metex.Unlock() // 解锁
	for _, v := range str {
		fmt.Printf("%c", v)
		time.Sleep(time.Second)
	}
}

/*
	互斥锁在大量读操作的时候 效率比较底下
	如果读操作比较多久建议使用 读写锁
特点：读共享 写独占 写优先
*/
var mutex1 sync.Mutex

func main1618() {
	go pr17("Hello ")
	go pr18("World")
	for {
		;
	}
}

func pr17(str string) {
	pr19(str)
}

func pr18(str string) {
	pr19(str)
}

func pr19(str string) {
	mutex1.Lock()
	defer mutex1.Unlock()
	for _, v := range str {
		fmt.Printf("%c", v)
		time.Sleep(time.Millisecond * 150)
	}
}

var num int = 0
var rem sync.RWMutex

func main1619() {
	rand.Seed(time.Now().UnixNano()) // 创建随机数种子
	for i := 1; i <= 5; i++ {
		go write(i)
	}
	for i := 1; i <= 5; i++ {
		go read(i)
	}
	for {
		;
	}
}

func write(i int) {
	for {
		rem.Lock()
		num = rand.Intn(100)
		fmt.Printf("第%d个go程写%d\n", i, num)

		rem.Unlock()
		time.Sleep(time.Millisecond * 300)
	}
}

func read(i int) {
	for {
		rem.RLock()
		fmt.Printf("\t第%d个go程读取%d\n", i, num)
		rem.RUnlock()
		time.Sleep(time.Millisecond * 300)
	}
}
