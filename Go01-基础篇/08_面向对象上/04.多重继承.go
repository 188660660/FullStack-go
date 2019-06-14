package main

import "fmt"

type test080401 struct {
	id   int
	name string
}

type test080402 struct {
	test080401
	sex string
	age int
}

//注意结构体不能嵌套本结构体
//结构体可以嵌套本结构体指针类型 链表
type test080403 struct {
	//*test080403 OK
	//test080403 //err
	test080402
	score int
}

func main080401() {
	var s = test080403{test080402{test080401{1, "最强王者"}, "男", 48}, 888}
	fmt.Println(s) //{{{1 最强王者} 男 48} 888}

	s.score = 998

	s.sex = "未知"
	s.age = 1600

	s.id = 6
	s.name = "尊者"

	fmt.Println(s) //{{{6 尊者} 未知 1600} 998}

	s.test080402.test080401.name = "武士"

	fmt.Println(s) //{{{6 武士} 未知 1600} 998}
}
