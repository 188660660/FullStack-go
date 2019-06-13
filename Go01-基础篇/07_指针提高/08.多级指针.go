package main

import "fmt"

//定义全局多级指针
var pp **int

func main070801() {
	a := 10
	p := &a
	pp := &p  //二级指针 二级指针存储一级指针的内存地址
	p3 := &pp //三级指针 三级指针存储二级指针的内存地址

	fmt.Println(p)  //0xc00000a0c8
	fmt.Println(&p) //0xc000006028
	fmt.Println(p3) //0xc000006030

	fmt.Println("-------------------")

	//一级指针的值
	fmt.Println(*pp) //0xc000064080
	fmt.Println(p)   //0xc000064080
	fmt.Println(&a)  //0xc000064080

	//变量a 的值
	fmt.Println(*p3)  //0xc000006028
	fmt.Println(**pp) //10
	fmt.Println(*p)   //10
	fmt.Println(a)    //0xc000006028

	fmt.Printf("%T\n", a)  //int
	fmt.Printf("%T\n", p)  //*int
	fmt.Printf("%T\n", pp) //**int
	fmt.Printf("%T\n", p3) //***int
}
