package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main1201() {
	c := make(chan int)
	ec := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
			fmt.Println("写入：", i)
			ec <- true
		}
	}()
	for i := 0; i < 5; i++ {
		fmt.Println("读取：", <-c)
		time.Sleep(time.Millisecond * 3)
		<-ec
	}
}

func main1202() {
	// 有缓冲通道
	c := make(chan int, 1)
	// ec := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
			fmt.Println("写入：", i)
			// ec<-true
		}
	}()
	for i := 0; i < 5; i++ {
		fmt.Println("读取：", <-c)
		// <-ec
	}
}

func main1203() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		for i := 7; i < 8; i++ {
			c <- i
		}
	}()
	for v := range c {
		fmt.Println(v)
	}
}

// 单双向通道作为参数
/*
	make(chan int) //双通道
	make(<-chan int) //写通道
	make(chan<- int)
*/
func main1204() {
	c := make(chan int)
	ec := make(chan bool)
	go wchannel(c)
	go rchannel(c, ec)
	// time.Sleep(5*time.Second)
	<-ec
}

func wchannel(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
		fmt.Println("写入：", i)
	}
	close(c)
}

func rchannel(c <-chan int, ec chan bool) {
	for v := range c {
		fmt.Println("读取：", v)
		ec <- true
	}
}

func main1205() {
	// 通道中的数组
	rand.Seed(time.Now().UnixNano()) // 创建随机数种子
	c := make([]chan int, 5)
	for i := range c {
		v := make(chan int)
		c[i] = v
		go setValue(v)
	}
	for _, v := range c {
		fmt.Println(<-v)
	}
}

func setValue(c chan int) {
	c <- rand.Intn(100)
}

// 给通道数组创建五个随机数
func main1206() {
	// 创建随机数种子
	rand.Seed(time.Now().UnixNano())
	// 创建通道数组
	c := [5]chan int{}
	for i := range c {
		c[i] = make(chan int)
		go setValue2(c[i])
	}
	for _, v := range c {
		fmt.Println(<-v)
	}
}

func setValue2(c chan int) {
	c <- rand.Intn(100)
}

func main1207() {
	// 通道切片
	rand.Seed(time.Now().UnixNano()) // 创建随机数种子
	var c []chan int
	for i := 0; i < 5; i++ {
		c = append(c, make(chan int))
	}
	go func() {
		for _, v := range c {
			v <- rand.Intn(100)
		}
	}()
	for _, v := range c {
		fmt.Println(<-v)
	}
}

func setValue3(c chan int) {
	c <- rand.Intn(100)
}

func main1208() {
	// 通道数组
	var c [5]chan int
	ec := make(chan bool, 5)
	for i := range c {
		c[i] = make(chan int)
		go setValue4(c[i], ec)
	}

	for _, v := range c {
		fmt.Println(<-v)
		<-ec
	}
}

func setValue4(c chan int, ec chan bool) {
	ec <- true
	c <- rand.Intn(100)
}

func main1209() {
	// 通道数组
	c := [5]chan int{}
	for i := 0; i < 5; i++ {
		c[i] = make(chan int)
		go setValue5(c[i])
	}
	for _, v := range c {
		fmt.Println(<-v)
	}
}
func setValue5(c chan int) {
	c <- rand.Intn(100)
}

// 通道切片
func main1210() {
	c := make([]chan int, 0)
	for i := 0; i < 5; i++ {
		ci := make(chan int)
		c = append(c, ci)
	}
	go func() {
		for _, v := range c {
			v <- rand.Intn(100)
		}
	}()
	for _, v := range c {
		fmt.Println(<-v)
	}
}

func main1211() {
	c := make(chan bool)
	for i := 0; i < 9; i++ {
		go setValue6(c, i)
	}
	<-c
}

func setValue6(c chan bool, i int) {
	fmt.Println(i)
	if i == 3 {
		c <- true
	}
}

func main1212() {
	c := make(chan bool)
	for i := 0; i < 9; i++ {
		go setValue7(c, i)
	}
	for i := 0; i < 9; i++ {
		<-c
	}
}

func setValue7(c chan bool, i int) {
	fmt.Println(i)
	c <- true
}

func main1213() {
	wg := sync.WaitGroup{}
	wg.Add(9)
	for i := 0; i < 10; i++ {
		go setValue8(&wg, i)

	}
	wg.Wait()
}

func setValue8(wg *sync.WaitGroup, i int) {
	fmt.Println(i)
	wg.Done()
}

func main1214() {
	wg := sync.WaitGroup{}
	wg.Add(9)
	for i := 0; i < 9; i++ {
		go setValue9(&wg, i)
	}
	wg.Wait()
}

func setValue9(wg *sync.WaitGroup, i int) {
	fmt.Println(i)
	wg.Done()
}
