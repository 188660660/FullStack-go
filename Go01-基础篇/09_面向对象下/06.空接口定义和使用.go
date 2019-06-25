package main

import "fmt"

/*
	空接口的定义和特性：
		接口中不包含任何方法,所有的类型都实现了空接口
		空接口可以存储任意类型的数据
*/
func main090601() {
	var Tnil1 interface{} //空接口 可以保存任意类型数据

	Tnil1 = 10                      //整形
	Tnil1 = "abc"                   //字符串型
	Tnil1 = 3.14                    //浮点型
	Tnil1 = [3]string{"A", "B"}     //数组
	Tnil1 = []string{"A", "B", "C"} //切片

	fmt.Println(Tnil1)
}

func main090602() {
	//fmt.Println("hello")
	//fmt.Println(10)
	//fmt.Println(3.14)
	//fmt.Println(test)

	//空接口 可以接受任意类型数据
	var i interface{}
	i = 10
	fmt.Println(i)         //10
	fmt.Printf("%p\n", &i) //0xc0000501c0

	i = "hello world"
	fmt.Println(i)         //hello world
	fmt.Printf("%p\n", &i) //0xc0000441d0

	var arr = [3]int{1, 2, 3}
	i = arr
	fmt.Println(i)         //[1 2 3]
	fmt.Printf("%p\n", &i) //0xc0000501c0
}

func main090603() {
	//空接口切片
	var i []interface{}
	i = append(i, 1, 2, "hello", "你瞅啥", [3]int{1, 2, 3})

	for j := 0; j < len(i); j++ {
		fmt.Println(i[j])
	}

	for k, v := range i {
		fmt.Println(k, v)
	}
}
