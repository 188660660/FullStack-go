package main

import "fmt"

func main051001() {
	//第一种直接定义赋值的写法
	m1 := map[int]string{
		100: "小明",
		101: "小张",
		102: "小王", //很行 结束 需要多加一个逗号
	}
	fmt.Printf("%T\n", m1) //map[int]string %T打印数据类型
	fmt.Println(m1)

	//第二种直接定义赋值的写法
	m2 := map[int]string{100: "小明", 101: "小王", 102: "小李"}
	fmt.Printf("%T\n", m2)
	fmt.Println(m2)

	//删除map元素的方法 delete(map,key)
	delete(m1, 101)
	fmt.Println(m1)

	//delete在删除map的元素的时候 下标如果不存在 也不会报错
	delete(m1, 999)
	delete(m1, 000)
	fmt.Println(m1)
}
