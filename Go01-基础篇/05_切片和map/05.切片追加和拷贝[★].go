package main

import "fmt"

func main050501() {
	var s []int
	//s[0] = 123
	/*
		在使用append添加数据时 切片的地址可能或发生变量
		如果容量扩充导致输出存储溢出 切片会自动找寻新的空间存储数据
		同时也会将之前的数据进行释放
		使用append添加数据
	*/
	s = append(s, 1, 2, 3, 4, 5, 6)

	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Printf("%p\n", s)

	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Printf("%p\n", s)
}

func main050502() {
	s1 := []int{1, 2, 3, 4, 5}
	//var s2 []int //0 //0
	s2 := make([]int, 5)
	//将s1切片中的数据 拷贝到s2中 s2中要有足够的容量
	//使用拷贝操作后s1 和 s2是两个独立的空间 不会相互影响

	copy(s2, s1) //拷贝后的地址,拷贝源地址
	s2[2] = 123

	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Printf("%p\f", s1)
	fmt.Printf("%p\f", s2)
}

func main050503() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, 2, 100)

	//如果要拷贝具体的切片中的片段需要使用切片中的截取
	copy(s2, s1[2:3]) //4
	fmt.Println(s2)

}

/*
	延伸扩展 - make()函数 [分配内存]
		第一个参数是类型，第二个参数是分配的空间，
		第三个参数是预留分配空间
		a:=make([]int, 5, 10)， len(a)输出结果是5，cap(a)输出结果是10

预留的空间需要[重新切片]才可以使用
package main

import "fmt"

func main(){
	a := make([]int, 10, 20)

	fmt.Printf("%d, %d\n", len(a), cap(a))

	fmt.Println(a)
	b := a[:cap(a)] ★
	fmt.Println(b)
}

10, 20
[0 0 0 0 0 0 0 0 0 0]
[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
*/
