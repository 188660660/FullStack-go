package main

import "fmt"

/*
	func add (a int , b int) int{
		return a + b
	}
*/
//方法的语法: func (方法接收者) 方法名(参数列表) 返回值类型
/*
	方法使用步骤：
		1.定义函数类型
		2.为已存在的函数类型起别名
*/

type Int int //1.预处理 2.编译 3.汇编 4.链接

func (c Int) add(b Int) Int {
	return c + b
}
func main080601() {
	var a Int = 10
	value := a.add(20)

	fmt.Println(value)
}

//联系demo 创建一个冒泡排序的方法
type BubbleSort []int //定义一个全局的数组切片类型

func (arr BubbleSort) BubbleSortDesc(args BubbleSort) BubbleSort {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
	return args
}

func main080602() {
	var arr BubbleSort = []int{3, 2, 1}
	arr = arr.BubbleSortDesc([]int{55, 8, 7, 6, 5, 4, 3, 2, 1, 0})
	fmt.Println(arr)
}
