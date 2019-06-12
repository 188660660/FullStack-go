package main

import "fmt"

func main021101() {
	var score int
	fmt.Scan(&score)
	/*
		!总结：
			if 条件判断根据是否满足条件指向对应的代码
			else 作为if 补充条件 如果条件不满足执行else代码
				格式 if 表达式{
					代码体
				}else {
					代码体
				}
	*/
	if score >= 60 {
		fmt.Println("你的这次考试及格了！")
	} else {
		fmt.Println("这次考试不及格")
	}
}

func main021102() {
	/*if语句嵌套
	if else if
	else 配对和if 一起使用  选择同级别下的if进行配对使用
	*/
	var score int
	fmt.Scan(&score)

	if score <= 700 {
		if score >= 600 {
			fmt.Println("北大青鸟")
		} else if score >= 500 {
			fmt.Println("新东方技术学院")
		} else {
			fmt.Println("蓝翔技校")
		}
	} else if score <= 800 {
		if score >= 750 {
			fmt.Println("北大")
		} else {
			fmt.Println("清华")
		}
	} else {
		fmt.Println("外太空研究院")
	}
}

func main021103() {
	a := 10
	if a > 5 {
		fmt.Println(a)
	}

	//采用就近原则 找到上面尚未配对的if进行匹配操作
	if a > 8 {
		fmt.Println("测试")
	} else {
		fmt.Println("其他")
	}
}

func main021104() {
	//小案例：比较三个数的大小
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a > b {
		if a > c {
			fmt.Printf("A：%d的数值最大", a)
		} else {
			fmt.Printf("C：%d的数值最大", c)
		}
	} else {
		if b > c {
			fmt.Printf("B: %d的数值最大", b)
		} else {
			fmt.Printf("C：%d的数值最大", c)
		}
	}
}
