package main

import "fmt"

func main040401() {
	/*
		var 变量名 数据类型
		数组的定义和使用：
				1.var 数组名[元素个数]数据类型 = [元素个数]数据类型{元素内容}
			      var arr[3]int = [3]int{1,2,3}
				2.数组名[下标] = 值 ：如：arr[0] = 1 ; arr[1] = 2
		数组下标是从0开始的 数组的最大长度可以使用len(array)-1求出
	*/
	//定义一个数组并测试
	var arr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])

	//打印数组名 可以显示出所有的数组元素
	fmt.Println(arr)

	//打印数组信息 - 方法一
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	//打印集合(数组)信息 - 方法二
	for _, value := range arr { //返回值 下标
		fmt.Println(value)
	}

	/*
		!小结：len函数的作用：
				A：可以len(string)输出字符串的有效长度
				B：可以len(...int)获取不定参函数的个数
				C：可以len([10]int)计算数组元素的个数
	*/
	fmt.Println(len(arr)) //数组10 即元素个数
}

func main040402() {
	//在数组定义时 可以初始化部分元素的值
	var arr1 [10]int = [10]int{1, 2, 3, 4, 5}
	//使用自动类型推导创建数组
	arr2 := [10]int{5, 4, 3, 2, 1}
	//使用自动类型推导 可以根据元素个数创建数组
	arr3 := [...]int{1, 2, 3, 4, 5, 6, 7} //个人感觉这种方式较优

	for _, v := range arr1 {
		fmt.Println(v)
	}
	for _, v := range arr2 {
		fmt.Println(v)
	}
	for _, v := range arr3 {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", arr1) //%T获取本身的数据类型 [10]int 1,2,3,4,5,0,0,0,0,0
	fmt.Printf("%T\n", arr2) //%T获取本身的数据类型 [10]int 5,4,3,2,1,0,0,0,0,0
	fmt.Printf("%T\n", arr3) //%T获取本身的数据类型 [7]int 1,2,3,4,5,6,7
}

func main040403() {
	//创建方法一 常规常见
	var arr1 [5]int = [5]int{1, 2, 3, 4, 5} //[5]在定义时叫元素个数 0~len(arr)-1 叫做数组下标
	//创建方法二 类型推导
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", arr1) //[5]int
	fmt.Printf("%T\n", arr2) //[5]int

	/*
		 //数组下标越界:
			var arr [5]int = [5]int{1, 2, 3, 4, 5}
			arr[5] = 100  //下标最长只到4 即arr[4] //err
			arr[-1] = 100 //下标从0开始 不可以为负数 //err
	*/

	//在GO语言中数组在定义后 元素的个数就已经固定 不能够添加
	fmt.Println(arr1)

	//在GO语言中数组元素是一个【常量】 不允许赋值 数组名代表整个数组
	//arr1 = 10 //err

	//在GO语言中 数组名也可以表示数组的首地址
	//即数组名的 内存地址 = 数组名[0] 所在的内存地址
	fmt.Printf("%p\n", &arr1) //0xc00008e030
	fmt.Printf("%p\n", &arr2) //0xc00008e060

	fmt.Printf("%p\n", &arr1[0]) //0xc00008e030
	fmt.Printf("%p\n", &arr1[1]) //0xc00008e038

	fmt.Printf("%p\n", &arr2[0]) //0xc00008e060
	fmt.Printf("%p\n", &arr2[1]) //0xc00008e068
}
