package main

import "fmt"

func main() {
	//定义一个数组
	//var arr [10]int = [10]int{85,76,99,120,85,66,77,99,25,16} //找出最重的一只小猪
	arr := [10]int{85, 76, 99, 120, 85, 66, 77, 99, 25, 16}

	max := arr[0] //假设arr[0] 为体重最重的那只小猪
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	fmt.Println(max)
}
