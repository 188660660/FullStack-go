package main

import "fmt"

/*
	!总结：
		特性：
			A.在GO语言中 所有的函数都是全局函数 都可以被项目中所有文件使用 在项目中 函数名必须是惟一的
			这句概念是我看一个视频的讲师是的 后来询问大牛有错误↓为他所述 后续验证是否正确
				1.不是全局函数，不能被所有文件引用。
				go函数受到包管理，通过函数首字母大小写，控制包内外是否可见。
				其次函数分为全局函数和struct对象函数，struct就不能直接引用。
				2.函数名字不必全局唯一。全局函数包内唯一。对象函数对象内唯一。
*/
//此处存疑 全局为什么还报未定义错误
//func test3(args ...int) {
//	fmt.Println(len(args))
//}

//func main() {
//	test3(1,2,3,4)
//}
func main030901() {
	//test3(1,2,3,4)
	//test3(5,6,7,8)
	test4(67, 81)
}

func test4(a int, b int) {
	test5(a, b)
}
func test5(a int, b int) {
	sum := a + b
	fmt.Println(sum)
}

func main030902() {
	//test6(10, 20, 30, 40)
	test7(10, 20, 30, 40)
}

//知识点：函数的作用域是全局的 可以在项目中的任意文件使用 不需要在意先后顺序
func test6(args ...int) {
	sum := 0
	for _, value := range args {
		sum += value
	}
	fmt.Println(sum)
}

//知识点：函数的传递方式
func test7(args ...int) {
	//test8(args[0:]...) //传递参数 必须常用这种方式 在本函数中使用不用加....
	test8(args[1:3]...)
	return
	//如果函数的参数为不定参 传递的方式为args[0:]
	fmt.Println(args[:0])  //[]
	fmt.Println(args[0:])  //[10 20 30 40] | 从0下标开始
	fmt.Println(args[1:])  //[20 30 40] |从1下标开始
	fmt.Println(args[:1])  //[10] | 开始取一个下标
	fmt.Println(args[1:3]) //[20 30] 从1下标 取到3下标？

	fmt.Println(args[3:4]) //[40]

}

func test8(args ...int) {
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
}
