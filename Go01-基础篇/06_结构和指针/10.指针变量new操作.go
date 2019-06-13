package main

import "fmt"

func main061001() {
	var p *int
	fmt.Println(p) //nil

	//为指针变量创建一块内存空间 - 堆区
	p = new(int)
	//打印P存储空间的值
	fmt.Println(p) //0xc00000a0f0 创建好的空间值为数据类型的默认值
	//打印P指向空间的值
	fmt.Println(*p)

	var a int
	fmt.Println(&a) //0xc00000a0f8 8为8个bit 字节 结论验证成立
}
