package main

import "fmt"

func main070601() {
	//指针切片
	var slice []*int
	a := 10
	b := 20
	c := 30
	d := 40
	slice = append(slice, &a, &b, &c)

	slice = append(slice, &d)

	fmt.Println(*slice[0]) //10
	fmt.Println(slice)     //[0xc000064080 0xc000064088 0xc000064090 0xc000064098]

	//取数去
	fmt.Println(len(slice))

	//循环取值 第一种方式
	for i := 0; i < len(slice); i++ {
		fmt.Println(*slice[i])
	}
	//循环取值 第二种方式
	for k, v := range slice {
		fmt.Println(k, *v)
	}
}
