package main

import "fmt"

/*
	Golang提供了指针用于操作数据内存，并通过引用来修改变量。
	只声明未赋值的变量，golang都会自动为其初始化为零值，基础数据类型的零值比较简单，引用类型和指针的零值都为nil，nil类型不能直接赋值，因此需要通过new开辟一个内存，或者通过make初始化数据类型，或者两者配合，然后才能赋值。
	指针也是一种类型，不同于一般类型，指针的值是地址，这个地址指向其他的内存，通过指针可以读取其所指向的地址所存储的值。
	函数方法的接受者，也可以是指针变量。无论普通接受者还是指针接受者都会被拷贝传入方法中，不同在于拷贝的指针，其指向的地方都一样，只是其自身的地址不一样。

*/
func main060901() {
	var a = 123

	//定义整形指针变量 指向a的地址
	/*
		指针类型定义
			var 指针 *数据类型 一级指针
	*/
	var p *int
	//将a的地址 赋值给指针变量p
	p = &a

	//通过指针变量 可以间接访问变量对应的内存空间
	*p = 345

	fmt.Println(a)  //345
	fmt.Println(*p) //345
	fmt.Println(p)  //0xc00000a0c8
	fmt.Println(&a) //0xc00000a0c8
}

func main0606902() {
	a := 10
	p := &a

	*p = 123
	fmt.Println(a)
	fmt.Printf("%T", p) //*int

	/*
		声明了一个指针 默认为nil(空指针 值为0)
		指向了内存地址编号为0的空间
	*/

	var x *int
	fmt.Println(x) //*int<nil>

	//p=0xc042058080//野指针  指针变量指向了一个未知的空间
	//访问野指针和空指针对应的内存空间都会报错

	fmt.Println(x) //*int<nil>
}
