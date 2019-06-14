package main

//定义学生数据结构
type student080801 struct {
	name   string
	age    int
	sex    string
	cscore int
	mscore int
	escore int
}

//为对象赋值
func (s *student080801) InitInfo(name string, age int, sex string, cscore int, mscore int, escore int) {
	s.name = name
	s.age = age
	s.sex = sex
	s.cscore = cscore
	s.mscore = mscore
	s.escore = escore
}
func main080801() {
	var stu student080801
	//初始化对象信息
	stu.InitInfo("小白", 15, "女", 100, 200, 300)
	//拼接语句
	SayHellow080901()
	//打印输出
	stu.PrintScore080901()
}
