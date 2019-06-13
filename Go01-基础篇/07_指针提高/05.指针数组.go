package main

import "fmt"

func main060501() {
	//var arr [3]int //->int
	//指针数组
	var arr [3]*int //->*int

	a := 10
	b := 20
	c := 30

	arr[0] = &a
	arr[1] = &b
	arr[2] = &c

	fmt.Println(arr)       //[0xc000064080 0xc000064088 0xc000064090]
	fmt.Printf("%p\n", &a) //0xc000064080
	fmt.Printf("%p\n", &b) //0xc000064088
	fmt.Printf("%p\n", &c) //0xc000064090

	//通过指针数组改变变量的值
	*arr[0] = 250

	fmt.Println(a)

	fmt.Println(len(arr))

	//变量指针数组对应的内存空间的值
	for i := 0; i < len(arr); i++ {
		fmt.Println(*arr[i])
	}
}
