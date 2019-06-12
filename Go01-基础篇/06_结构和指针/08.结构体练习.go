package main

import "fmt"

//定义结构体存储5名学生 三门成绩 求出每名学生的总成绩和平均成绩
type Student060801 struct {
	id    int
	name  string
	score []int //成绩存储为切片类型
}

func main060801() {
	//用结构体定义学生的信息
	stuinfo := []Student060801{
		{1, "小明", []int{85, 62, 78}},
		{2, "小王", []int{85, 63, 78}},
		{3, "小李", []int{85, 64, 78}},
		{4, "小张", []int{85, 65, 78}},
		{5, "小红", []int{85, 66, 78}},
	}

	//五名学生
	for i := 0; i < len(stuinfo); i++ {
		sum := 0
		var pvg float64
		//三门成绩
		for j := 0; j < len(stuinfo[i].score); j++ {
			sum += stuinfo[i].score[j]
		}
		pvg = float64(sum) / 3
		fmt.Printf("%s的总成绩为：%d 平均成绩为：%.2f\n", stuinfo[i].name, sum, pvg)
	}
}
