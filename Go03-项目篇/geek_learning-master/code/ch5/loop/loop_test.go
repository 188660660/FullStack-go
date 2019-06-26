package loop_test

import "testing"

func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n) //0 1 2 3 4
		n++
	}
}
