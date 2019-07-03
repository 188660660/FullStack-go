package main

import (
	"fmt"
	"log"

	"github.com/Unknwon/goconfig"
)

func main() {
	// 加载配置文件
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		log.Fatalf("无法加载配置文件：%v", err)
	}

	// 操作键值对
	log.Printf("作者：%s", cfg.MustValue("", "author"))
	log.Printf("GitHub：%s", cfg.MustValue("", "mygithub"))

	log.Println("视频教程列表：")
	for i := 1; i <= 3; i++ {
		log.Printf("#%d：%s", i, cfg.MustValue("courses",
			fmt.Sprintf("#%d", i)))
	}

	log.Println(cfg.MustValue("dir.Go名库讲解.01-goconfig", "name"))

	cfg.SetKeyComments("courses", "#3", "https://github.com/Unknwon/go-rock-libraries-showcases")
	if err := goconfig.SaveConfigFile(cfg, "output.ini"); err != nil {
		log.Fatalf("无法保存配置文件：%v", err)
	}
}
