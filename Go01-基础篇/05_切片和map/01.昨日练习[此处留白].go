package main

import (
	"fmt"
	"math/rand"
	"time"
)

//未完成
func main050101() {
	//通过键盘获取20个小写字母 统计个数
	//获取用的输入
	var str [20]byte
	for i := 0; i < len(str)-1; i++ {
		fmt.Scanf("%c", &str[i]) //%c 表示字符型数据类型 Scanf从标准输入扫描文本
	}
	fmt.Printf("%c", str)

	var ch [26]int //出现范围 26个英文字母 0-25 下标
	//记录字母出现的个数
	for i := 0; i < len(str); i++ {
		ch[str[i]-'a']++ //此处不理解
	}
	//打印字母出现的个数
	for i := 0; i < len(ch); i++ {
		if ch[i] > 0 {
			fmt.Printf("字母%c,一共出现%d次\n", 'a'+i, ch[i])
		}
	}
}

//随机一注双色球彩票信息 红球6个 1-33 不能重复
func main() {
	//获取随机数种子
	rand.Seed(time.Now().UnixNano())

	var redball [6]int
	for i := 0; i < len(redball); i++ {
		temp := rand.Intn(32) + 1
		for j := 0; j < i; j++ {
			if temp == redball[j] {
				temp = rand.Intn(32) + 1
				j = -1
				continue
			}
		}
		redball[i] = temp
	}
	fmt.Println(redball, "+", rand.Intn(16)+1)
}
