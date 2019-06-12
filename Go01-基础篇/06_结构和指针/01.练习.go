package main

import "fmt"

func main060101() {
	//计算字母出现了几次
	var arr [20]byte //字符切片
	for i := 0; i < len(arr); i++ {
		fmt.Scanf("%c", &arr[i])
	}
	m := make(map[byte]int)

	for i := 0; i < len(arr); i++ {
		m[arr[i]]++
	}

	for k, v := range m {
		if v > 0 {
			fmt.Printf("%c %d\n", k, v)
		}
	}

}
