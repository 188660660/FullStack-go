package main

import "fmt"

func test0801(arr [3]int) {
	arr[2] = 123
	for _, v := range arr {
		fmt.Println(v)
	}
	fmt.Println(arr)
	fmt.Printf("%p\n", &arr) //取地址
}
func main040801() {
	//指定数组下标进行初始化操作
	//第一种常规方式
	var arr1 [3]int = [3]int{1, 2, 3}
	//第二种推导方式
	//arr2 := [3]int{3,2,1}

	/*
		!总结：
			1.数组作为函数参数传递是传递的数组名
			2.数组作为函数参数传递是传递的值 而不是指向地址
	*/
	test0801(arr1)
	fmt.Printf("%p\n", &arr1)
	//fmt.Println(arr1)
	//fmt.Println(arr2)
}

func main040802() {
	//定义一个无序数组 实现函数冒泡排序
	arr := [10]int{9, 1, 5, 6, 7, 3, 10, 2, 4, 8}

	asc := BubbleSortAsc(arr)
	desc := BubbleSortDesc(arr)

	fmt.Println(asc)  //正序排序
	fmt.Println(desc) //倒序排序
}

//冒泡排序正序
func BubbleSortAsc(arr [10]int) [10]int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
	return arr
}

//冒泡排序倒序
func BubbleSortDesc(arr [10]int) [10]int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
