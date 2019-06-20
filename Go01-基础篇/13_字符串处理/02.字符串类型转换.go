package main

import (
	"fmt"
	"strconv"
)

func main130201() {
	/*
		1.将其他类型转成字符串 布尔型转字符串
			Bool型转成字符串
	*/
	s := strconv.FormatBool(false)

	fmt.Println(s)        //false
	fmt.Printf("%T\n", s) //string

	/*
		2.将整形转成字符串
		FormatInt(数据，进制) 二进制 八进制 十进制 十六进制
	*/
	s = strconv.FormatInt(123, 2)
	fmt.Printf("%T\n", s)

	//Itoa只能转十进制 是进制转换的另一种方式
	s = strconv.Itoa(123) //trconv.Itoa把数字转为字符串
	fmt.Println(s)

	fmt.Printf("%T\n", s)

	/*
		浮点型转成字符串
		FormatFloat(数据,'f',保留小数位置(谁四舍五入),位数(64,32))
	*/
	s = strconv.FormatFloat(1.2355, 'f', 2, 64)
	fmt.Println(s)
}

func main130202() {
	//1.将字符串转成bool
	str1 := "true"
	//只能将"true" "false" 转成bool类型 有多余的数据是无效的
	str2, err := strconv.ParseBool(str1)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("%T:%v\n", str2, str2)
	}

	//2.将字符串转成整形
	str3 := "12345"
	str4, err := strconv.ParseInt(str3, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("%T:%v\n", str4, str4)
	}

	//3.将字符串转成浮点型
	str5 := "1.255"
	value, err := strconv.ParseFloat(str5, 64)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("%T:%v\n", value, value)
	}
}

func main130203() {
	b := make([]byte, 0, 1024)
	//将指定类型放在切片中
	b = strconv.AppendBool(b, true)
	b = strconv.AppendInt(b, 123, 10)
	b = strconv.AppendFloat(b, 1.24, 'f', 5, 64)
	b = strconv.AppendQuote(b, "字符串")
	fmt.Println(string(b))
}
