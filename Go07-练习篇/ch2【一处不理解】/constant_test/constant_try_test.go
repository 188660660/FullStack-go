package constant_test

import "testing"

const(
	a = iota + 1 //快速设置连续的值
	b
	c
)

//标记★
const(
	d = 1 << iota
	e
	f
)
func TestConstTry1(t *testing.T) {
	t.Log(a,b,c)
}

func TestConstTry2(t *testing.T)  {
	d := 0
	t.Log(d&d == d,d&e==e,d&f==f)
}
