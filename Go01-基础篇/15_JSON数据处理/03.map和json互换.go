package main

import (
	"encoding/json"
	"fmt"
)

//map转成json
func main150301() {
	m := make(map[string]interface{}, 4)
	m["name"] = "Tom"
	m["age"] = 25
	m["score"] = 88.5
	m["is_stu"] = true
	mJson, err := json.Marshal(m) //将map转成json
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(mJson)) //将json转成字符串输出
	//{"age":25,"is_stu":true,"name":"Tom","score":88.5}
}

//json转成map
func main150302() {
	str := `{"age":25,"is_stu":true,"name":"Tom","score":88.5}`
	m := make(map[string]interface{}, 4)
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(m)        //map[age:25 is_stu:true name:Tom score:88.5]
	fmt.Println(m["age"]) //25
}
