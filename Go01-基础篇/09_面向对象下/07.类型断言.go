package main

import "fmt"

/*
	value, ok := x.(T)
	x表示要断言的接口变量；
	T表示要断言的目标类型；
	value表示断言成功之后目标类型变量；
	ok表示断言的结果，是一个bool型变量，true表示断言成功，false表示失败，如果失败value的值为nil。

	网上百度 golang中x.(type)只能在switch中使用
*/
func main090701() {
	arr := make([]interface{}, 6)
	arr[0] = 1
	arr[1] = 3.14
	arr[2] = "字符串"
	arr[3] = 'a'
	arr[4] = true
	arr[5] = []int{1, 2, 3}

	//方法一 for 循环实现
	for _, v := range arr {
		//data,ok := v.([]int)
		//if ok {
		//	fmt.Println(data)
		//}
		if data, ok := v.(int); ok {
			fmt.Println(data, "是整形")
		} else if data, ok := v.(float64); ok {
			fmt.Println(data, "是浮点型")
		} else if data, ok := v.(byte); ok {
			fmt.Println(data, "是字符型")
		} else if data, ok := v.(string); ok {
			fmt.Println(data, "是字符串型")
		} else if data, ok := v.(bool); ok {
			fmt.Println(data, "是布尔型")
		} else if data, ok := v.([]int); ok {
			fmt.Println(data, "是整形切片")
		} else {
			fmt.Println(data, "未知类型")
		}
	}

	//方法二 使用switch 实现
	for _, v := range arr {
		switch value := v.(type) {
		case int:
			fmt.Println(value, "是整形")
		case float64:
			fmt.Println(value, "是字符串型")
		case []int:
			fmt.Println(value, "是整形切片")
		default:
			fmt.Println("未知类型")
		}
	}
}
