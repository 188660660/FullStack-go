package main

import "fmt"

/*
	!思考：if和switch的区别？
			if 可以嵌套 可以判断区间  执行效率比较低
			switch 执行效率高  不能嵌套和区间判断
	特性：
		1.switch 选择项可以是一个整型变量
		2.swich中的值不能是浮点型数据 浮点型数据是一个约等于的数据
*/
func main021201() {
	var w int
	fmt.Scan(&w)

	switch w {
	case 1:
		fmt.Println("今天是星期一")
	case 2:
		fmt.Println("今天是星期二")
	case 3:
		fmt.Println("今天是星期三")
	case 4:
		fmt.Println("今天是星期四")
	case 5:
		fmt.Println("今天是星期五")
	case 6:
		fmt.Println("今天是星期六")
	case 7:
		fmt.Println("今天是星期七")
	default:
		fmt.Println("请输入正确的数值！")
	}
}

func main021202() {
	//switch判断成绩是否合格
	fmt.Println("请输入查询成绩！")
	var score int
	fmt.Scanf("%d\n", &score)
	switch score >= 60 {
	case true:
		fmt.Println("此次考试成绩及格！")
	case false:
		fmt.Println("此次考试成绩不及格！")
	}
}

func main021203() {
	var score int
	fmt.Println("请输入考试成绩！")
	fmt.Scan(&score)

	switch score / 10 {
	case 10:
		//fmt.Println("A")
		fallthrough //让switch执行下一个分支的代码  如果不写 执行到下一个分支就会自动停止
	case 9:
		fmt.Println("A")
	case 8:
		fmt.Println("B")
	case 7:
		fmt.Println("C")
	case 6:
		fmt.Println("D")
	default:
		fmt.Println("E")
	}
}
