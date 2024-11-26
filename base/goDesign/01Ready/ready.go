package main

import "fmt"

/*
用来测试源代码编译过程中的生成的汇编语言，如何能够从汇编语言的结果中找到优化空间
*/
func testCanSetData() {
	a := 123
	b := 222
	c := 333
	a = 222
	fmt.Println("______the res of a is :", a)
	fmt.Println("______the res of b is :", b)
	fmt.Println("______the res of b is :", c)
}

func main() {
	testCanSetData()
}
