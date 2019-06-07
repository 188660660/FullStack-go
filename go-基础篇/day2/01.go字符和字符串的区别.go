package main

import "fmt"

func main0101() {
	//var a byte = 'a'
	//var b string = "a"
	var c string = "as\\asas\nas\077asdas"
	//%s 遇到\0停止 但自己测试的时候遇到\0并未出现停止
	fmt.Println(c)
	fmt.Printf("%s", c)
	//fmt.Println(a==b) 不同的类型进行比较运算时会报类型不符的错误
}

func main0102() {
	var str1 string = "hello world"
	//在go语言中一个汉字算作3个字符 是为了和Linux统一处理
	var str2 string = "新中国zh"

	num1 := len(str1)
	num2 := len(str2)

	fmt.Println(str1)
	fmt.Println(num1)
	fmt.Println(num2)
}

func main0103() {
	//字符串不只可以使用双引号表述 还可以进行反引号换行书写
	var a string = `阿萨
莎莎`
	fmt.Println(a)
}

/*
扩展延伸：
Go语言的字符串有两种方式来表示：

双引号，可以使用转义字符，如s := "Go语言字符串\n不能跨行赋值"
反引号，字符串跟反引号中的格式一样，即Raw Type
s := `Go原格式字符串
    可以跨行`
Go语言的字符串是以UTF-8格式编码并存储的
*/
