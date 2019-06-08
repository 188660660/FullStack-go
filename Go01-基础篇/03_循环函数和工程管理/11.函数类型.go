package main

import "fmt"

/*
	！总结：
		什么是函数类型？
			定义：定义使用函数的方式,就是函数类型，
			特性: 所谓的函数类型就是将函数作为一种类型，可以用它来定义变量
*/
func test10() {
	fmt.Println("传说中的暴龙兽")
}

func test11(a int, b int) {
	fmt.Println(a + b)
}

func test12(a int, b int) int {
	return a + b
}

func test13() {
	fmt.Println("传说中的基尔兽")
}

/*
	A.type可以定义函数类型
	B.type可以为已存在的函数起别名
*/
type FUNCTYPE func()
type FUNCTEST func(int, int)
type funcdemo func(int, int) int

func main1101() {
	//定义函数类型和变量
	var f FUNCTYPE
	f = test10
	f()

	//f = test11 /err 如果函数类型不一致就会报错
	//f()

	var f1 FUNCTEST
	f1 = test11
	f1(10, 20)

	var f2 funcdemo
	f2 = test12
	value := f2(25, 86)
	fmt.Println(value)

	//类型和参数一致可以重复使用
	var f3 FUNCTYPE
	f3 = test13
	f3()
}

func test14(a int, b int) {
	fmt.Println(a + b)
}
func main() {
	//函数调用
	test14(25, 30)

	/*
		如果使用print打印 将会出现一个地址
		！函数名本身就是一个指针类型数据 在内存中的【代码区】进行存储 使用的时候在栈区
			(变量存储在栈区)->四大区补充：代码区 数据区 堆区 栈区
	*/
	fmt.Printf("%T\n", test14) //0x491370

	//自动类型推导创建函数类型
	f1 := test14
	f1(10, 20)
	fmt.Printf("%T\n", f1)

	//直接定义函数类型
	var f2 func(int, int)
	f2 = test14
	f2(80, 90)
	//fmt.Println(f2) //0x492f70

}
