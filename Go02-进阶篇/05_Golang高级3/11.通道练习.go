package main

import (
	"fmt"
	"time"
)

func main1101() {
	// 两个用户依次打印"hello world"
	c := make(chan bool)
	ec := make(chan bool)
	go user1(c)
	go user2(c, ec)
	<-ec
}
func user1(c chan bool) {
	printer("Hello ")
	c <- true

}

func user2(c, ec chan bool) {
	<-c
	printer("World")
	ec <- true
}

func printer(str string) {
	for _, v := range str {
		fmt.Printf("%c", v)
		time.Sleep(300 * time.Millisecond)
	}
}
