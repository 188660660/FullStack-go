package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main050701() {
	//创建随机数种子
	rand.Seed(time.Now().UnixNano())
	//生成100-999的随机数
	random := make([]int, 3)
	random[0] = rand.Intn(9) + 1
	random[1] = rand.Intn(10)
	random[2] = rand.Intn(10)

	fmt.Println(random)
	usernum := make([]int, 3)

	var num int
	var flag int = 0
	for { //结束条件在里面
		for {
			fmt.Println("请输入一个三位数")
			fmt.Scan(&num)

			if num >= 100 && num <= 999 {
				break
			}
			fmt.Println("输入错误,请重新输入")
		}

		usernum[0] = num / 100
		usernum[1] = num / 10 % 10
		usernum[2] = num % 10

		for i := 0; i < 3; i++ {
			if usernum[i] > random[i] {
				fmt.Printf("您输入的第%d个数太大了\n", i+1)
			} else if usernum[i] < random[i] {
				fmt.Printf("您输入的第%d个数太小了\n", i+1)
			} else {
				fmt.Printf("恭喜您,第%d位数相同\n", i+1)
				flag++
			}
		}

		if flag == 3 {
			fmt.Println("成功")
			break
		} else {
			flag = 0
		}
	}
}
