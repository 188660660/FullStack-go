package main

import "fmt"

func main040701() {
	//go语言中的冒泡排序
	arr := [10]int{100, 99, 66, 77, 55, 33, 110, 25, 67, 89}

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}
func main040702() {
	//go语言中的冒泡排序
	arr := [10]int{100, 99, 66, 77, 55, 33, 110, 25, 67, 89}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ { //年纪最大的排前面
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}
