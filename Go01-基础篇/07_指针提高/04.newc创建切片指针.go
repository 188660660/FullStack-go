package main

import "fmt"

func main070401() {
	//创建切片指针
	var slice *[]int
	fmt.Printf("%p\n", slice) //0x0空指针
	slice = new([]int)
	fmt.Printf("%p\n", slice) //0xc00005a440分配了内存空间

	*slice = append(*slice, 1, 2, 3, 4)

	fmt.Println(len(*slice)) //len = 4

	//循环输出 方法一：
	for i := 0; i < len(*slice); i++ {
		fmt.Println((*slice)[i])
	}

	//循环输出 方法二：
	for k, v := range *slice {
		fmt.Printf("下标：%d 值：%v\n", k, v)
	}
}
