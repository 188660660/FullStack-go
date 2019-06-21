package main

import (
	"fmt"
	"regexp"
)

func main140201() {
	str := []string{"aba", "ABA", "aBA", "Aba", "ab\nAb", "abXXXXXXXXXXXb"}

	reg := regexp.MustCompile(`^a.*\w`) //区分大小写匹配 aba,aBA,ab,Ab,abXXXXXXXXXXXb
	//reg = regexp.MustCompile(`(?i:^a).*\w`) //不区分大小写匹配 aba,ABA,aBA,Aba,ab,Ab,abXXXXXXXXXXXb

	reg = regexp.MustCompile(`(?m:^a)\w+`)  //多行模式
	reg = regexp.MustCompile(`(?mi:^a)\w+`) //多行模式 且不区分大小写

	reg = regexp.MustCompile(`a.*b`)      //默认就是贪婪模式
	reg = regexp.MustCompile(`(?U)a.*b`)  //(?U)设置成非贪婪模式
	reg = regexp.MustCompile(`(?U:)a.*b`) //(?U)设置成非贪婪模式

	for _, v := range str {
		result := reg.FindAllString(v, -1)
		if result == nil {
			continue
		}
		fmt.Println(result)
		//if reg.MatchString(v){
		//	fmt.Println(v)
		//}
	}
}

func mainmain140202() {
	//?s: 忽略空白字符
	str := `<html>
	<div>
		锄禾日当午，
汗滴禾下土。
谁知盘中餐，
粒粒皆辛苦。
	</div>
</html>`

	reg := regexp.MustCompile(`<div>(.+)</div>`) //不忽略换行
	//reg = regexp.MustCompile(`<div>(?s:(.+))</div>`)	//忽略换行

	result := reg.FindAllString(str, -1)
	fmt.Println(result)
}
