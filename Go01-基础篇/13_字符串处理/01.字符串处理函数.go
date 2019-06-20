package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main130101() {
	//value := strings.Contains("abc","ab") //true | 字符串中是否包含子字符串 返回布尔类型

	//value := strings.Join([]string{"a","b","c"},"-") //a-b-c | 将切片的元素用什么链接起来
	//value := strings.Split("a-b-c","-")//[a b c] | 将字符串分割成切片

	//value := strings.Index("abcde","cd")//2 | 返回子字符串第一次出现的位置 如果找不到就返-1
	//value := strings.LastIndex("abcde","csss")//2 | 返回子字符串最后一次出现的位置 如果找不到就返-1

	//value := strings.Repeat("ab",3) //ababab | 将字符串重复几次
	//value := strings.Replace("abcdfbdbdwwb","b","*",-1)//a*cdf*d*dww* | 将字符串替换 param：需要替换的字符串 替换的能容 被什么替换 为0<时表示全部替换

	//value := strings.Trim("==你好吗==","=") //你好吗 | 去除掉指定的界定符
	//value := strings.Fields("a asa   ass sdww s sdsd") //[a asa ass sdww s sdsd] | 按空格将字符串转成切片

	//value := strings.EqualFold("GO","go") //true | 忽略大小写进行比较 返回布尔型

	//value := strings.HasPrefix("abc.jpg.jsss","abc.")	//true | 判断是否具有指定的前缀 返回布尔型
	//value := strings.HasSuffix("abc.jpg.jsss",".jsss") //true | 判断是否具有指定的后缀 返回布尔型

	//fmt.Println(value)
}

func mainCount() {
	/*
		语法：
			func Count(s , sep string) int
				Count 计算字符串 sep 在s中的非重叠个数
				如果sep为空字符串 则返回S中的字符(非字节)个数+1
				如果使用Rabin-Karp算法实现
	*/
	s := "Hello World!!!!!"
	n := strings.Count(s, "!") //5
	n = strings.Count(s, "!!") //2
	n = strings.Count(s, "")   //17 16+1
	fmt.Println(n)
}

func mainContains() {
	/*
		语法：
			func Contains(s , substr string) bool
				1.Contains判断字符串s 中是否包含子串substr
				2.如果substr为空 则返回true
	*/
	s := "HelloWorld!!!!"
	b := strings.Contains(s, "W") //true
	b = strings.Contains(s, "D")  //false
	b = strings.Contains(s, "")   //true !
	b = strings.Contains(s, " ")  //false
	fmt.Println(b)
}

func mainContainsAny() {
	/*
		语法：
			func ContainAny(s , chars string) bool
			1.ContainAny判断字符串s中是否包含chars中的任何一个字符
			2.如果chars为空 就返回false
	*/
	s := "Hello,中国"
	b := strings.ContainsAny(s, "abc")     //false
	b = strings.ContainsAny(s, "def")      //true
	b = strings.ContainsAny(s, "传说中的暴龙兽！") //true
	b = strings.ContainsAny(s, " ")        //false
	b = strings.ContainsAny(s, "")         //false
	fmt.Println(b)
}

func mainContainsRune() {
	/*
		语法：
			func ContainsRune(s string,r rune) bool
			1.ContainsRune 判断字符串 s 中是否包含字符 r
	*/
	s := "HelloWorld!!! 中国"
	b := strings.ContainsRune(s, '\n') //false
	b = strings.ContainsRune(s, '国')   //true
	b = strings.ContainsRune(s, 0)     //false
	b = strings.ContainsRune(s, ' ')   //true
	fmt.Println(b)
}

func mainIndex() {
	/*
		语法：
			func Index(s ,sep string) int
			1.Index返回子串sep在字符串sz中第一次出现的位置
			2.如果找不到 返回-1 , 如果sep 为空 ,这返回 0
			3.使用Rabin-Karp算法实现
	*/
	s := "Hello,中国！"
	i := strings.Index(s, "h") //-1
	i = strings.Index(s, "H")  //0
	i = strings.Index(s, "！")  //12
	i = strings.Index(s, "")   //0
	i = strings.Index(s, " ")  //-1
	fmt.Println(i)
}

func mainLastIndex() {
	/*
		语法：
			func LastIndex(s ,sep string) int
			1.LastIndex返回子串sep在字符串s中最后一次出现的位置
			2.如果找不到，则返回-1,如果sep为空,则返回字符串的长度
			3.使用朴素字符串比较算法实现
	*/
	s := "Hello,世界！Hello"
	i := strings.LastIndex(s, "h") // -1
	i = strings.LastIndex(s, "")   // 20
	i = strings.LastIndex(s, "H")  // 15
	fmt.Println(i)
}

func mainIndexRune() {
	/*
		语法：
			func IndexRune(s string , r rune) int
			1.IndexRune返回字符r 在字符串s中第一次出现的位置
			2.如果找不到，则返回-1
	*/
	s := "Hello,世界！Hello"
	i := strings.IndexRune(s, '\n') //-1
	i = strings.IndexRune(s, 'H')   //0
	i = strings.IndexRune(s, 0)     //-1
	fmt.Println(i)
}

func mainIndexAny() {
	/*
		语法：
			func(s ,chars string) int
			1.IndexAny 返回字符串chars中的任何一个字符在字符串s中第一次出现的位置
			2.如果找不到 这返回-1 如果chars为空 这返回-1
	*/
	s := "Hello,世界！ Hello"
	i := strings.IndexAny(s, "abcde") //1
	i = strings.IndexAny(s, "dof")    //4
	i = strings.IndexAny(s, "f")      //-1
	i = strings.IndexAny(s, "")       //-1
	fmt.Println(i)
}

func mainLastIndexAny() {
	/*
		语法：
			func LastIndexAny(s ,chars string) int
			1.LastIndexAny 返回字符串chars中的任何一个字符在字符串s 中最后一次出现的位置
			2.如果找不到 返回-1 如果chars为空 也返回-1
	*/
	s := "Hello,世界！ Hello"
	i := strings.LastIndexAny(s, "abv") //-1
	i = strings.LastIndexAny(s, "")     //-1
	i = strings.LastIndexAny(s, " ")    //15
	i = strings.LastIndexAny(s, "cde")  //17
	fmt.Println(i)
}

/******************用法不明******************/
func mainSplitN() {
	/*
		语法：fun SplitN(s , sep string , n int)[]string
		1.如果sep为空,则将s切分成Unicode字符列表
		2.如果s中没有sep子串,则整个s作为[]string的第一个元素返回
		3.参数最多且分出几个子串,超出的部分将不再切分
		4.如果n为0 ,这返回 nil , 如果n 小于0 这不限制切分个数 ,全部切分
	*/
	//%q单引号围绕的字符字面值，由Go语法安全地转义 Printf("%q", 0x4E2D)

	s := "Hello,世界！ Hello"
	ss := strings.SplitN(s, "", 2)
	fmt.Printf("%q\n", ss) //["H" "ello,世界！ Hello"]

	ss = strings.SplitN(s, " ", 2)
	fmt.Printf("%q\n", ss) //["Hello,世界！" "Hello"]

	ss = strings.SplitN(s, " ", -1)
	fmt.Printf("%q\n", ss) //["Hello,世界！" "Hello"]

	ss = strings.SplitN(s, " ", 3)
	fmt.Printf("%q\n", ss) //["Hello,世界！" "Hello"]

	ss = strings.SplitN(s, "", 3)
	fmt.Printf("%q\n", ss) //["H" "e" "llo,世界！ Hello"]
}

func mainSplitAfterN() {
	/*
		语法：
		func SplitAfterN(s , sep string , n int)[]string
		1.SplitAfterN以sep为分割符,将s切分成多个子串,结果中包含sep本身
		2.如果sep为空,则将s切分成Unicode字符列表
		3.如果s中没有sep子串，则将整个s作为[]string的第一个元素返回
		4.参数n表示最多切分出几个子串,超出的部分将不再切分
		5.如果n为0,则返回 nil ,如果n小于0，则不限制切分个数，全部切分
	*/
	s := "Hello, 世界! Hello!"
	ss := strings.SplitAfterN(s, " ", 2)
	fmt.Printf("%q\n", ss) //["Hello, " "世界! Hello!"]

	ss = strings.SplitAfterN(s, "", 2)
	fmt.Printf("%q\n", ss) //["H" "ello, 世界! Hello!"]

	ss = strings.SplitAfterN(s, " ", -1)
	fmt.Printf("%q\n", ss) //["Hello, " "世界! " "Hello!"]

	ss = strings.SplitAfterN(s, "", 3)
	fmt.Printf("%q\n", ss) //["H" "e" "llo, 世界! Hello!"]
}

func mainSplit() {
	/*
		语法：
			func Split(s ,sep string)[]string
			1.Split以sep为分割符,将s切片分成多个子切片,结果中不包含sep本身
			2.如果s中没有sep子串,则将s切分成Unicode字符列表
			3.如果s中没有sep子串,则将整个s作为[]string的第一个元素返回
	*/
	s := "Hello, 世界! Hello!"
	ss := strings.Split(s, "")
	fmt.Printf("%q\n", ss) //单个字符列表 ["H" "e" "l" "l" "o" "," " " "世" "界" "!" " " "H" "e" "l" "l" "o" "!"]

	ss = strings.Split(s, " ")
	fmt.Printf("%q\n", ss) //["Hello," "世界!" "Hello!"]

	ss = strings.Split(s, ",")
	fmt.Printf("%q\n", ss) //["Hello" " 世界! Hello!"]
}

func mainSplitAfter() {
	/*
		语法：
			func SplitAfter(s ,sep string)[]string
			1.SplitAfter以sep为分隔符,将s切分成多个切片,结果中包含sep本身
			2.如果sep为空,这将s切分成Unicode字符列表
			3.如果s中没有sep子串,则将整个s作为[]string的第一个元素返回
	*/
	s := "Hello, 世界! Hello!"
	ss := strings.SplitAfter(s, " ")
	fmt.Printf("%q\n", ss) //["Hello, " "世界! " "Hello!"]

	ss = strings.SplitAfter(s, ",")
	fmt.Printf("%q\n", ss) //["Hello," " 世界! Hello!"]

	ss = strings.SplitAfter(s, "")
	fmt.Printf("%q\n", ss) //单个字符列表 ["H" "e" "l" "l" "o" "," " " "世" "界" "!" " " "H" "e" "l" "l" "o" "!"]
}

/******************TOPENDEOF******************/

func mainFields() {
	/*
		语法：
			func Fields(s string)[]string
			1.Fields以连续的空白字符 为分割符,将s切分成多个子串,结果中不包含空白字符本身
			2.空白字符有:
				\t,\n,\v,\r\f,'',U+0085(NEL),U+00A0(NBSP)
			3.如果s中只包含空白字符,这返回一个空列表
	*/
	s := "Hello, 世界! Hello!"
	ss := strings.Fields(s)

	fmt.Printf("%q\n", ss) //["Hello," "世界!" "Hello!"]
	fmt.Println(ss)        //[Hello, 世界! Hello!]
}

/*************start*************/
func isSlash(r rune) bool {
	return r == '\\' || r == '/'
}

func mainFieldFunc() {
	/*
		语法：
			func FieldFunc(s string,f func(rune) bool)[]string
			1.FieldsFunc 以一个或多个满足f(rune)的字符为分隔符
			2.将s切分成多个子串,结果中不包含分隔符本身。
	*/
	s := "C:\\Windows\\System32\\FileName"
	ss := strings.FieldsFunc(s, isSlash)
	fmt.Printf("%q\n", ss) //["C:" "Windows" "System32" "FileName"]
}

/**************end**************/

func mainJoin() {
	/*
		语法：
			func Join(a []string,sep string)string
			1.Join将a中的子串连接成一个单独的字符串，子串之间用sep分隔
			2.可以使用空格进行连接
	*/
	ss := []string{"A", "B", "C"}
	s := strings.Join(ss, "") //ABC
	s = strings.Join(ss, "|") //A|B|C
	fmt.Println(s)
}

func mainHasPrefix() {
	/*
		语法：
			func HasPrefix(s,prefix string)bool
			1.HasPrefix判断字符串 s是否以prefix开头
	*/
	s := "Hello 世界!"
	b := strings.HasPrefix(s, "x") //false
	b = strings.HasPrefix(s, "He") //true
	fmt.Println(b)
}

func mainHasSuffix() {
	/*
		语法：
			func HasSuffix(s,suffix string)bool
			1.判断字符串s以prefix结尾
	*/
	s := "Hello 世界!"
	b := strings.HasSuffix(s, "Hello") //false
	b = strings.HasSuffix(s, "")       //为空字符串的时候为true
	b = strings.HasSuffix(s, " ")      //false
	b = strings.HasSuffix(s, "!")      //true
	fmt.Println(b)
}

/*************start*************/
func Slash(r rune) rune {
	if r == '\\' {
		return '/'
	}
	return r
}

func mainMap() {
	/*
		语法：
			func Map(mapping func(rune)rune,s string)string
			1.Map：将s 中满足mapping(rune)的字符替换为mapping(rune)的返回值
			2.如果mapping(rune) 返回负数,则相应的字符将被删除
	*/
	s := "C:\\Windows\\System32\\FileName"
	ms := strings.Map(Slash, s)
	fmt.Printf("%q\n", ms) //"C:/Windows/System32/FileName"
}

/**************end**************/

func mainRepeat() {
	/*
		语法：
		func Repeat(s string,count int)string //repeat:英语重复的意思
		1.Repeat将count个字符串s连接成一个新的字符串
	*/
	s := "Hello"
	rs := strings.Repeat(s, 3)
	fmt.Printf("%q\n", rs) //"HelloHelloHello"
	fmt.Println(rs)        //HelloHelloHello
}

/*
	语法A：
		func ToUpper(s string) string
		1.ToUpper将s中的所有字符修改其为大写
		2.对于非ASCII字符，它的大写格式需要查表转换
	语法B：
		func ToLower(s string)string
		1.ToLower将s中的所有字符修改为小写格式
		2.对于非ASCII字符，它的小写格式需要查表转换
	语法C：
		func ToTitle(s string) string
		1.ToTitle将s中的所有字符修改为Title格式
		2.大部分字符的Title格式是Upper格式
		3.只有少数字符的Title格式是特殊字符
		4.ToTitle主要给Title函数调用
*/
func mainConversionSet() {
	s := "heLLo worLd Ａｂｃ"
	us := strings.ToUpper(s)
	ls := strings.ToLower(s)
	ts := strings.ToTitle(s)

	fmt.Println(us)        //HELLO WORLD ＡＢＣ
	fmt.Printf("%q\n", us) //"HELLO WORLD ＡＢＣ"

	fmt.Println(ls)        //hello world ａｂｃ
	fmt.Printf("%q\n", ls) //"hello world ａｂｃ"

	fmt.Println(ts)        //HELLO WORLD ＡＢＣ
	fmt.Printf("%q\n", ts) //"HELLO WORLD ＡＢＣ"
}

//获取非ASCII字符的Title格式列表
func mainToTitleUnAscii() {
	for _, cr := range unicode.CaseRanges {
		// u := uint32(cr.Delta[unicode.UpperCase]) // 大写格式
		// l := uint32(cr.Delta[unicode.LowerCase]) // 小写格式
		t := uint32(cr.Delta[unicode.TitleCase]) // Title 格式
		// if t != 0 && t != u {
		if t != 0 {
			for i := cr.Lo; i <= cr.Hi; i++ {
				fmt.Printf("%c -> %c\n", i, i+t)
			}
		}
	}
}

/*
	语法A：
		func ToUpperSpecial(_case unicode.SpecialCase, s string) string
		1.ToUpperSpecial 将s中的所有字符修改为大写格式
		2.优先使用_case中的规则进行转换
	语法B：
		func ToLowerSpecial(_case unicode.SpecialCase, s string) string
		1.ToLowerSpecial 将 s 中的所有字符修改为小写格式。
		2.优先使用_case中的规则进行转换
	语法C：
		func ToTitleSpecial(_case unicode.SpecialCase, s string) string
 		1.ToTitleSpecial 将 s 中的所有字符修改为其 Title 格式。
		2.优先使用_case中的规则进行转换

_case 规则说明举例：
	unicode.CaseRange{'A', 'Z', [unicode.MaxCase]rune{3, -3, 0}}
	·其中 'A', 'Z' 表示此规则只影响 'A' 到 'Z' 之间的字符。
	·其中 [unicode.MaxCase]rune 数组表示：
	当使用 ToUpperSpecial 转换时，将字符的 Unicode 编码与第一个元素值（3）相加
	当使用 ToLowerSpecial 转换时，将字符的 Unicode 编码与第二个元素值（-3）相加
	当使用 ToTitleSpecial 转换时，将字符的 Unicode 编码与第三个元素值（0）相加
*/

func mainSpecial() {
	// 定义转换规则
	var _MyCase = unicode.SpecialCase{
		// 将半角逗号替换为全角逗号，ToTitle 不处理
		unicode.CaseRange{',', ',',
			[unicode.MaxCase]rune{'，' - ',', '，' - ',', 0}},
		// 将半角句号替换为全角句号，ToTitle 不处理
		unicode.CaseRange{'.', '.',
			[unicode.MaxCase]rune{'。' - '.', '。' - '.', 0}},
		// 将 ABC 分别替换为全角的 ＡＢＣ、ａｂｃ，ToTitle 不处理
		unicode.CaseRange{'A', 'C',
			[unicode.MaxCase]rune{'Ａ' - 'A', 'ａ' - 'A', 0}},
	}
	s := "ABCDEF,abcdef."
	us := strings.ToUpperSpecial(_MyCase, s)
	fmt.Printf("%q\n", us) // "ＡＢＣDEF，ABCDEF。"
	ls := strings.ToLowerSpecial(_MyCase, s)
	fmt.Printf("%q\n", ls) // "ａｂｃdef，abcdef。"
	ts := strings.ToTitleSpecial(_MyCase, s)
	fmt.Printf("%q\n", ts) // "ABCDEF,ABCDEF."
}

func mainTitle() {
	/*
		语法：
			func Title(s string) string
			1.Title 将 s 中的所有单词的首字母修改为其 Title 格式
			2.BUG: Title 规则不能正确处理 Unicode 标点符号
	*/
	s := "heLLo worLd"
	ts := strings.Title(s)
	fmt.Printf("%q\n", ts) // "HeLLo WorLd"
}

func isSlashTrim(r rune) bool {
	return r == '\\' || r == '/'
}
func mainTrimRightFunc() {
	/*
		语法：
			func TrimRightFunc(s string, f func(rune) bool) string
			1.TrimRightFunc 将删除 s 尾部连续的满足 f(rune) 的字符
	*/
	s := "\\\\HostName\\C\\Windows\\"
	ts := strings.TrimRightFunc(s, isSlashTrim)
	fmt.Printf("%q\n", ts) // "\\\\HostName\\C\\Windows"
}

func mainTrimFunc() {
	/*
		语法：
		func TrimFunc(s string, f func(rune) bool) string
		1.TrimFunc 将删除 s 首尾连续的满足 f(rune) 的字符
	*/
	s := "\\\\HostName\\C\\Windows\\"
	ts := strings.TrimFunc(s, isSlashTrim)
	fmt.Printf("%q\n", ts) // "HostName\\C\\Windows"
}

func mainIndexFunc() {
	/*
		语法：
			func IndexFunc(s string, f func(rune) bool) int
		1.返回 s 中第一个满足 f(rune) 的字符的字节位置。
		2.如果没有满足 f(rune) 的字符，则返回 -1
	*/
	s := "C:\\Windows\\System32"
	i := strings.IndexFunc(s, isSlashTrim)
	fmt.Printf("%v\n", i) // 2
}

func mainLastIndexFunc() {
	/*
		语法：
			func LastIndexFunc(s string, f func(rune) bool) int
		1.返回 s 中最后一个满足 f(rune) 的字符的字节位置。
		2.如果没有满足 f(rune) 的字符，则返回 -1
	*/
	s := "C:\\Windows\\System32"
	i := strings.LastIndexFunc(s, isSlashTrim)
	fmt.Printf("%v\n", i) // 10
}

func mainTrim() {
	/*
		语法：
			func Trim(s string, cutset string) string
			1.Trim 将删除 s 首尾连续的包含在 cutset 中的字符
	*/
	s := " Hello 世界! "
	ts := strings.Trim(s, " Helo!")
	fmt.Printf("%q\n", ts) // "世界"
}

func mainTrimLeft() {
	/*
		语法：
			func TrimLeft(s string, cutset string) string
		 	1.TrimLeft 将删除 s 头部连续的包含在 cutset 中的字符
	*/
	s := " Hello 世界! "
	ts := strings.TrimLeft(s, " Helo")
	fmt.Printf("%q\n", ts) // "世界! "
}

func mainTrimRight() {
	/*
		语法:
			func TrimRight(s string, cutset string) string
			1.TrimRight 将删除 s 尾部连续的包含在 cutset 中的字符
	*/
	s := " Hello 世界! "
	ts := strings.TrimRight(s, " 世界!")
	fmt.Printf("%q\n", ts) // " Hello"
}

func mainTrimSpace() {
	/*
		语法：
			func TrimSpace(s string) string
			1.TrimSpace 将删除 s 首尾连续的的空白字符
	*/
	s := " H ello 世 界! "
	ts := strings.TrimSpace(s)
	fmt.Printf("%q\n", ts) // "H ello 世 界!"
}

func mainTrimPrefix() {
	/*
		语法：
			func TrimPrefix(s, prefix string) string
			1.TrimPrefix 删除 s 头部的 prefix 字符串
			2.如果 s 不是以 prefix 开头，则返回原始 s
	*/
	s := "Hello 世界!"
	ts := strings.TrimPrefix(s, "Hello")
	fmt.Printf("%q\n", ts) // " 世界"
}

func mainTrimSuffi() {
	/*
		语法：
			func TrimSuffix(s, suffix string) string
			1.TrimSuffix 删除 s 尾部的 suffix 字符串
			2. 如果 s 不是以 suffix 结尾，则返回原始 s
	*/
	s := "Hello 世界!!!!!"
	ts := strings.TrimSuffix(s, "!!!!")
	fmt.Printf("%q\n", ts) // "Hello 世界!"
	/*
		注：TrimSuffix只是去掉s字符串结尾的suffix字符串，只是去掉１次，
			而TrimRight是一直去掉s字符串右边的字符串，
			只要有响应的字符串就去掉，是一个多次的过程，这也是二者的本质区别．
	*/
}

func mainReplace() {
	/*
		语法：
			func Replace(s, old, new string, n int) string
			1.Replace 返回 s 的副本，并将副本中的 old 字符串替换为 new 字符串
			2.替换次数为 n 次，如果 n 为 -1，则全部替换
			3.如果 old 为空，则在副本的每个字符之间都插入一个 new
	*/
	s := "Hello 世界！"
	s = strings.Replace(s, " ", ",", -1)
	fmt.Println(s)
	s = strings.Replace(s, "", "|", -1)
	fmt.Println(s)
}

func mainEqualFold() {
	/*
			语法：
				func EqualFold(s, t string) bool
		 	1.EqualFold 判断 s 和 t 是否相等。忽略大小写，同时它还会对特殊字符进行转换
			2.比如将“ϕ”转换为“Φ”、将“Ǆ”转换为“ǅ”等，然后再进行比较
	*/
	s1 := "Hello 世界! ϕ Ǆ"
	s2 := "hello 世界! Φ ǅ"
	b := strings.EqualFold(s1, s2)
	fmt.Printf("%v\n", b) // true
}

//// reader.go

func mainNewReader() {
	/*
		语法：
		// Reader 结构通过读取字符串，实现了 io.Reader，io.ReaderAt，
		// io.Seeker，io.WriterTo，io.ByteScanner，io.RuneScanner 接口
		type Reader struct {
			s string // 要读取的字符串
			i int // 当前读取的索引位置，从 i 处开始读取数据
			prevRune int // 读取的前一个字符的索引位置，小于 0 表示之前未读取字符
		}

		// 通过字符串 s 创建 strings.Reader 对象
		// 这个函数类似于 bytes.NewBufferString
		// 但比 bytes.NewBufferString 更有效率，而且只读
			func NewReader(s string) *Reader { return &Reader{s, 0, -1} }
	*/
}

func mainLen() {
	/*
		语法：
			func (r *Reader) Len() int
			1.Len 返回 r.i 之后的所有数据的字节长度
	*/
	s := "Hello 世界!"
	// 创建 Reader
	r := strings.NewReader(s)
	// 获取字符串的编码长度
	fmt.Println(r.Len()) // 13
}

func mainReader() {
	/*
		语法：
			func (r *Reader) Read(b []byte) (n int, err error)
			1. Read 将 r.i 之后的所有数据写入到 b 中（如果 b 的容量足够大）
			2. 返回读取的字节数和读取过程中遇到的错误
			3. 如果无可读数据，则返回 io.EOF
	*/
	s := "Hello World!"
	// 创建 Reader
	r := strings.NewReader(s)
	// 创建长度为 5 个字节的缓冲区
	b := make([]byte, 5)
	// 循环读取 r 中的字符串
	for n, _ := r.Read(b); n > 0; n, _ = r.Read(b) {
		fmt.Printf("%q, ", b[:n]) // "Hello", " Worl", "d!"
	}
}

func mainReadAt() {
	/*
		语法：
			func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
		 1. ReadAt 将 off 之后的所有数据写入到 b 中（如果 b 的容量足够大）
		 2. 返回读取的字节数和读取过程中遇到的错误
		 3. 如果无可读数据，则返回 io.EOF
		 4. 果数据被一次性读取完毕，则返回 io.EOF
	*/
	s := "Hello World!"
	// 创建 Reader
	r := strings.NewReader(s)
	// 创建长度为 5 个字节的缓冲区
	b := make([]byte, 5)
	// 读取 r 中指定位置的字符串
	n, _ := r.ReadAt(b, 0)
	fmt.Printf("%q\n", b[:n]) // "Hello"
	// 读取 r 中指定位置的字符串
	n, _ = r.ReadAt(b, 6)
	fmt.Printf("%q\n", b[:n]) // "World"
}

func mainReadByte() {
	/*
		语法：
			func (r *Reader) ReadByte() (b byte, err error)
			1. ReadByte 将 r.i 之后的一个字节写入到返回值 b 中
			2. 返回读取的字节和读取过程中遇到的错误
			3. 如果无可读数据，则返回 io.EOF
	*/
	s := "Hello World!"
	// 创建 Reader
	r := strings.NewReader(s)
	// 读取 r 中的一个字节
	for i := 0; i < 3; i++ {
		b, _ := r.ReadByte()
		fmt.Printf("%q, ", b) // 'H', 'e', 'l',
	}
}

func mainUnreadByte() {
	/*
		语法：
			func (r *Reader) UnreadByte() error
			1. UnreadByte 撤消前一次的 ReadByte 操作，即 r.i--
	*/
	s := "Hello World!"
	// 创建 Reader
	r := strings.NewReader(s)
	// 读取 r 中的一个字节
	for i := 0; i < 3; i++ {
		b, _ := r.ReadByte()
		fmt.Printf("%q, ", b) // 'H', 'H', 'H',
		_ = r.UnreadByte()    // 撤消前一次的字节读取操作
	}
}

func mainReadRune() {
	/*
		语法：
			func (r *Reader) ReadRune() (ch rune, size int, err error)
			1. ReadRune 将 r.i 之后的一个字符写入到返回值 ch 中
			2. ch： 读取的字符
			3. size：ch 的编码长度
			4. err： 读取过程中遇到的错误
			5. 如果无可读数据，则返回 io.EOF
			6. 如果 r.i 之后不是一个合法的 UTF-8 字符编码，则返回 utf8.RuneError 字符
	*/
	s := "你好 世界！"
	// 创建 Reader
	r := strings.NewReader(s)
	// 读取 r 中的一个字符
	for i := 0; i < 5; i++ {
		b, n, _ := r.ReadRune()
		fmt.Printf(`"%c:%v", `, b, n)
		// "你:3", "好:3", " :1", "世:3", "界:3",
	}
}

func mainUnreadRune() {
	/*
		语法：
			func (r *Reader) UnreadRune() error
			1. 撤消前一次的 ReadRune 操作
	*/

	s := "你好 世界！"
	// 创建 Reader
	r := strings.NewReader(s)
	// 读取 r 中的一个字符
	for i := 0; i < 5; i++ {
		b, _, _ := r.ReadRune()
		fmt.Printf("%q, ", b)
		// '你', '你', '你', '你', '你',
		_ = r.UnreadRune() // 撤消前一次的字符读取操作
	}
}

func mainSeek() {
	/*
		语法：
			func (r *Reader) Seek(offset int64, whence int) (int64, error)
			1. Seek 用来移动 r 中的索引位置
			2. offset：要移动的偏移量，负数表示反向移动
			3. whence：从那里开始移动，0：起始位置，1：当前位置，2：结尾位置
			4. 如果 whence 不是 0、1、2，则返回错误信息
			5. 如果目标索引位置超出字符串范围，则返回错误信息
			6. 目标索引位置不能超出 1 << 31，否则返回错误信息
	*/
	s := "Hello World!"
	// 创建 Reader
	r := strings.NewReader(s)
	// 创建读取缓冲区
	b := make([]byte, 5)
	// 读取 r 中指定位置的内容
	_, _ = r.Seek(6, 0)   // 移动索引位置到第 7 个字节
	_, _ = r.Read(b)      // 开始读取
	fmt.Printf("%q\n", b) //"World"
	_, _ = r.Seek(-5, 1)  // 将索引位置移回去
	_, _ = r.Read(b)      // 继续读取
	fmt.Printf("%q\n", b) //"World"
}

func mainNewBuffer() {
	/*
		语法：
			func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
			1. WriteTo 将 r.i 之后的数据写入接口 w 中
	*/
	s := "Hello World!"
	// 创建 Reader
	r := strings.NewReader(s)
	// 创建 bytes.Buffer 对象，它实现了 io.Writer 接口
	buf := bytes.NewBuffer(nil)
	// 将 r 中的数据写入 buf 中
	_, _ = r.WriteTo(buf)
	fmt.Printf("%q\n", buf) // "Hello World!"
}

// replace.go

func mainNewReplacer() {
	/*
		语法：
		// Replacer 根据一个替换列表执行替换操作
		type Replacer struct {
			Replace(s string) string
			WriteString(w io.Writer, s string) (n int, err error)
		}

		func NewReplacer(oldnew ...string) *Replacer
		1. NewReplacer 通过“替换列表”创建一个 Replacer 对象。
		2. 按照“替换列表”中的顺序进行替换，只替换非重叠部分。
		3. 如果参数的个数不是偶数，则抛出异常。
		4. 如果在“替换列表”中有相同的“查找项”，则后面重复的“查找项”会被忽略
	*/
}

func mainReplaceA() {
	/*
		语法：
			func (r *Replacer) Replace(s string) string
			1. Replace 返回对 s 进行“查找和替换”后的结果
			2. Replace 使用的是 Boyer-Moore 算法，速度很快
	*/
	srp := strings.NewReplacer("Hello", "你好", "World", "世界", "!", "！")
	s := "Hello World!Hello World!hello world!"
	rst := srp.Replace(s)
	fmt.Print(rst) // 你好 世界！你好 世界！hello world！
}

// ↕ 注：这两种写法均可
func mainReplaceB() {
	wl := []string{"Hello", "Hi", "Hello", "你好"}
	srp := strings.NewReplacer(wl...)
	s := "Hello World! Hello World! hello world!"
	rst := srp.Replace(s)
	fmt.Print(rst) // Hi World! Hi World! hello world!
}

func mainWriteString() {
	/*
		语法：
			func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error)
			1.WriteString 对 s 进行“查找和替换”，然后将结果写入 w 中
	*/
	wl := []string{"Hello", "你好", "World", "世界", "!", "！"}
	srp := strings.NewReplacer(wl...)
	s := "Hello World!Hello World!hello world!"
	_, _ = srp.WriteString(os.Stdout, s)
	// 你好 世界！你好 世界！hello world！
}
