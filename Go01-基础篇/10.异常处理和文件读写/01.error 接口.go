package main

import (
	"errors"
	"fmt"
)

/*
	程序错误类型：
		语法错误：编辑时异常 比如：语法错误 缺少括号等
		逻辑错误：编译时错误 比如错误的算法导致运行结果不符
		运行错误：运行时异常 比如内存泄露 以零作整数等
*/
func test100101(a int, b int) (value int, err error) {
	//0不能作为除数
	if b == 0 {
		err = errors.New("0不能作整数")
	} else {
		value = a / b
	}
	return
}

func main100101() {
	value, err := test100101(10, 0)
	if err == nil {
		fmt.Println(value)
	} else {
		fmt.Println(err)
	}
}

//再写一遍
func main100102() {
	value, err := test100102(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}

func test100102(a int, b int) (value int, err error) {
	if b == 0 {
		err = errors.New("除数不可以为0")
	} else {
		value = a / b
	}
	return
}
