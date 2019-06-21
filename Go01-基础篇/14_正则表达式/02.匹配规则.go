package main

import (
	"fmt"
	"regexp"
)

func main140101() {
	//例题1：中括号,脱字符
	str := "arc asc atc ac auc"
	//1.中括号，表示一位字符
	//reg := regexp.MustCompile(`a[rst]c`) //返回正则表示式对象 反引号创建字面量对象

	//2.脱字符，不能包括中括号中的字符
	reg := regexp.MustCompile(`a[^rst]c`)
	result := reg.FindAllString(str, -1) //查找所有符合条件的结果

	//result  := reg.FindString(str) //只查出一个符合的结果
	fmt.Println(result)
}

func main140102() {
	//例题2:量词
	str := "a1c  a33c  a156c  ac   ad"

	reg := regexp.MustCompile(`a[0-9]{2}c`)  //[a33c]
	reg = regexp.MustCompile(`a[0-9]*c`)     //*表示大于等于0次 [a1c a33c a156c ac]
	reg = regexp.MustCompile(`a[0-9]+c`)     //+表示大于等于1次 [a1c a33c a156c]
	reg = regexp.MustCompile(`a[0-9]?c`)     //?表示0次或1次 [a1c ac]
	reg = regexp.MustCompile(`a[0-9]{0,5}c`) //表示0-5位的数字 [a1c a33c a156c ac]
	reg = regexp.MustCompile(`a[0-9]{2,}c`)  //表示>=2位的数字 [a33c a156c]

	result := reg.FindAllString(str, -1)
	fmt.Println(result)
}

func main140103() {
	//例题3：或者
	str := "arc,asc,atc,arsc,ac"
	//reg := regexp.MustCompile(`a[r-t]c`) //表示范围
	reg := regexp.MustCompile(`a[r|t]c`) //[arc atc] | 表示或者
	result := reg.FindAllString(str, -1)
	fmt.Println(result)
}

func main140104() {
	//例题4：匹配开头和结尾
	str := []string{"aacqqqq", "abc", "acc", "aec", "aabc", "ac", "qqasaccs"}

	reg := regexp.MustCompile(`^a[a-c]c`)  //匹配开头 aacqqqq abc acc
	reg = regexp.MustCompile(`a[a-c]c$`)   //匹配结尾 abc acc aabc
	reg = regexp.MustCompile(`^a[a-z]c$`)  //匹配开头和结尾 abc  acc aec
	reg = regexp.MustCompile(`^a[a-z]+q$`) //匹配开头和结尾 aacqqqq

	//切片数据如何匹配
	for _, v := range str {
		if reg.MatchString(v) {
			//字符串是否满足正则表达式 满足返回true 反之返回false
			fmt.Println(v)
		}
	}
}

func main140105() {
	//任意字符
	stu := []string{"李白", "李小白", "小李白", "李晓飞刀", "杜甫", "白居易"}
	reg := regexp.MustCompile(`^李.+`)   //一次或多次 李白 李小白 李晓飞刀
	reg = regexp.MustCompile(`^[^李].+`) //找出不姓李的学生

	reg = regexp.MustCompile(`^[^李|杜].+`) //找出不姓李 也不姓杜的学生 A
	reg = regexp.MustCompile(`^[^李杜].+`)  //找出不姓李 也不姓杜的学生  B 等价 A

	reg = regexp.MustCompile(`.+白$`) //.+表示任意字符 出现一次到多次

	for _, v := range stu {
		if reg.MatchString(v) {
			fmt.Println(v)
		}
	}
}

func mainmain140106_1() {
	//例题6：字符簇-匹配小数
	str := []string{"a10", "10a", "10", "3.14"}

	reg := regexp.MustCompile(`\d+\.+\d`) //匹配小数
	for _, v := range str {
		if reg.MatchString(v) {
			fmt.Println(v)
		}
	}
}

func mainmain140106_2() {
	//字符簇-匹配邮箱
	email := []string{"aa@aa.com", "aa@", "@aa", "aa@aa.", "aa.@aa.com", "aa@aa.com.cn"}
	reg := regexp.MustCompile(`^\w+@\w+\.\w+$|^\w+@\w+\.\w+\.\w+$`)
	for _, v := range email {
		if reg.MatchString(v) {
			fmt.Println(v)
		}
	}
}

//看一下视频
func main140107() {
	//例题7：匹配单词边界
	str := "This is a girl"
	reg := regexp.MustCompile(`\Bis\b`) //
	str = reg.ReplaceAllString(str, "at")
	fmt.Println(str) //That is a girl
}

func mainmain140108() {
	//匹配空白和非空白
	str := "This is a girl"
	reg := regexp.MustCompile(`\S+is\s`)
	str = reg.ReplaceAllString(str, "That ")
	fmt.Println(str)
}

//Exercise
func mainExercise() {
	//练习：匹配单词边界 匹配空白和非空白
	str := "This is a girl"
	//reg := regexp.MustCompile(`\Bis\b`)
	//str = reg.ReplaceAllString(str,"at")

	reg := regexp.MustCompile(`\S+is\s`)
	str = reg.ReplaceAllString(str, "That ")
	fmt.Println(str)
}
