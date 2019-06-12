package main

import "fmt"

func main030601() {
	/*
		!案例：百钱白鸡
			cock 公鸡 5钱
			hen 母鸡 3钱
			chicken 小鸡 1/3钱
	*/

	chicken := 0
	count := 0
	//100/5公鸡最大值为20 所以此处我们作为结束条件
	for coke := 0; coke <= 20; coke++ {
		for hen := 0; hen <= 33; hen++ {
			count++
			//小鸡数量= 100-公鸡-母鸡
			chicken = 100 - coke - hen
			if 100 == coke*5+hen*3+chicken/3 && chicken%3 == 0 {
				fmt.Printf("公鸡数量:%d 母鸡数量:%d 小鸡数量：%d\n", coke, hen, chicken)
			}
		}
	}
	fmt.Println("执行次数", count)
}
