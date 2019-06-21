package main

import "fmt"

/*
	defer的作用：
		defer语句经常被用于处理成对的操作，
		如打开、关闭、连接、断开连接、加锁、释放锁。
		通过defer机制，不论函数逻辑多复杂，都能保证在任何执行路径下，
		资源被释放。释放资源的defer应该直接跟在请求资源的语句后

	执行顺序：
		defer 延迟调用 后进先出 先进后出 从后往前
*/
func main100301() {
	//A D C B
	fmt.Println("A")
	defer fmt.Println("B")
	defer fmt.Println("C")
	fmt.Println("D")
}

func test100301() {
	fmt.Println("传说中的暴龙兽！")
}

func test100302() {
	fmt.Println("传说中的基尔兽！")
}

var value int

func test100303(a int, b int) {
	value = a + b
}

func test100304(value int) {
	fmt.Println(value)
}

func test100305(a int, b int) (value int) {
	value = (a + b) * 2
	fmt.Println(value)
	return
}

func main100302() {
	//defer test100301()
	//函数中有返回值 不能使用defer延迟调用 ★ 存疑
	//defer test100305(10,20)
	//fmt.Println(v)

	test100302()
	defer test100303(10, 20)
	test100304(value)
}

func main100303() {
	a := 10
	b := 20
	//定义函数类型变量
	//f := func(a int, b int) {
	//	fmt.Println(a)
	//	fmt.Println(b)
	//}

	//f(a,b)
	defer func(a int, b int) {
		fmt.Println(a)
		fmt.Println(b)
	}(a, b)

	a = 100
	b = 200

	fmt.Println(a)
	fmt.Println(b)
}

func main100304() {
	/*
		内存空间的管理方式有两种
		第一种：队列(先进先出)
		第二种：出栈(先进后出)
	*/
	//即使defer后面的语句崩溃，已经加载的defer语句也是要依次执行（因为在发生错误之前已经加载）
	defer fmt.Println("锄禾日当午") //执行，原因是panic前已经加载
	defer fmt.Println("汗滴禾下土") //执行
	panic("这是一个致命的错误")
	defer fmt.Println("谁知盘中餐") //不执行,panic以后不能加载
	/*
	   汗滴禾下土
	   锄禾日当午
	   panic: 这是一个致命的错误
	*/
}

func main100305() {
	a, b := 10, 20
	defer func() {
		fmt.Println("匿名函数：", a, b) //  100,200
	}()
	a, b = 100, 200
	fmt.Println("main函数：", a, b) //100,200
}

func main100306() {
	//测试
	//fmt.Println("A")
	//defer fmt.Println("B")
	//defer fmt.Println("C")
	//fmt.Println("D")

	fmt.Println("A")
	defer func() {
		fmt.Println("B")
		fmt.Println(1)
		fmt.Println("C")
	}()
	fmt.Println("D")
}
