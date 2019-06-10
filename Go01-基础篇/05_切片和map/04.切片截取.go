package main

import "fmt"

func main050401() {
	//定义一个切片
	s := []int{1, 2, 3, 4, 5}

	//切片名[起始下标：] ★包含起始下标
	//slice := s[2:] //[3,4,5]

	//切片名[:结束位置] ★不包含结束下标
	//slice := s[:2] //[1,2]

	//切片名[起始位置:结束位置] ★不包含结束下标
	//slice := s[0:2] //[1,2]

	//切片名[起始位置：结束位置：容量]
	//大于等于len 小于等于cap len <= xxx <= cap
	slice := s[0:2:2]
	fmt.Println(slice)
	fmt.Println(cap(slice))
}

func main050402() {
	//定义一个切片
	s := []int{1, 2, 3, 4, 5}
	//下标 下标 下标
	slice1 := s[2:3] // [3] 相邻的只截取最后一个 不包含结束位置
	slice2 := s[2:4] //[3 4] 非相邻的包含结束位置 不包含结束位置

	s[1] = 123
	fmt.Println(s)
	fmt.Println(slice1)
	fmt.Println(slice2)

	//切片名 本身就是地址
	//一个字节在内存中占8bit(位)
	fmt.Printf("%p\n", s)     //0xc00008e030
	fmt.Printf("%p\n", &s[1]) //0xc00008e038
	fmt.Printf("%p\n", &s[2]) //0xc00008e040
	//fmt.Printf("%p\n",&s[3])
	fmt.Printf("%p\n", slice1) //0xc00008e040
	fmt.Printf("%p\n", slice2) //0xc00008e040
	//疑问 2 和slice2 地址一样？

	//切片名[:]获取切片中的所有元素 s[:] = s 这种写法
	slice3 := s[:]
	slice4 := s

	fmt.Println(s)
	fmt.Println(slice3)
	fmt.Println(slice4)

	news := make([]int, 5)
	copy(news, s)
	fmt.Printf("%p\n", news)
	fmt.Printf("%p\n", s)
}
