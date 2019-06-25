package main

import "fmt"

func main050201() {
	/*
		一维数组语法：
					var 数组名 [元素个数]元素类型
					= [元素个数]数据类型{元素参数一,元素参数二,...}
	*/
	var demo1 = [3]int{1, 2, 3}
	/*
		二维数组语法：
				    var 数组名 [行元素个数][列元素个数]数据类型
					= [行元素个数][列元素个数]数据类型{{A,B,C},{},{}}
	*/
	var demo2 = [3][5]int{{1, 2, 3, 4, 5}, {2, 3, 4, 5, 6}}

	//通过行和列下标找到具体的元素
	fmt.Println(demo2[0][4])
	fmt.Println(demo1)
	fmt.Println(demo2)

	//打印二维数组 A 遍
	for i := 0; i < len(demo2); i++ {
		for j := 0; j < len(demo2[i]); j++ {
			fmt.Print(demo2[i][j], " ")
		}
		fmt.Println()
	}
	//打印二维数组 B遍
	for i := 0; i < len(demo2); i++ {
		for j := 0; j < len(demo2[i]); j++ {
			fmt.Print(demo2[i][j], " ")
		}
		fmt.Println()
	}
}

func main050202() {
	var arr [3][4]int
	//len(二维数组名)：获取到的是arr函数
	fmt.Println(len(arr))

	fmt.Println(arr[0]) //表示二维数组的第一行
	//len(二维数组名(下标n))
	fmt.Println(arr[2][0]) //表示二维数组的第三列 第一列
	fmt.Println(arr[2])    //表示二维数组的第三列的长度

	//打印二维数组 C遍
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}
