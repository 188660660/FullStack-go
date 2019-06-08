package main

import "fmt"

//！len()既可以计算字符串的长度 也可以计算不定参函数的长度
func plus(add ...int) {
	sum := 0
	//方法一：
	//for i:=0;i<len(add);i++{
	//	sum += add[i]
	//}
	//方法二：
	for _, value := range add {
		sum += value
	}
	/*
		!总结：
		for 和range 都已遍历集合中的数据信息 数组 切片 结构体数组和map
	*/
	fmt.Println(sum)
}
func main0801() {
	test3(25, 67)
	test3(89, 71)
	//plus(1, 2, 3)
	//plus(1, 2, 3, 4, 5)
	//plus(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}

func test3(args ...int) {
	fmt.Println(args)
	fmt.Println(len(args))
	//fmt.Printf("数据类型：%T\n",args)

	for i := 0; i < len(args); i++ {
		fmt.Printf("下标：%d 数值： %d \n", i, args[i])
	}
}
