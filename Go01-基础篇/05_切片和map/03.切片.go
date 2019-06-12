package main

import "fmt"

func main050301() {
	/*
		!语法总结：
			数组的定义 var 数组名 [元素个数]数据类型
			切片的定义 var 切片名 []数据类型
	*/
	//常规声明切片类型
	var s1 []int
	fmt.Println(s1) //输出：[]

	//自动推导切片类型
	s2 := make([]int, 5)
	//给切片赋值
	s2[0] = 123
	s2[1] = 456
	s2[2] = 789
	s2[3] = 781
	s2[4] = 951
	//s2[5] = 887 //err 索引超出范围

	//通过append添加切片信息
	fmt.Println(len(s2)) //5

	fmt.Println(s2)
	s2 = append(s2, 555)
	fmt.Println(s2)

	//通过len查看切片长度
	fmt.Println(len(s2)) //6
}

func main050302() {
	s := make([]int, 5)

	s[0] = 123
	s[1] = 456
	s[2] = 789
	s[3] = 881
	//s[4] = 881 未赋值的默认值为0

	fmt.Println(s)

	//切片循环 - 推荐方法一
	for i, v := range s {
		fmt.Println(i, v)
	}

	//切片循环 - 常规方法二
	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}

	fmt.Println(s)
}

func main050303() {
	//不写元素个数的叫切片 必须写元素个数的叫数组
	var s []int = []int{1, 2, 3, 4, 5}

	s = append(s, 6, 7, 8, 9)

	//切片容量大于等于切片的长度
	fmt.Println(s)      //[1 2 3 4 5 6 7 8 9]
	fmt.Println(len(s)) //9
	fmt.Println(cap(s)) //10

	//容量每次扩展为上次的倍数
	s = append(s, 11, 2, 33, 56)
	fmt.Println("长度:", len(s)) //13
	fmt.Println("容量：", cap(s)) //20

	/*
		特性：
			1.如果整体数据没有超过1024字节 每次扩展为上次的倍数
			2.超过1024直接 每次扩展为上次的1/4
			(自己运行测试 容量在超出的时候才会翻倍)
	*/
	s = append(s, 6, 7, 8, 9, 10, 11, 15, 18, 66)
	fmt.Println("长度:", len(s)) //17
	fmt.Println("容量：", cap(s)) //20
}

func main050304() {
	//定义一个切片
	s1 := make([]int, 5)
	var s2 []int

	s1 = append(s1, 1, 2, 3, 4, 5)
	s2 = append(s2, 5, 4, 3, 2, 1)

	//获取切片的容量和长度股
	fmt.Println(len(s1)) //10
	fmt.Println(len(s2)) //5
	fmt.Println(cap(s1)) //10
	fmt.Println(cap(s2)) //6

	/*
		在切片打印时 只能打印有效长度的数据
		cap不能作为数组打印的条件
	*/

	for i := 0; i < len(s1); i++ {
		fmt.Println(s1[i])
	}

}
