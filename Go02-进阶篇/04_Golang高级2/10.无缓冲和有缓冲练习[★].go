package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Printf("sum:")
	fmt.Printf("%#v\n", sum)
	c <- sum // 把 sum 发送到通道 c
	fmt.Println("after channel pro")
}

func main100101() {
	/*
		无缓冲是同步的，例如 make(chan int)，就是一个送信人去你家门口送信，你不在家他不走，你一定要接下信，他才会走，无缓冲保证信能到你手上。
		有缓冲是异步的，例如 make(chan int, 1)，就是一个送信人去你家仍到你家的信箱，转身就走，除非你的信箱满了，他必须等信箱空下来，有缓冲的保证信能进你家的邮箱。
		修改一下上面笔记中的程序如下：
	*/

	// 通道不带缓冲，表示是同步的，只能向通道 c 发送一个数据，只要这个数据没被接收然后所有的发送就被阻塞
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	fmt.Println("go [0,3]")
	go sum(s[:len(s)/2], c) //a

	//这里开启一个新的运行期线程，这个是需要时间的，本程序继续往下走

	fmt.Println("go [3,6]")
	go sum(s[len(s)/2:], c) //b
	fmt.Println("go2 [0,3]")
	go sum(s[:len(s)/2], c) //c
	fmt.Println("go2 [3,6]")
	go sum(s[len(s)/2:], c) //d

	/*
	   a b c d和main一起争夺cpu的，他们的执行顺序完全无序，甚至里面不同的语句都相互穿插
	   但无缓冲的等待是同步的，所以接下来a b c d都会执行，一直执行到c <- sum后，开始同步阻塞
	   因此after channel pro是打印不出来的, 等要打印after channel pro的时候，main就结束了
	*/

	fmt.Println("go3 start waiting...")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("go3 waited 1000 ms")

	//因为a b c d都在管道门口等着，这里度一个，a b c d就继续一个，这个结果的顺序可能是acbd
	aa := <-c
	bb := <-c
	fmt.Println(aa)
	fmt.Println(bb)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

// 无缓冲和有缓冲
// 无缓冲 make(chan int)是同步的
// 有缓冲 make(chan int ,3)是异步的

func sum10(num []int, c chan int) {
	sum := 0
	for _, v := range num {
		sum += v
	}
	c <- sum
	fmt.Println("after channel pro")
}
func main() {
	// 通道不带缓冲 表示是同步的 只能向通道c发送一个数据 只要这个数据没有被接受 那么其他的数据都将会被阻塞
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	fmt.Println("go [0,3]")
	go sum10(s[:len(s)/2], c)
	fmt.Println(<-c)

	// 这里开启一个新的运行期线程 这个是需要时间的 本程序继续往下走
	fmt.Println("go [3,6]")
	go sum(s[len(s)/2:], c)
	fmt.Println(<-c)

	fmt.Println("go2 [0,3]")
	go sum(s[:len(s)/2], c) // c
	fmt.Println(<-c)

	fmt.Println("go2 [3,6]")
	go sum(s[len(s)/2:], c) // d
	fmt.Println(<-c)

	/*
	   a b c d和main一起争夺cpu的，他们的执行顺序完全无序，甚至里面不同的语句都相互穿插
	   但无缓冲的等待是同步的，所以接下来a b c d都会执行，一直执行到c <- sum后，开始同步阻塞
	   因此after channel pro是打印不出来的, 等要打印after channel pro的时候，main就结束了
	*/
}