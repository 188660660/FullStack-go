package main

import "fmt"

//定义全局结构体
type Student struct {
	id    int
	name  string
	sex   string
	age   int
	score int
	addr  string
}

func main060401() {
	//var 结构体变量名 结构体类型

	/*
		定义结构体数组
		var 结构体数组名 [元素个数]结构体类型
	*/
	var arr [5]Student

	//len(数组名) 计算数组元素个数

	fmt.Println(arr)
	fmt.Println(len(arr))

	for i := 0; i < len(arr); i++ {
		_, _ = fmt.Scan(
			&arr[i].id,
			&arr[i].name,
			&arr[i].sex,
			&arr[i].age,
			&arr[i].score,
			&arr[i].addr,
		)
	}

	//结构体 排序  根据结构体成员进行排序
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; i++ {
			if arr[j].score > arr[j+1].score {
				//结构体数组中的元素 允许相互赋值 将结构体成员中改的所有数据相互交换
				//arr[j].score,arr[j+1].score  = arr[j+1].score,arr[j].score
				//结构体成员依次交换  不建议采用
				//arr[j].id,arr[j+1].id=arr[j+1].id,arr[j].id
				//arr[j].name,arr[j+1].name=arr[j+1].name,arr[j].name
				//arr[j].sex,arr[j+1].sex=arr[j+1].sex,arr[j].sex
				//arr[j].score,arr[j+1].score=arr[j+1].score,arr[j].score
				//arr[j].age,arr[j+1].age=arr[j+1].age,arr[j].age
				//arr[j].addr,arr[j+1].addr=arr[j+1].addr,arr[j].addr
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func main060402() {
	//[元素个数]数组 []切片len
	arr := []Student{ //写中括号的值就是结构体数组 不谢的话就是数组
		{101, "曹操", "男", 58, 90, "许昌"},
		{102, "夏侯惇", "男", 40, 100, "荆州"},
		{103, "张辽", "男", 38, 99, "许昌"},
	}
	//结构体切片 可以添加 使用append
	arr = append(arr, Student{104, "许褚", "男", 58, 90, "许昌"})
	arr = append(arr, Student{105, "典韦", "男", 40, 100, "荆州"})
	for k, v := range arr {

		fmt.Println(k, v)
	}
}
