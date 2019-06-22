package main

import (
	"encoding/json"
	"fmt"
)

//JSON(JavaScript Object Notation) 是一种轻量级的数据交换格式。格式的特点：键值对
type Student1501 struct {
	//类的首字母大写才能调用 代表着公共的
	Name string
	Sex  string
	Age  int
}
type Student1502 struct {
	//类的首字母大写才能调用 代表着公共的
	name string
	sex  string
	age  int
}

//将结构体转成json
func main150101() {
	stu := Student1501{"tom", "男", 88}
	jsons, err := json.Marshal(stu) //将结构体转成json
	if err != nil {
		fmt.Println(err)
		return
	}
	//转成json并且格式化
	jsons, _ = json.MarshalIndent(stu, "", " ")
	fmt.Println(string(jsons))
}

func main150102() {
	stu := Student1501{"小白", "女", 16}
	stuJson, _ := json.Marshal(stu)
	stuJson, _ = json.MarshalIndent(stu, "", "")
	fmt.Println(string(stuJson))
}

//2.将JSON转成结构体
func main150201() {
	str := `{"Name":"tom","Sex":"男","Score":88}`
	var stu Student1501
	err := json.Unmarshal([]byte(str), &stu) //将json转成结构体
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(stu)
}

//练习
func main150202() {
	str := `{"Name":"tom","Sex":"男","Age":88}`
	var stu Student1501
	err := json.Unmarshal([]byte(str), &stu)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(stu) //{tom 男 0}
	fmt.Println(stu.Name, stu.Sex, stu.Age)
}
