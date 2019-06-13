package main

import "fmt"

func test070301(slice *[]int) {
	//但切片指针改变了原先的值 会释放掉原先的内存空间
	*slice = append(*slice, 123)
}
func main070301() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("%p\n", s)

	//切片指针作为函数参数是地址传递 形参可以改变实参的值
	test070301(&s)
	fmt.Printf("%p\n", s)
	fmt.Println(s)
}
