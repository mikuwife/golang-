package main

import "fmt"

// 实现两数相加
// 1. 面向过程(函数)
func Add01(a, b int) int {
	return a + b
}

// 2. 面向对象(方法)
type long int

func (a long) Add02(b long) long {
	return a + b
}

func main() {
	var res int
	res = Add01(1, 2) // 普通函数调用
	fmt.Println("Add01... ", res)

	var a long = 200
	res1 := a.Add02(100) // 对象方法调用
	fmt.Println("Add02... ", res1)

}
