package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main1401() {
	//1.测试一个字符串是否符合一个表达式
	match, _ := regexp.MatchString("p([a-z])+ch", "peach")
	fmt.Println(match) //true

	//2.上面我们是直接使用字符串,但对于一些其他的正则任务，你需要使用Compile 一个优化的Regexp结构体
	r, _ := regexp.Compile("p([a-z]+)ch")
	//2-1：这个结构体有很多方法 这里类似我们前面看到的一个匹配测试
	fmt.Println(r.MatchString("peach")) //true
	//2-2：这是查找匹配字符串的
	fmt.Println(r.FindString("peach hellow")) //peach
	//2-3:这个也是查找第一次匹配的字符串的 当时返回的匹配开始和结束位置索引 而不是匹配的内容

	fmt.Println(r.FindStringIndex("peach hellow")) //[0 5]
	//Submatch 返回完全匹配和局部匹配的字符串。例如,此处会返回p([a-z]+)ch 和`([a-z]+)`的信息
	fmt.Println(r.FindStringSubmatch("peach hellow")) //[peach ea]
	//类似的 这个返回完全匹配和局部匹配的索引位置
	fmt.Println(r.FindStringSubmatchIndex("peach hellow")) //[0 5 1 3]

	//带ALL的这个函数返回所有的匹配项，而不仅仅是首次匹配项，例如查找匹配表达式的所有线
	fmt.Println(r.FindAllString("peach punch pinch", -1)) //[peach punch pinch]
	//All 同样对应到上面的所有函数
	fmt.Println(r.FindAllStringSubmatch("peach punch pinch", -1)) //[[peach ea] [punch un] [pinch in]]

	//这个函数提供一个正整数来限制匹配次数
	fmt.Println(r.FindAllString("peach punch pinch", 2)) //[peach punch]
	//上面的例子中,我们使用了字符串作为参数，并使用了路MathString这样的方法。我们也可以提供[]byte参数并将string从函数中去掉
	fmt.Println(r.Match([]byte("peach")))
	//创建正则表达式常量时，可以使用Compile的变体MustCompile.因为Compile返回两个值 不能适用于常量
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r) //p([a-z]+)ch

	//regexp包 也可以用来替换部分字符串为其他值
	fmt.Println(r.ReplaceAllString("a peach", "<frult>"))

	//Func 变量允许传递匹配内容到一个给定的函数中
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
