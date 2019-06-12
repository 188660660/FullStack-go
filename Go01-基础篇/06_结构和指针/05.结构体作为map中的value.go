package main

import "fmt"

//定义全局结构体
type stu060501 struct {
	name  string
	age   int
	score int
}

func main060601() {

	//定义map
	m := make(map[int]stu060501)

	m[101] = stu060501{"猪八戒", 18, 60}
	m[102] = stu060501{"沙和尚", 28, 60}
	m[103] = stu060501{"孙悟空", 38, 60}

	fmt.Println(m)
	fmt.Printf("%T\n", m)      //map[int]main.stu060501
	fmt.Printf("%T\n", m[101]) //main.stu060501

	//打印map 结构体
	for k, v := range m {
		fmt.Printf("键：%d 值:%v\n", k, v) //%v会打印实际类型的值。
	}
}

func main() {
	m := make(map[int][]stu060501)

	m[1] = append(m[101], stu060501{"曹操", 50, 68})
	m[2] = append(m[101], stu060501{"张飞", 50, 68})
	m[3] = append(m[102], stu060501{"刘备", 50, 68})
	m[4] = append(m[103], stu060501{"孙权", 50, 68})

	fmt.Printf("%T\n", m) //map[int][]main.stu060501
	fmt.Println(m)        //map[1:[{曹操 50 68}] 2:[{刘备 50 68}] 3:[{孙权 50 68}]]

	//m[1][101].name="张文远" 判断有没有这个名字

	for k, v := range m {
		for i, data := range v {
			if data.name == "曹操" {
				fmt.Println(k, i, data)
			}
			fmt.Println("key:", k, "index", i, "value", data)
		}
	}
}
