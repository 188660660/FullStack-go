package main

import "fmt"

/*
	map是干嘛的？
		map 也就是 Python 中字典的概念，它的格式为“map[keyType]valueType”。
 		map 的读取和设置也类似 slice 一样，通过 key 来操作，
       只是 slice 的index 只能是｀int｀类型，而 map 多了很多类型，
       可以是 int ，可以是 string及所有完全定义了 == 与 != 操作的类型。
*/
func main050801() {
	//map[keytype]valuetype

	//var m map[int]string
	m := make(map[int]string, 1)
	/*
		1.map[key] key是一个 基本数据类型[基本数据类型：整形 浮点型 字符串型 字符型 布尔型]
		2.map的长度是自动扩容的
		3.map中的数据是无序存储的
	*/
	m[100] = "中国"
	m[99] = "日本"
	m[56] = "韩国"

	//for k,v k=>key v=>value
	for k, v := range m {
		fmt.Printf("下标：%d 值：%s \n", k, v)
	}
	fmt.Println(m)
}

func main050802() {
	m := make(map[string]int)

	m["传说"] = 100
	m["稀有"] = 88

	//可以通过验证 key 对应的value是否有值 根据条件作出相应操作
	_, v := m["稀有"]

	if v {
		fmt.Print("该值key已经存在！")
	} else {
		m["传说"] = 100
	}

	fmt.Println(m)        //map[传说:100 稀有:88]
	fmt.Printf("%T\n", m) //map[string]int
	fmt.Printf("%p", m)   //0xc00007a300

	/*
			map使用注意事项：
				1.map中的 key和value 不能翻过来操作
		         m[10] = "史诗"  //err 因为定义时就已经规定好了类型
				2.map在定义时 key 的值是唯一的 不允许重复
	*/
}
