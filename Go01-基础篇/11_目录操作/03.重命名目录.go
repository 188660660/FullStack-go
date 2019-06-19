package main

import (
	"fmt"
	"os"
)

/*
	重命名目录：Rename(老的路径，新的路径)
*/
func main100301() {
	err := os.Rename("./aa", "./user")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("文件重命名成功！")
}
