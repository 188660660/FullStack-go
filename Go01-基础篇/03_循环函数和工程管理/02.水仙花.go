package main

import "fmt"

func main030202() {
	/*
		水仙花案例：
			水仙花 一个三位数 100-999
			各个位数的立方和等于这个数本身
	*/
	for i := 100; i <= 999; i++ {
		//核心:num = a*a*a+b*b*b+c*c*c
		//第一步获取 a b c 各个位数的值
		a := i / 100     //取百位数
		b := i / 10 % 10 //取十位数
		c := i % 10      //取个位数
		if i == a*a*a+b*b*b+c*c*c {
			fmt.Printf("%d为水仙花数", i)
		}
	}

	/*
		a := i/100
		b := i/10%10
		c := 1%10
	*/
}
