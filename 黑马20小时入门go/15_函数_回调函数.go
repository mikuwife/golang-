package main

import "fmt"

type funcType func(int, int) int

func mult(a, b int) (result int) {
	result = a * b
	return
}

/*
回调函数：函数中有一个参数类型是函数类型的就是回调函数
*/
func Calc(a, b int, fTest funcType) (result int) {
	fmt.Println("计算中...")
	// 多态: 多种形态，调用同一个接口，不同的表现
	result = fTest(a, b)
	return
}

func main() {
	fmt.Println(Calc(10, 20, mult))
}
