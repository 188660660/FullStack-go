package main

import "fmt"

func main020701() {
	a := 10
	b := 10
	//a = a + 5
	//a += 5
	//a -= 5
	//a /= 5
	//a %= 5

	a += 5 * 3 //运算顺序: a = a + 15
	b += 5 * b //运算顺序：a = a + 50

	fmt.Println(a)
	fmt.Println(b)

	/*
		!总结：
			加等于【+=】 减等于【-=】乘等于【*=】除等于【/=】取余等于【%=】
	*/
}
