package main

import (
	"fmt"
	"reflect"
)

type Student1506 struct {
	name  string "学生姓名"
	sex   string "学生性别"
	score int    "学生成绩"
}

func main1506() {
	var stu = Student1506{"小李", "男", 88}
	t := reflect.TypeOf(stu)  // 反射对象的类型
	v := reflect.ValueOf(stu) // 反射对象的值
	fmt.Println(t)            // main.Student1506
	fmt.Println(v)            // {小李 男 88}
	fmt.Println(t.NumField()) //3 成员个数
	for i := 0; i < t.NumField(); i++ {
		fmt.Println("字段名：", t.Field(i).Name, "值：", v.Field(i), "备注：", t.Field(i).Tag, "数据类型：", t.Field(i).Type)
	}
}
