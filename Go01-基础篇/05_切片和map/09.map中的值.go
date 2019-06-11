package main

import "fmt"

func main050901() {
	// m: = make(map[byte]int)
	m := make(map[int]string, 1)
	//map中的value值允许重复出现
	m[0] = "萌萌哒"
	m[1] = "萌萌哒"
	m[2] = "棒棒哒"

	fmt.Println(m[2])
	//在map中 如果出现没有定义的map 默认值为空 nil
	fmt.Println("---", m[3], "---")

	k, v := m[3]
	if v {
		fmt.Println(k)
	} else {
		fmt.Println("没有对应值")
	}

	//循环输出map值
	for k, v := range m {
		fmt.Println(k, "---", v)
	}
}
