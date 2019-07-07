package main

import (
	"fmt"
	"time"
)

// 方法一：通过休眠的方式
func main1401() {
	// 延迟3S后打印当前时间
	time.Sleep(3 * time.Second)
	fmt.Println(time.Now())
}

// 方法二：创建定时器
func main1402() {
	Timer := time.NewTimer(2 * time.Second)
	Timer.Reset(time.Second)
	for i := 0; i < 5; i++ {
		go printer14(i)
	}
	fmt.Println(<-Timer.C)
}
func printer14(i int) {
	fmt.Println("输出：", i)
}

// 方法三：time.After(时间间隔)
func main1403() {
	Timer := time.After(2 * time.Second)
	fmt.Println(<-Timer)
}

func main1404() {
	time.Sleep(time.Second * 2)
	fmt.Println(time.Now())

	Timer := time.NewTimer(3 * time.Second)
	fmt.Println(<-Timer.C) // 通道中声明了 容量为1 因此不会死锁

	Timer2 := time.After(5 * time.Second)
	fmt.Println(<-Timer2)
}

func main1405() {
	Timer := time.NewTimer(10 * time.Second)
	Timer.Reset(time.Second)
	go func() {
		<-Timer.C
		fmt.Println("传说中的钢铁城！")
	}()
	Timer.Stop() // 关闭定时器
	for {
		;
	}
}

func main1406() {
	Ticker := time.NewTicker(2 * time.Second) // 间隔两秒
	ec := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			<-Ticker.C
			fmt.Println("传说中的暴龙兽！")
		}
		ec <- true
	}()
	<-ec
}

func main1407() {
	Ticker := time.NewTicker(3 * time.Second)
	EcTicker := make(chan bool)
	go func() {
		num := 0
		for {
			<-Ticker.C
			fmt.Println("钢铁加鲁鲁兽进化！")
			num++
			if num == 3 {
				EcTicker <- true
			}
		}
	}()
	<-EcTicker
}

// 多学一招：通过定时器也可以实现时钟周期的效果
func main1408() {
	num := 0
	for {
		Timer := time.NewTimer(3 * time.Second)
		<-Timer.C
		fmt.Println("传说中钢铁侠！")
		num++
		if num == 3 {
			break
		}
	}
}

func main1409() {
	// 练习：在11点35分打印"锄禾日当午"
	Ticker := time.NewTicker(5 * time.Second) // 每隔五秒 监听一次
	for {
		<-Ticker.C
		if time.Now().Hour() == 1 && time.Now().Minute() == 28 {
			fmt.Println("锄禾日当午！")
			break
		}
	}
}

func main14010() {
	Ticker := time.NewTicker(time.Second)
	for {
		nowTime := <-Ticker.C
		if nowTime.Hour() == 1 && nowTime.Minute() == 33 {
			fmt.Println("传说中的鸿雁~！")
			break
		}
	}
}
