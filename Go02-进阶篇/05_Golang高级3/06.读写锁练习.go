package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var num06 int = 0
var rwmutex sync.RWMutex

func main050601() {
	/*
			互斥锁在大量读操作的时候 就显得效率比较底下
		读过读操作较多 就建议使用读写锁
		读共享 写独占 写优先
	*/
	/*
		var Rwmutex
	*/
	// 五个程序同时写
	// 创建随机数种子
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 5; i++ {
		go write06(i)
	}

	// 五个程序同时写
	for i := 1; i <= 5; i++ {
		go read06(i)
	}
	for {
		;
	}
}

func write06(i int) {
	for {
		num06 = rand.Intn(100)
		rwmutex.Lock()
		fmt.Println("第", i, "个程序 写了", num06)
		rwmutex.Unlock()
		time.Sleep(time.Millisecond * 300)
	}
}

func read06(i int) {
	for {
		rwmutex.RLock()
		fmt.Println("\t第", i, "个程序 读了", num06, "")
		rwmutex.RUnlock()
		time.Sleep(time.Millisecond * 300)
	}
}

/*
	写操作的锁定和解锁
	* func (*RWMutex) Lock
	* func (*RWMutex) Unlock
	读操作的锁定和解锁
	* func (*RWMutex) Rlock
	* func (*RWMutex) RUnlock
	读写锁总结
		1.同时只能有一个goroutine能够获得写锁定
		2.同时可以有任意多个goroutine获得读锁定
		3.同时只能存在写锁定或者读锁定 读写互斥
*/
// demo1
var count int

func main050602() {
	ch := make(chan struct{})
	for i := 0; i < 5; i++ {
		go write0602(i, ch)
	}

	for i := 0; i < 5; i++ {
		go read0602(i, ch)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
}

func read0602(n int, ch chan struct{}) {
	rwmutex.RLock()
	fmt.Printf("goroutine %d 进入读操作...\n", n)
	v := count
	fmt.Printf("goroutine %d 读取结束，值为：%d\n", n, v)
	rwmutex.RUnlock()
	ch <- struct{}{}
}

func write0602(n int, ch chan struct{}) {
	rwmutex.Lock()
	fmt.Printf("goroutine %d 进入写操作...\n", n)
	v := rand.Intn(1000)
	count = v
	fmt.Printf("goroutine %d 写入结束，新值为：%d\n", n, v)
	rwmutex.Unlock()
	ch <- struct{}{}
}

var m *sync.RWMutex

func main050603() {
	m = new(sync.RWMutex)
	// 多个同时读
	go read1(1)
	go read1(2)
	time.Sleep(time.Second * 2)
}

func read1(i int) {
	fmt.Println(i, "读取")
	m.RLock()
	fmt.Println(i, "读取中")
	time.Sleep(time.Millisecond * 300)
	m.RUnlock()
	fmt.Println(i, "读取结束")
}

// 050603的例子证明了  同读是可以的 不受影响
var m1 *sync.RWMutex

func main050604() {
	m1 = new(sync.RWMutex)
	// 写的时候啥都不能干
	go write0604(1)
	go read0604(2)
	go write0604(3)
	for {
		;
	}
}

func write0604(i int) {
	println(i, "write start")
	m1.Lock()
	println(i, "writing")
	time.Sleep(1 * time.Second)
	m1.Unlock()
	println(i, "write over")
}

func read0604(i int) {
	println(i, "read start")
	m1.RLock()
	println(i, "reading")
	time.Sleep(1 * time.Second)
	m1.RUnlock()
	println(i, "read over")
}

/*
练习注意： 这个案例是在网上找的一个demo  不知道大家看明白了没有 我跟随着练习 出来的结果感觉不能很好地证明
可能是理解有偏差
	由于读写互斥，上面这个示例中，写开始的时候，
	读必须要等写进行完才能继续。
	不然他只能继续等待，这就像只有一个茅坑，
	别我蹲着，你着急也不能去抢（为什么？有门关着呢！）。
1 write start
1 writing
2 read start
3 write start
1 write over
2 reading
2 read over
3 writing
3 write over
*/
