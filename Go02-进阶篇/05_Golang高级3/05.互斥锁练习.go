package main

import (
	"fmt"
	"sync"
	"time"
)

// 互斥锁 相互竞争谁抢到锁 谁使用
var mutex0501 sync.Mutex

func main050501() {
	go getstr051("Hello")
	go getstr052("World")
	for {
		;
	}
}

func getstr051(str string) {
	prt05(str)
}

func getstr052(str string) {
	prt05(str)
}

func prt05(str string) {
	mutex0501.Lock()
	for _, v := range str {
		fmt.Printf("%c", v)
		time.Sleep(time.Millisecond * 300)
	}
	mutex0501.Unlock()
}

func main050502() {
	num := 0
	var lock sync.Mutex
	for i := 1; i < 1000; i++ {
		go func(idx int) {
			lock.Lock()
			defer lock.Unlock()
			num += 1
			fmt.Println(num)
		}(1)

	}
	time.Sleep(time.Second)
}

func main050503() {
	for i := 1; i <= 1000; i++ {
		var mutex sync.Mutex
		num := 0
		go func(idx int) {
			mutex.Lock()
			defer mutex.Unlock()
			num += i
			fmt.Println(num)
		}(1)
	}
	time.Sleep(time.Second)
}

func main050504() {
	ch := make(chan struct{}, 2)
	var l sync.Mutex
	go func() {
		l.Lock()
		defer l.Unlock()
		fmt.Println("gotroutinel:我会锁定大概2S")
		time.Sleep(time.Second * 2)
		fmt.Println("goroutine1：我解锁了你们去抢吧")
		ch <- struct{}{}
	}()
	go func() {
		fmt.Println("go2:等待解锁")
		l.Lock()
		defer l.Unlock()
		fmt.Println("go2太好了 我也解锁了")
		ch <- struct{}{}
	}()
	// 等待GO程 执行结束
	for i := 0; i < 2; i++ {
		<-ch
	}
}

// 互斥锁 只能同时被一个GO程锁定 因此互斥锁大多用于同步的使用
