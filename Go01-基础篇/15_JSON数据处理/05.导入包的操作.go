package main

import (
	. "fmt"       //加. 省略包名
	xs "math"     //给包取别名
	_ "math/rand" //忽略此包,不使用包中的方法(下划线)
)

func main150501() {
	Println("你好啊！")
	Println(xs.Abs(200))
	//println(rand.Int(200))
}
