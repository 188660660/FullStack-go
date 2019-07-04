package main

import "fmt"

//1、如下代码输出什么
/*
func main0101() {
	str:=[]string{"a","b","c"}
	for _,v:=range str {
		go func() {
			fmt.Println(v)
		}()
	}
	time.Sleep(time.Second*2)
}
*/

//2、通过channel实现数据同步
/*
var c=make(chan bool)		//声明通道，用来控制用户执行的先后顺序
var flag=make(chan int)	//用来控制主线程不退出
func printer(str string){
	for _,ch:=range str{
		fmt.Printf("%c",ch)
		time.Sleep(time.Millisecond*300)
	}
}
func user1(){
	printer("hello ")
	c<-true
}
func user2(){
	<-c					//读阻塞，等到另一个go程往c中写入值
	printer("world")
	flag<-10			//写阻塞
}
func main() {
	go user1()
	go user2()
	<-flag				//一致阻塞，一直到user2执行完毕
}
*/

//3、通道赋值和取值
func main0103() {
	var c = make(chan int)     //保存数据
	var flag = make(chan bool) //控制数据同步
	//写入
	go func() {
		for i := 1; i <= 5; i++ {
			c <- i
			fmt.Println("写入：", i)
			flag <- true //阻塞，写入完毕阻塞，一直等到读取完毕再解除阻塞
		}
	}()
	//读取
	for i := 1; i <= 5; i++ {
		num := <-c
		fmt.Println("读取：", num)
		<-flag //解除阻塞
	}
}
