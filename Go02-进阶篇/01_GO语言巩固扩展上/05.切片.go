package main

import (
	"fmt"
	"sort"
)

func main0501() {
	stu1 := []string{"tom", "berry", "ketty"}
	stu2 := []string{"李白", "杜甫", "白居易", "小白"}
	stu1 = append(stu1, stu2...) //切片中追加切片
	fmt.Println(stu1)

	var num []int = []int{1, 2, 3, 4, 5, 6}
	slice := num[3:] //切片截取，456
	//slice=num[:2]			//12
	//slice=num[1:6]		//不包含6，所有正确 23456
	//slice:=num[:]			//获取切片中的所有内容
	////        起始下标  结束下标  最大下标
	//slice:=num[2:3:6]		//切片的容量，结束下标<=最大下标<=原切片的大小 //3
	//[low:high:max] //获取下标从low开始的元素。len=hign-low，cap=max-low
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice)) //1 4

	str := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	s1 := str[2:5] //cde
	fmt.Println(s1)
	s2 := s1[1:5] // defg
	fmt.Println(s2)
}

func main0502() {
	str := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	s1 := str[2:5:6] //c","d","e" f g
	fmt.Println(s1)
	s2 := s1[1:5] //"d","e","f","g","h"
	fmt.Println(s2)
}
func main0503() {
	var num1 []int = []int{1, 2, 3, 4, 5}
	var num2 []int = []int{10, 20, 30, 40}
	copy(num1[2:4], num2[1:3])
	fmt.Println(num1)
}

func main0504() {
	a, b, c := 10, 20, 30
	var slice []*int
	slice = append(slice, &a, &b, &c)
	*slice[1] = 200
	fmt.Println(b)
}

func main0505() {
	var slice []int = []int{10, 20, 30, 40, 50}
	var p *[]int
	p = &slice                //将数组的地址付给指针
	fmt.Printf("%p\n", p)     //0x120300e0
	fmt.Printf("%p\n", slice) //0x1202d800  因为一个在堆区，一个在栈区，所以地址不一样
	fmt.Printf("%p\n", *p)    //0x11ecd800	*p=slice
	*p = append(*p, 1, 2, 3)  //通过指针更改切片的值
	fmt.Println(*p)
}

func main0506() {
	var p *[]int
	fmt.Printf("%p\n", p) //指向0的指针
	p = new([]int)
	fmt.Printf("%p\n", p)
	*p = append(*p, 1, 2, 3)
	for i := 0; i < len(*p); i++ {
		fmt.Println((*p)[i])
	}
}

func main0507() {
	var num []int = []int{10, 20, 30, 40, 50}
	fun(&num)
	fmt.Println(num) //[10 20 30 40 50 1 2 3]
}
func fun(p *[]int) { //参数是切片指针
	*p = append(*p, 1, 2, 3)
}

func main0508() {
	stu := []string{"tom", "berry", "ketty"}
	for _, v := range stu {
		stu = append(stu, "李白")
		fmt.Println(v) //tom-berry-ketty
	}
	fmt.Println(stu)
}

func mainx() {
	num1 := []int{10, 5, 8, 7, 12, 34}
	num2 := []float64{1.2, 5.6, 3.4, 7.8}
	num3 := []string{"tom", "berry", "ketty"}
	//sort.Ints(num1)				//默认升序
	//sort.Float64s(num2)
	//sort.Strings(num3)
	sort.Sort(sort.Reverse(sort.IntSlice((num1)))) //反向排序（降序）
	sort.Sort(sort.Reverse(sort.Float64Slice(num2)))
	sort.Sort(sort.Reverse(sort.StringSlice(num3)))

	fmt.Println(num1) //[5 7 8 10 12 34]			[34 12 10 8 7 5]
	fmt.Println(num2) //[1.2 3.4 5.6 7.8]			[7.8 5.6 3.4 1.2]
	fmt.Println(num3) //[berry ketty tom]			[tom ketty berry]
}
