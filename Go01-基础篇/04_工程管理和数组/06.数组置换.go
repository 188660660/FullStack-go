package main

import "fmt"

func main040601() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//0 9 1-10 {10,2,3,4,5,6,7,8,9,1}
	//1 8 2-9  {10,9,3,4,5,6,7,8,2,1}
	//2 7 3-8  {10,9,8,4,5,6,7,3,2,1}
	//3 6 4-7  {10,9,8,7,5,6,4,3,2,1}
	//4 5 5-6  {10,9,8,7,6,5,4,3,2,1}
	//5 4 6-5 if start>end break

	//定义起始和结束下标
	start := 0
	end := len(arr) - 1

	for ; start <= end; start++ {
		if start > end {
			break
		}
		arr[start], arr[end] = arr[end], arr[start]
		end--
	}
	fmt.Println(arr)
}

func main040602() {
	//数组大小顺序颠倒
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//定义开始和结束的下标
	start := 0
	end := len(arr) - 1
	for {
		//结束条件 关键
		if start > end {
			break //结束for程序运行
		}
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
	fmt.Println(arr)
}
