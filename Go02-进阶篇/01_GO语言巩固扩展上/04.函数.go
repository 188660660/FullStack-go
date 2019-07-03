package main

import (
	"fmt"
	"strconv"
)

func add(num1, num2 int) int {
	return num1 + num2
}
func sub(num1, num2 int) int {
	return num1 - num2
}

type FuncType func(int, int) int

func cal(a, b int, callback FuncType) int {
	return callback(a, b)
}
func main0401() {
	//函数的一个参数是函数类型，这个函数参数就是回调函数。在OOP中实现多态
	rs1 := cal(100, 200, add)
	rs2 := cal(100, 200, sub)
	fmt.Println(rs1, rs2)
}

/*
	将字符串转为Int的例子，在转换失败的情况下执行回调函数，
	输出错误信息
*/
type Callback func(msg string)

//将字符串转换为int64，如果转换失败调用Callback
func stringToInt(str string, callback1 Callback) int64 {
	if value, err := strconv.ParseInt(str, 10, 64); err != nil {
		callback1(err.Error())
		return 0
	} else {
		return value
	}
}

func errLog(msg string) {
	fmt.Println("Convert error:", msg)
}
func main0402() {
	var stu [3]string = [3]string{"tom", "berry", "ketty"}
	//var slice[]string=stu[:]
	var slice []string = stu[:2]
	fmt.Println(slice)
	fmt.Println(strToInt("18", errLog1))
}

type calback2 func(msg string)

func strToInt(str string, callback calback2) int64 {
	if value, err := strconv.ParseInt(str, 10, 64); err != nil {
		callback(err.Error())
		return 0
	} else {
		return value
	}
}
func errLog1(msg string) {
	fmt.Println(msg)
}
