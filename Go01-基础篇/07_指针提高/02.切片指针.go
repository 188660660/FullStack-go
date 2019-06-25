package main

import "fmt"

func main070201() {
	var slice = []int{1, 2, 3, 4, 5}

	//p := &slice //*[]int

	var p *[]int //二级指针
	p = &slice

	fmt.Printf("%T\n", p) //*[]int

	//切片名本身就是一个地址
	fmt.Printf("%p\n", slice) //0xc00008e030
	fmt.Printf("%p\n", p)     //0xc00005a420
	fmt.Printf("%p\n", *p)    //0xc00008e030
}

func main070202() {
	var slice = []int{1, 2, 3, 4, 5}

	fmt.Printf("%p\n", slice) //0xc00008e030

	//第一种写法
	//p := &slice //*[]int
	fmt.Printf("%p\n", slice) //0xc00008e030
	var p *[]int
	p = &slice

	*p = append(*p, 6, 7, 8, 9, 10)
	fmt.Printf("%p\n", slice) //0xc00006a050
	fmt.Printf("%p\n", p)     //0xc00005a420
	fmt.Printf("%p\n", *p)    //0xc00006a050
	fmt.Println(slice)        //[1 2 3 4 5 6 7 8 9 10]
}
