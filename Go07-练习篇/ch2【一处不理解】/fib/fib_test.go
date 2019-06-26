package fib

import "testing"

/*
	1. 源码文件以 _test 结尾：xxx_test.go
	2. 测试方法名以 Test 开头：func TestXXX(t *testing.T) {…}
*/
func TestListCh201(t *testing.T) {
	//var a int = 1
	//var b int = 2
	//a := 1
	//b := 2
	var (
		a int = 1
		b int = 3
	)
	//a,b = b,a
	for i:=0;i<5 ;i++  {
		tem := a //tem1 a1
		a = b //a3 b3
		b = tem
	}
	t.Log(a)
	t.Log(b)
}
