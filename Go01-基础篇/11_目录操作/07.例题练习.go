package main

import (
	"fmt"
	"os"
)

/*
	1、判断C:\是否有user文件夹，如果没有就创建teacher文件夹，如果有就改名为teacher
	2、如果c:\有stu文件夹就删除，没有就创建
*/
const path = `E:\Go_WorkSpace\FullStack-go\Go01-基础篇\11_目录操作`

func main110701() {
	fp, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	f, err := fp.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, f := range f {
		if f.IsDir() {
			if f.Name() == "Test1203" {
				err := os.Rename(path+"\\Test1203", path+"\\teacher")
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("文件修改成功！")
			} else {
				CreateDir(path, "\\Test1203")
			}
			return
		}
	}
	CreateDir(path, "\\Test1203")
}

func CreateDir(path string, filepath string) {
	err := os.Mkdir(path+filepath, 777)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("目录创建成功！")
}

//2、如果c:\有stu文件夹就删除，没有就创建
func main110702() {
	_, err := os.Stat(path + "\\stu")
	if err != nil {
		CreateDir(path, "\\stu")
	} else {
		err = os.Remove(path + "\\stu")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
