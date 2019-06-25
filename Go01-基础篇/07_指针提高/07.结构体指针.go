package main

import "fmt"

type student070701 struct {
	name string
	id   int
	age  int
	sex  string
}

func main070701() {
	/*
		定义结构体变量:
			var 结构体名 结构数据类型
		var stu Student=Student{
			id:101,name:"多啦A梦",
			age:100,sex:"男"
		}
		结构体变量.成员 = 值
	*/
	var stu = student070701{
		id: 101, name: "多啦A梦",
		age: 100, sex: "男",
	}

	//定义结构体指针指向变量的地址
	var p *student070701

	//结构体指针指向变量的地址
	p1 := &stu
	p2 := &stu.name
	p3 := &stu.id

	fmt.Printf("%T\n", p1) //*main.student070701
	fmt.Printf("%T\n", p2) //*string
	fmt.Printf("%T\n", p3) //*int

	fmt.Printf("%T\n", p)    //*main.student070701
	fmt.Printf("%T\n", &stu) //*main.student070701

	//结构体成员地址
	fmt.Printf("%p\n", &stu.name) //0xc00007a300
	fmt.Printf("%p\n", &stu.id)   //0xc00007a310
	fmt.Printf("%p\n", &stu.age)  //0xc00007a318
	fmt.Printf("%p\n", &stu.sex)  //0xc00007a320
}

func main070702() {
	var stu = student070701{
		id: 101, name: "多啦A梦",
		age: 100, sex: "男",
	}

	var p *student070701
	p = &stu

	//通过结构体指针间接操作结构体成员
	(*p).name = "大雄"
	//通过指针可以直接操作结构体成员
	//p.name = "静香"
	p.age = 18
	p.sex = "男"

	fmt.Println(stu)
}

func main070703() {
	//结构体切片
	var stu = make([]student070701, 3)

	p := &stu //*[]student070701 结构体切片指针

	fmt.Println(p)
	fmt.Printf("%T\n", p)

	(*p)[0] = student070701{"张", 1, 18, "男"}
	(*p)[1] = student070701{"王", 1, 18, "男"}
	(*p)[2] = student070701{"里", 1, 18, "男"}

	fmt.Println(*p)       //[{张 1 18 男} {王 1 18 男} {里 1 18 男}]
	fmt.Printf("%T\n", p) //*[]main.student070701

	*p = append(*p, student070701{"小猪佩奇", 1, 10, "女"})

	fmt.Println(stu)      //[{张 1 18 男} {王 1 18 男} {里 1 18 男} {小猪佩奇 1 10 女}]
	fmt.Printf("%T\n", p) //*[]main.student070701

	//循环方法一：
	for i := 0; i < len(*p); i++ {
		fmt.Printf("下标：%d 值：%v\n", i, (*p)[i])
	}

	fmt.Println("----------------")

	//循环方法二：
	for k, v := range *p {
		fmt.Printf("%T\n", v)
		fmt.Printf("下标：%d 值：%v\n", k, v)
	}
}
