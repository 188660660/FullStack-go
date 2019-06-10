package main

import "fmt"

func test050601(slice []int) {
	fmt.Printf("%p\n", slice)
	slice[0] = 123
}
func main050601() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("%p\n", s)
	test050601(s)
	fmt.Println(s)
}

func main050602() {
	s := []int{8, 9, 7, 5, 6, 4, 2, 1, 3, 4}
	/*
		切片传参的特点：
			切片作为函数传参数 是地址传递 形参可以改变实参的值
			在项目开发中 建议使用切片 代替数组的使用
	*/

	BubbleSortDesc0601(s)
	fmt.Println(s)
}

func BubbleSortDesc0601(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func test050602(slice []int) []int {
	/*
		！如果使用 append 操作切片可能会改变切片的地址 需要使用返回值
		如果想不改变原地址 那就不在函数中进行追加？
	*/
	slice = append(slice, 1, 2, 3)
	fmt.Printf("%p\n", slice)
	fmt.Println(slice)
	return slice
}

func main050603() {
	s := []int{1, 2, 3}
	fmt.Printf("%p\n", s)
	s = test050602(s)
	fmt.Println(s)
	fmt.Printf("%p\f", s)
}
