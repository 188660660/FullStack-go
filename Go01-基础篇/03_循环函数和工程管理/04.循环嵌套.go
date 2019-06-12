package main

import "fmt"

func main030401() {
	//嵌套循环中  执行次数为外层*内层?
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Println(i, j)
		}
	}
}

func main030402() {
	/*
		案例：打印九九乘法口诀
			1*1=1
			1*2=2 2*2=4
			1*3=3 2*3=6 3*3=9
	*/
	//外层控制行 内层控制列
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, i*j)
		}
		fmt.Println()
	}
}
