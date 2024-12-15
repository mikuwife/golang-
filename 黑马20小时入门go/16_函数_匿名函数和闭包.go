package main

import "fmt"

func main() {
	a := "hello"
	b := "world"

	// 匿名函数，没有函数名
	f1 := func() {
		fmt.Println("a = ", a)
		fmt.Println("b = ", b)
	}
	f1() // 调用

	// 给函数类型起别名
	type FuncType func()
	// 声明变量
	var f2 FuncType
	f2 = f1
	f2()

	// 定义匿名函数同时调用
	func() {
		fmt.Println("hello, world")
	}()

	func(a, b int) {
		fmt.Println("a + b = ", a+b)
	}(10, 29)
}
