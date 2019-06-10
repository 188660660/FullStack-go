package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	rand.Seed()  //使用给定的seed来初始化生成器到一个确定的状态。
	rand.Init()  //返回一个取值范围在[0,n)的伪随机int值，如果n<=0会panic。
	time.Now().UnixNano() //即从时间点January1,1970 UTC 到时间点t所经过的时间（单位纳秒）。
	time.Now().Nanosecond() //返回对应的那一秒内的纳秒偏移量，范围[0,999999999]。
*/
func main111() {
	//1.导入头文件 math/rand time
	//2.添加随机数种子
	//3.使用随机数
	start := time.Now().Nanosecond() //开始时间
	rand.Seed(time.Now().UnixNano())
	fmt.Println(start)

	for i := 0; i < 10000; i++ {
		//根据时间 1970.1.1.0.0.0
		fmt.Println(rand.Intn(100)) //0-99 最大值-1 取余 %
	}
	fmt.Println(time.Now())
	fmt.Println(time.Now().Second())

	end := time.Now().Nanosecond()
	fmt.Println(end - start)

}

func main040902() {
	//验证伪随机数 每次的结果值输出都是一样的 要做随机数混淆处理
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(123))
	}
	/*

	 */
	//time.Now().UnixNano() 精确到纳秒，通过纳秒就可以计算出毫秒和微秒
	/*
		!补充：
		如果不使用rand.Seed(seed int64)，每次运行，得到的随机数会一样，程序不停止，一直获取的随机数是不一样的；
		每次运行时rand.Seed(seed int64)，seed的值要不一样，这样生成的随机数才会和上次运行时生成的随机数不一样；
		rand.Intn(n int)得到的随机数int i，0 <= i < n。
	*/
}

func main040903() {
	//创建一个 100的随机数
	var arr [10]int
	//send    使用给定的seed来初始化生成器到一个确定的状态。
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	//BubbleSortAsc09(arr) //结果值为什么不正确 找原因
	//BubbleSortDesc09(arr)

	arr = BubbleSortAsc09(arr)
	arr = BubbleSortDesc09(arr)
	fmt.Println(arr)
}

func BubbleSortAsc09(arr [10]int) [10]int { //冒泡排序正序
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
	return arr
}

func BubbleSortDesc09(arr [10]int) [10]int { //冒泡排序倒序
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main040904() {
	//创建一个数组 用来进行保存
	var arr [10]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr)-1; i++ {
		arr[i] = rand.Intn(101) //rand.Init 和 rand.Init一个带参数 一个不带参数？
	}
	fmt.Println(arr)
	res := BubbleSortDesc09(arr) //将res进行倒序排序保存
	fmt.Println(res)
}
