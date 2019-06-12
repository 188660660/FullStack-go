package main

import "fmt"

func main030501() {
	//break 跳出循环 在嵌套循环中 跳出本层循环
	for i := 0; i <= 10; i++ {
		if i == 2 {
			break //终止循环
		}
		fmt.Println(i)
	}

	for i := 0; i <= 10; i++ {
		if i == 2 {
			continue //跳出当前循环 程序继续
		}
		fmt.Println(i)
	}
}
