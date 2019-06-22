package main

import (
	"encoding/json"
	"fmt"
)

//结构体和json转换的扩张 主要用于保护数据结构安全
type Student150201 struct {
	Name  string `json:"name"`    //给字段改名
	Sex   string `json:"-"`       //该字段不输出 (注意此处为中划线)
	Score int    `json:",string"` //修改字段的数据类型
}

func main15020101() {
	stu := Student150201{"tom", "男", 66}
	json, err := json.Marshal(stu)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(json))
	/*
		{"Name":"tom","Sex":"男","Score":66}
		修改后
		{"name":"tom","Score":"66"}
	*/
}
