package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"log"
	"strconv"
)

func main() {
	path := `E:\Go_WorkSpace\local-project\project5-1\src\conf.ini`
	//加载配置文件
	cfg, err := goconfig.LoadConfigFile(path)
	//2.  加载并读取 conf.ini 使得命令行输出与 sample_output.txt 一致的内容
	if err != nil {
		log.Printf("配置文件加载失败：%s", err)
		return
	}
	/*
		命令行输出
		1.2014/02/01 10:48:07 作者：无闻
		2.2014/02/01 10:48:07 GitHub：https://github.com/Unknwon
		3.2014/02/01 10:48:07 视频教程列表：
		4.2014/02/01 10:48:07 #1：Go编程基础
		5.2014/02/01 10:48:07 #2：Go Web基础
		6.2014/02/01 10:48:07 #3：Go名库讲解
		7.2014/02/01 10:48:07 欢迎使用 goconfig！
	*/
	p1, err1 := cfg.GetValue("", "author")
	if err1 != nil {
		log.Printf("P1错误信息：%s", err1)
		return
	}
	log.Printf("作者:%s", p1)

	p2, err2 := cfg.GetValue("", "mygithub")
	if err2 != nil {
		log.Printf("P2错误信息：%s", err2)
		return
	}
	log.Printf("GitHub：%s", p2)

	p3 := cfg.GetSectionComments("courses")
	if p3 == "" {
		log.Println("P3注释信息不存在！")
		return
	}
	log.Printf("%s", p3)

	p4, err4 := cfg.GetSection("courses")
	if err4 != nil {
		log.Printf("P4获取失败：%s", err4)
		return
	}
	log.Printf("#1:%s", p4["#1"])
	log.Printf("#2:%s", p4["#2"])
	log.Printf("#3:%s", p4["#3"])

	p5, err5 := cfg.GetValue("dir.Go名库讲解.01-goconfig", "name")
	if err5 != nil {
		log.Printf("P5子孙节点获取失败：%s", err5)
		return
	}
	log.Printf("%s", p5)

	/*
		3.  操作配置文件对象并保存，保存结果需与 output.ini 一致
	*/
	//修改键注释 - 获取
	cm, err := cfg.GetSection("courses")
	if err != nil {
		log.Printf("获取courses的节点信息失败：%s", err)
		return
	}
	clen1 := "#" + strconv.Itoa(len(cm))
	comment := `https://github.com/Unknwon/go-rock-libraries-showcases`

	//此处需要注意！
	v := cfg.SetKeyComments("courses", clen1, comment)
	if v {
		log.Printf("super分区注释插入或删除成功！%v", v)
	} else {
		log.Printf("super分区注释重写成功！%v", v)
	}

	//保存配置文件
	err = goconfig.SaveConfigFile(cfg, "output2.ini")
	if err != nil {
		log.Printf("数据保存失败：%s", err)
		return
	} else {
		fmt.Println("数据保存成功！")
	}
}
