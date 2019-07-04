package main

import (
	"fmt"
	"time"
)

func main040101() {
	str := []string{"a", "b", "c"}
	for _, v := range str {
		go func() {
			fmt.Println(v)
		}()
	}
	time.Sleep(3 * time.Second) //全部都是 a bc 的组合结果
}
