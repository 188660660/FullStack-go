package main

import "fmt"

func demo(m map[int]string) {
	//map在函数中添加删除数据 会影响主调函数中的实参
	m[102] = "猪八戒"
	m[103] = "沙和尚"

	fmt.Println(len(m)) //3

	delete(m, 101)
	fmt.Println(m) //map[102:猪八戒 103:沙和尚]
}
func main051101() {
	m := make(map[int]string, 1)
	m[101] = "孙悟空"

	fmt.Println(len(m)) //1
	fmt.Println(m)      //map[101:孙悟空]
	//fmt.Println(cap(m)) //err cap()只能计算切片的容量
	//map作为函数参数是地址传递(引用传递)
	demo(m)

	fmt.Println(m) //map[102:猪八戒 103:沙和尚]
}

//m := make(map[int][]int) map value切片类型
//m := make(map[int][]struct) map value结构体类型
