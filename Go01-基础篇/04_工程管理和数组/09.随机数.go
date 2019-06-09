package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

func main() {
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
