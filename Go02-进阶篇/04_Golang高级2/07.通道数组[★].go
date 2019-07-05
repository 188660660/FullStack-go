package main

import (
	"fmt"
	"math/rand"
	"time"
)

//设置值
func setValue10(c chan int) {
	c <- rand.Intn(100)
}

func main040701() {
	rand.Seed(time.Now().UnixNano())
	var num []chan int = make([]chan int, 0)
	for i := 1; i <= 5; i++ {
		c := make(chan int)
		num = append(num, c) //将通道添加到切片中

	}
	go func() {
		for _, v := range num {
			v <- rand.Intn(100)
		}
	}()
	//	取值
	for _, v := range num {
		fmt.Println(<-v)
	}
}

func main040702() {
	// 需求 往通道切片里面写入五个100内的随机数
	rand.Seed(time.Now().UnixNano()) // 创建随机数种子
	var num []chan int = make([]chan int, 0)
	for i := 0; i < 5; i++ {
		c := make(chan int)
		num = append(num, c) // 将通道追加到切片中去
	}
	go func() {
		for _, v := range num {
			v <- rand.Intn(100)
		}
	}()
	for _, v := range num {
		fmt.Println(<-v)
	}
}

func main040703() {
	// 需求 往通道切片里面写入五个100内的随机数
	rand.Seed(time.Now().UnixNano()) // 创建随机数种子
	var num = make([]chan int, 0)
	for i := 0; i < 5; i++ {
		c := make(chan int)
		num = append(num, c)
	}
	go func() {
		for _, v := range num {
			v <- rand.Intn(100)
		}
	}()
	for _, v := range num {
		fmt.Println(<-v)
	}
}

// 需求 往通道切片里面写入五个100内的随机数

func main040704() {
	rand.Seed(time.Now().UnixNano()) // 创建随机数种子
	var num = make([]chan int, 0)    // 创建切片通道 容量为0
	for i := 0; i < 5; i++ {
		c := make(chan int)  // 创建通道
		num = append(num, c) // 将创建好的通道追加到 切片通道中去
	}

	go func() {
		for _, v := range num {
			v <- rand.Intn(99) // 给切片中的每个通道都附加一个随机值
		}
	}()

	for _, v := range num {
		fmt.Println(<-v) // 打印输出每个切片通道中的值
	}
}

// 通道数组
func main040705() {
	rand.Seed(time.Now().UnixNano())
	var num [5]chan int = [5]chan int{}
	for i := 0; i < 5; i++ {
		c := make(chan int)
		num[i] = c
		go setValue1(c)
	}
	for _, v := range num {
		fmt.Println(<-v)
	}
}

func setValue1(c chan int) {
	c <- rand.Intn(99)
}