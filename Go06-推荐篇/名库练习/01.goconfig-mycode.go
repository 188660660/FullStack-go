package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"log"
	"strconv"
)

var path = `E:\Go_WorkSpace\local-project\project5-1\src\config1.ini`
var path2 = `E:\Go_WorkSpace\local-project\project5-1\src\config2.ini`
var path3 = `E:\Go_WorkSpace\local-project\project5-1\src\config3.ini`

func main01() {
	//加载配置文件
	cfg, err := goconfig.LoadConfigFile(path3)
	if err != nil {
		fmt.Println(err)
		return
	}

	//操作键值对
	log.Printf("search：%s", cfg.MustValue("", "search"))
	log.Printf("Githup:%s", cfg.MustValue("", "mygithup")) //没有为空
	//fmt.Println(cfg)
	value, err := cfg.GetValue(goconfig.DEFAULT_SECTION, "google")
	isInsert := cfg.SetValue(goconfig.DEFAULT_SECTION, "testes", "这是新的值")
	//isInsert 为true为写入成功 false表示原有值 覆盖写入
	if err != nil {
		fmt.Println("新值设置失败！")
		return
	}
	if isInsert {
		fmt.Println("写入成功！")
	} else {
		fmt.Println("覆盖写入成功！")
	}
	fmt.Printf(value)
	//默认会生成一个txt文本的文件
	err = goconfig.SaveConfigFile(cfg, "22")
	if err != nil {
		fmt.Println(err)
		return
	}
}

//基本使用方法
func main02() {
	cfg, err := goconfig.LoadConfigFile(path)
	if err != nil {
		fmt.Println("配置文件加载失败！")
		return
	}
	//获取默认值 如果没有值返回 就会返回一个"" DEFAULT_SECTION表示默认分区
	getinfo1, err := cfg.GetValue(goconfig.DEFAULT_SECTION, "google")
	//获取mysql分区下的mysql配置信息 ★如果键名重复 会进行覆盖
	getinfo2, err := cfg.GetValue("testament", "user")
	if err != nil {
		fmt.Println("username 获取失败！", err)
		return
	}
	//读取整个分区的信息 返回的数据格式为：map[string]string
	getinfo3, err := cfg.GetSection("super")

	fmt.Println(getinfo1, getinfo2) //www.google.com root
	fmt.Println(getinfo3)           //map[password:123456 url:(127.0.0.1:3306)/baidu username:root]

	//写入操作 - 修改mysql信息 [★] 读取正常 此处验证并没有正常写入修改：已解决 没有保存设置 goconfig.SaveConfigFile()
	isInsert := cfg.SetValue("mysql", "username", "admin456")                 //指定分区
	isInsert = cfg.SetValue(goconfig.DEFAULT_SECTION, "username", "admin123") //默认分区 没有则创建
	if err != nil {
		//fmt.Println("isInsert操作失败！")
		log.Printf("无法设置键值%s:%s", "username", err)
		return
	}

	if isInsert {
		//fmt.Println("新值写入成功！")
		log.Printf("新值:%s写入成功！", "username")
	} else {
		log.Printf("原有值:%s覆盖成功！", "username")
	}

	//默认会生成一个txt文本的文件
	err = goconfig.SaveConfigFile(cfg, "22")
	//err = goconfig.SaveConfigFile(cfg,"22.md")
	if err != nil {
		fmt.Println(err)
		return
	}
}

/*
-  注释读写操作：
	comment := cfg.GetSectionComments("super")
	comment = cfg.GetKeyComments("super", "key_super")
	v := cfg.SetKeyComments("super", "key_super", "# 这是新的键注释")
	v = cfg.SetSectionComments("super", "# 这是新的分区注释")
*/
func main03() {
	cfg, err := goconfig.LoadConfigFile(path)
	if err != nil {
		fmt.Println("加载配置文件失败！")
		return
	}
	//获取注释信息 	此处验证并没有返回注释：看无闻讲解说默认只有键值的注释 没有分区注释
	getInfo1 := cfg.GetSectionComments(goconfig.DEFAULT_SECTION) //默认分区
	getInfo2 := cfg.GetSectionComments("")                       //默认分区
	getInfo3 := cfg.GetSectionComments("super")                  //指定demo分区

	log.Printf("默认分区注释信息为：%s", getInfo1)  //输出: ""
	log.Printf("默认分区注释信息为：%s", getInfo2)  //输出: ""
	log.Printf("super注释信息为：%s", getInfo3) //输出: # 该行注释会作为 super 分区节点的注释,冒号或等号两侧的空白符号均会被自动删除

	comment1 := cfg.GetKeyComments("super", "key_super") //输出: 获取分区键的注释
	log.Printf("super分区下key_super注释信息为：%s", comment1)    //输出: # 该注释会作为键 key_super 的注释
	/*
		脚下留心：
			注释分为：分区注释，键值对注释
		小技巧：
			如果不想写一长串 goconfig.DEFAULT_SECTION 可以使用 ""进行替代
				s1 := cfg.GetSectionComments(goconfig.DEFAULT_SECTION)
				s2 := cfg.GetSectionComments("")
			即：s1与s2获取实现效果等同
	*/
	//设置新的分区注释 返回布尔值
	v := cfg.SetSectionComments("super", "新注释测试")
	//如果写入空字符串代表删除
	//v = cfg.SetSectionComments("super","")
	if v {
		log.Printf("super分区注释插入或删除成功！%v", v)
	} else {
		log.Printf("super分区注释重写成功！%v", v)
	}

	//设置新的键注释
	k := cfg.SetKeyComments("super", "key_super", "测试新的键注释！")
	//设置空字符串这代表删除
	//k = cfg.SetKeyComments("super","key_super","")
	if k {
		log.Printf("super分区下key_super插入或删除成功！%v", k)
	} else {
		log.Printf("super分区下key_super注释重写成功！%v", k)
	}
	err = goconfig.SaveConfigFile(cfg, "test2.ini")
	if err != nil {
		fmt.Println("配置信息保存失败！")
		return
	}

}

/*
	-  类型转换读取：  vInt, err := cfg.Int("must", "int")
	-  Must 系列方法：  vBool := cfg.MustBool("must", "bool")
	-  删除指定键值：  ok := cfg.DeleteKey("must", "string")
	-  保存配置文件：  err = goconfig.SaveConfigFile(cfg, "conf_save.ini")
*/
func main04() {
	//加载配置文件
	cfg, err := goconfig.LoadConfigFile(path)
	if err != nil {
		fmt.Println("加载配置文件失败！")
		return
	}

	//A：类型转换读取测试
	vInt, err := cfg.Int("must", "int")
	//vInt,err = cfg.Int("must","float64") //err
	if err != nil {
		//fmt.Println("无法获取对应的键值：",err)
		//Fatalf等价于{Printf(v...); os.Exit(1)} 退出其本函数
		log.Fatalf("无法获取对应的的键值：%s", err)
	}
	log.Printf("获取到的Int类型值为：%v", vInt)

	/*
		B：Must方法类型读取
		问题思考:如果是一个布尔型的false和一个字符串型的false 那如何定义:是否配置文件中 基本不会出现这种情况。。
	*/
	vBool := cfg.MustBool("must", "boo1") // true -> true
	vBool = cfg.MustBool("must", "boo2")  // 0 -> false
	vBool = cfg.MustBool("must", "boo3")  // false -> false
	vBool = cfg.MustBool("must", "int")   // 123 -> false
	//验证得知：只有在为布尔型true的时候 才会输出true其他类型或布尔型false一律输出为false

	//Must获取类型 must如果不存在返回0 0值即为布尔型的false 不会返回错误
	vBool = cfg.MustBool("must", "int222") // 123 -> false

	/*
		-  类型转换读取：  vInt, err := cfg.Int("must", "int")
		-  Must 系列方法：  vBool := cfg.MustBool("must", "bool")
		脚下留心：
			以上两种方法都可以获取配置信息的数据类型；
			区别在于 类型转换在读取不存在和不符合的数据类型时，会输出一个错误
			而must方法 会返回一个零值 即布尔型的false 这样更方便与我们做类型的判断
			if !vBool {
				log.Print("输出为false")
			}
	*/
	log.Printf("获取到的Bool类型值为：%v", vBool)

	//C.删除指定键值  返回布尔类型
	ok1 := cfg.DeleteKey("must", "string")
	if !ok1 {
		log.Printf("must键值string删除失败：%v", ok1)
	}
	//D.删除分类节点
	ok2 := cfg.DeleteSection("testament")
	if !ok1 {
		log.Printf("must分类节点：%v", ok2)
	}

	//保存配置文件 param A 要保存的对象 param B
	err = goconfig.SaveConfigFile(cfg, "test.ini")
	if err != nil {
		log.Printf("无法保存配置文件：%v", err)
		return
	}
}

var path4 = `E:\Go_WorkSpace\local-project\project5-1\src\`

/*
	05-多文件覆盖加载
	语法：
		cfg, err := goconfig.LoadConfigFile("conf.ini", "conf2.ini")
		err = cfg.AppendFiles("conf3.ini")
*/
func main05() {
	//加载配置文件
	cfg, err := goconfig.LoadConfigFile(path4+`config4.ini`, path4+`config5.ini`)
	if err != nil {
		log.Printf("配置文件加载失败！%s", err) //仅输出 不写return仍会继续执行后续程序
		return
	}
	v, err := cfg.GetValue("", "google")
	v, err = cfg.GetValue("", "key_default")
	if err != nil {
		log.Printf("value获取失败！%s", err)
		return
	}
	/*
		总结：
			如果两边的值都没有重复的相当于把两份配置文件合并成了一份
		如果有重复着 则后者覆盖前者的键值
	*/

	//追加文件配置 可以追加多个值
	err = cfg.AppendFiles(path4 + `config6.ini`)
	v, err = cfg.GetValue("", "google")
	//如果有相同的字段 追加后的值也会覆盖掉原有的值
	v, err = cfg.GetValue("", "key_default")
	if err != nil {
		log.Printf("value获取失败！%s", err)
		return
	}
	fmt.Println(v)
	/*
		为什么使用log 而不使用fmt 使用log的好处：
			添加了输出时间
			线程安全
			方便对日志信息进行转存，形成日志文件
		相关文章Go09延伸篇
	*/

	/*
		06-配置文件重载 重载的顺序是根据我们覆盖和追加的顺序(可以通过修改方法来重写)
		应用场景：
			比如问在网页中修改了一个配置文件 这个时候 如果通过Reload()方法进行方法重载
			就省去了我们重新启动服务的便捷操作 通过重载的方式 使得我们修改的配置 即时生效
	*/
	err = cfg.Reload()
	if err != nil {
		log.Printf("配置文件重载错误！%s", err)
		return
	}

	/*
		07-为配置文件设置缺省值 [缺省值的意思就是默认值] 返回信息为刚才设置的默认值
		-  为 Must 系列方法设置缺省值：  vBool := cfg.MustBool("must", "bool404", true)
		常用作于设置布尔值 进行判断 一开始自己错误的认为 设置默认值 保存配置后 具体文件
		会发生改变 经测试不是这样的 主要应用在数据输出上 不会对原先的文件进行修改保存
	*/
	vBool := cfg.MustValue("must", "string2", "666")
	fmt.Println("+", vBool, "+")
	err = goconfig.SaveConfigFile(cfg, "test2")
	if err != nil {
		log.Printf("数据信息保存失败！%s", err)
		return
	}

	v, err = cfg.GetValue("must", "string")
	if err != nil {
		log.Printf("value获取失败！%s", err)
		return
	}
	fmt.Println(v)

	/*
		08-递归读取键值 search=%(键名)s
		递归读取 必须是已经加载好的键值数据才可以被读取
		最大递归层级为200层 防止内存溢出
	*/
	_, err = cfg.GetValue("parent", "search")
	if err != nil {
		log.Printf("获取递归读取键值失败！%s", err)
		return
	}
	//log.Printf("输出递归读取键值：%s",inf)
	//fmt.Printf("输出递归读取键值：%s",inf)

	/*
		09-子孙区覆盖读取
		什么是子孙区覆盖读取 parent.child
		子孙区会覆盖父级的键值信息
	*/
	pcv, err := cfg.GetValue("parent.child", "age") //第一种情况如果两者都有重复键 则输出为子孙区的信息
	pcv, err = cfg.GetValue("parent.child", "sex")  //第二种情况如果仅父级有 则输出父级的键值信息 不会报错
	//pcv , err = cfg.GetValue("parent.child", "sex2") //第三种情况如果两者都没有 这输出 not found 错误
	if err != nil {
		log.Printf("子孙分区数据信息读取失败！%s", err)
		return
	}
	fmt.Println(pcv)

	/*
		-  09自增键名获取  ini写法语句 以-=起始 获取通过GetSection()方法进行获取
		-  10获取整个分区：  sec, err := cfg.GetSection("auto increment")
	*/
	//如何获取自增键名
	av, err := cfg.GetSection("auto increment")
	if err != nil {
		log.Printf("自增键名获取失败！%s", err)
		return
	}
	fmt.Println(av)       //map[#1:hello #2:go #3:config]
	fmt.Println(av["#1"]) //hello

	//添加测试 定位到最后一个键 [测试写入提交]
	av, _ = cfg.GetSection("auto increment")
	var a string
	a = "#" + strconv.Itoa(len(av)+1) //先定义再追加

	cfg.SetValue("auto increment", a, "哈哈")
	err = goconfig.SaveConfigFile(cfg, "list")
	if err != nil {
		log.Printf("数据保存失败！%s", err)
		return
	}

	//更新测试 先定位到指定的键
	cfg.SetValue("auto increment", "#3", "哈哈")
	err = goconfig.SaveConfigFile(cfg, "list")
	if err != nil {
		log.Printf("数据保存失败！%s", err)
		return
	}
}
