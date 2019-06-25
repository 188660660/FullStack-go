package main

import "fmt"

func main070101() {

	var arr = [5]int{1, 2, 3, 4, 5}

	fmt.Printf("%p\n", &arr)    //0xc00000c300
	fmt.Printf("%p\n", &arr[0]) //0xc00000c300

	//p := &arr
	//定义指针指向数组
	var p *[5]int
	//将指针变量和数组建立关系
	p = &arr

	//可以通过指针间接操作数组
	fmt.Println(*p)      //[1 2 3 4 5]
	fmt.Println((*p)[0]) //1

	(*p)[0] = 123
	fmt.Println(arr) //[123 2 3 4 5]

	//直接使用指针[下标] 操作数组元素
	fmt.Println(p[0])     //123
	fmt.Printf("%p\n", p) //0xc00008e030
	fmt.Printf("%T\n", p) //*[5]int
}

func main070102() {
	var arr = [5]int{1, 2, 3, 4, 5}
	//指向数组的指针
	//p := &arr //*[5]int

	p := &arr[0]
	fmt.Printf("%T", p) //*int

	p2 := &arr
	//len(指针变量) 输出元素个数
	fmt.Println(len(p2)) //*[5]int5

	//循环输出打印
	for i := 0; i < len(p2); i++ {
		fmt.Println(p2[i])
	}

	fmt.Printf("%p\n", p2)      //0xc00000c300
	fmt.Printf("%p\n", &arr)    //0xc00008e030
	fmt.Printf("%p\n", &arr[0]) //0xc00008e030
}
