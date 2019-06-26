package fib

import (
	"testing"
)
/*
1. 源码文件以 _test 结尾：xxx_test.go
2. 测试方法名以 Test 开头：func TestXXX(t *testing.T) {…}
*/
func TestFibList(t *testing.T) {
	// var a int = 1
	// var b int = 1
	// var (
	// 	a int = 1
	// 	b     = 1
	// )
	a := 1
	// a := 1
	b := 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}

}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	// tmp := a
	// a = b
	// b = tmp
	a, b = b, a
	t.Log(a, b)
}
