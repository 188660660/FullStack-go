package main

import "fmt"

//1.接口的实现
type USBer interface {
	Read()
	Write()
}

//创建对象
type USBdev struct {
	id     int
	name   string
	rspeed int
	wspeed int
}

type Mobile struct {
	USBdev
	call string
}

type Upan struct {
	USBdev
}

//实现方法
func (m *Mobile) Read() {
	fmt.Printf("%s正在读取数据速度为：%d\n", m.name, m.rspeed)
}

func (m *Mobile) Write() {
	fmt.Printf("%s正在写入数据速度为：%d\n", m.name, m.wspeed)
}

func (u *Upan) Read() {
	fmt.Printf("%s正在读取数据速度为：%d\n", u.name, u.rspeed)
}

func (u *Upan) Write() {
	fmt.Printf("%s正在写入数据速度为：%d\n", u.name, u.wspeed)
}

//实现多态
func Test(t USBer) {
	t.Read()
	t.Write()
}

func main090401() {
	var usb USBer
	usb = &Mobile{
		USBdev{101, "手机", 5, 10}, "隔壁老王",
	}
	Test(usb)

	usb = &Upan{USBdev{102, "U盘", 20, 30}}
	Test(usb)

}
