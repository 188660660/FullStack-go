package main

import "fmt"

type test090501 struct {
	id   int
	name string
}

type test090502 struct {
	age int
	sex string
}

type test090503 struct {
	test090501
	test090502
	score int
}

func main080501() {
	var stu test090503 = test090503{test090501{1, "王者"}, test090502{18, "男"}, 888}

	stu.id = 117
	stu.name = "宗师"

	stu.age = 998
	stu.sex = "未知"

	stu.score = 100

	fmt.Println(stu) //{{117 宗师} {998 未知} 100}
}
