package main

import "fmt"

func main040401() {
	var c = make(chan int, 3)
	fmt.Println("通道长度为：", len(c), "\t通道容量为：", cap(c))
	//var ec = make(chan int,5)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
			fmt.Println("写入：", i, "长度:", len(c), "\t容量：", cap(c))
			//ec<-i
		}
	}()
	for i := 0; i < 5; i++ {
		num := <-c
		fmt.Println("读取：", num, "长度:", len(c), "\t容量：", cap(c))
		//<-ec
	}
	//异步写入数据 显示结果和计算器硬件以及CPU有关 写入 要快一些 读取在缓冲中读取 会比较忙
}

/*
注意：默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须又接收端相应的接收数据。
以下实例通过两个 goroutine 来计算数字之和，在 goroutine 完成计算后，它会计算两个结果的和：
*/
func sum040401(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}
func main040402() {
	s := []int{3, 5, 2, 6, 9, 8, 7}
	c := make(chan int)
	go sum040401(s[:len(s)/2], c)
	go sum040401(s[len(s)/2:], c)
	x, y := <-c, <-c // 从通道 c 中接收
	fmt.Println(x, y, x+y)
}

func main040403() {
	//这里我们定义了一个可以存储整数类型的带缓冲通道
	ch := make(chan int, 2)
	// 因为 ch 是带缓冲的通道，我们可以同时发送两个数据
	// 而不用立刻需要去同步读取数据
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range c {
		fmt.Println(i)
	}
}
