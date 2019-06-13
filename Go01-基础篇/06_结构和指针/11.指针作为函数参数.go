package main

import "fmt"

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

//函数参数为指针类型
func main061101() {
	a := 10
	b := 20

	//值未作交换 值传递
	//swap(a, b) //10,20
	//指针作为函数参数传递是地址传递
	swap(&a, &b)
	fmt.Println(a) //20
	fmt.Println(b) //10
}
