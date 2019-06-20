package main

func main130301() {
	/*
	      Format 系列函数把其他类型的转换为字符串。
	   	(1)Format:英文单词格式的意思
	   		1.布尔型转字符串
	   			语法：
	   				strconv.FormatBool(数据)

	   		2.整形型转字符串
	   			语法A：常用推荐
	   				strconv.Itoa(666)
	   					输出： "666" string类型

	   			语法B：
	   				strconv.FormatInt(数据,任意进制>=2)


	   		3.浮点型型转字符串
	   			语法：
	   				strconv.FormatFloat(a,b,c,d)
	   			参数说明：
	   				a : 指需要转换的数据
	   				b : 指打印格式，一般写作 'f' 以小数形式
	   				c : 指小数点位数
	   				d : 指使用float64 还是 float32进行处理
	   			示例：
	   				str := 3.14159
	   				str = strconv.FormatFloat(str,'f',2,64)
	   					输出：str : 3.14 类型 float64

	      Parse 系列函数把字符串转换为其他类型
	      	(2)Parse:英文单词解析的意思
	   		1.字符串转布尔类型
	      			语法：
	      				strconv.ParseBool("数据")

	      		2.字符串转整形
	      			语法A：常用推荐
	      				strconv.Atoi("666")
	      					输出： 666 int类型

	      			语法B：
	      				strconv.ParseInt(A,B,C)
	      					func ParseInt(s string, base int, bitSize int) (i int64, err error)
	   					   参数A 数字的字符串形式
	   					   参数B 数字字符串的进制 比如二进制 八进制 十进制 十六进制
	      					   参数C 返回结果的bit大小 也就是int8 int16 int32 int64

	      		3.字符串转浮点型
	   			语法：
	   				strconv.ParseFloat(A,B)
	   				   参数A 数字的字符串形式
	   				   参数B 返回结果的bit大小 也就是int8 int16 int32 int64

	      Append 系列函数将整数等转换为字符串后，添加到现有的字节数组中。
	      	(3)Append:英文单词附加的意思
	      		1.将布尔类型转成字符串追加到数组
	      			语法：
	      				strconv.AppendBool(追加的数组,追加的bool)

	   		2.将Int类型转成字符串追加到数组
	   		语法：
	   			strconv.AppendInt(追加的数组,追加数,指定追加的进制)
	      			示例：strconv.AppendInt(slice,123,10)

	      		2.将字符串转换为带引号的字符串文字。
	      		语法：
	      			strconv.AppendQuote(追加的数组,"字符串")
	         			示例：strconv.AppendQuote(slice,"你好啊")
	*/
}
