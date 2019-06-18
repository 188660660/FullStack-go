package main

import "fmt"

/*
	recover注意事项：
		1.recover()使用一定要在错误出现之前 进行拦截
		2.通过匿名函数和recover()进行错误拦截
	recover的特性：
		1.recover可以从panic 异常中重新获取控制权
		2.recover()返回值为接口类型
	recover的作用：
		1.提高程序的友好性和健壮性
*/
func demo100401(param int) {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("demo100401函数异常：13-15行", err)
		}
	}()

	var p *int
	*p = 123 //err 无效的内存地址 recover是有作用范围的

	fmt.Println("1")

	//如果传递超出数组的下标值会导致数组下标越界
	var arr [10]int
	arr[param] = 100 //panic常用作系统处理 会导致程序崩溃
	fmt.Println(arr)
}

func demo100402() {
	fmt.Println("传说中的暴龙兽！")
}

func main100401() {
	demo100401(11)
	demo100402()
}

func fun100401(age int) {
	if age >= 45 || age <= 18 {
		panic("辞退该员工！")
	}
	fmt.Println("允许继续留任")
}
func main100402() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	var age int
	fmt.Scan(&age)
	fun100401(age)
}

func test100401(num int) (value int) {
	value = 10 / num
	return
}

//写一个recove接口的案例
func main100403() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	test100401(0)
}
