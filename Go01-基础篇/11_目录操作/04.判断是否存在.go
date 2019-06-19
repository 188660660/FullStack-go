package main

import (
	"fmt"
	"os"
)

//Stat()用来获取文件的信息，如果报错，说明文件或文件夹不存在。
func main100401() {
	_, err := os.Stat("./user")
	if err != nil {
		fmt.Println(err)
	} else {
		err := os.Rename("./user", "./test")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("目录重命名成功！")
	}
	return
}
