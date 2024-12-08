package main

import "fmt"

func main() {

	// 变量初始化 声明同时赋值
	var a int = 20
	fmt.Println("a = ", a)

	// 赋值
	a = 200
	fmt.Println("a = ", a)

	// 自动类型推导
	c := "miku"
	// 格式化输出
	fmt.Printf("c type is %T\n", c)

	// 多重赋值
	num1, num2 := 10, 20
	fmt.Printf("num1 = %d, num2 = %d\n", num1, num2)

	num1, num2 = num2, num1
	fmt.Printf("num1 = %d, num2 = %d\n", num1, num2)

	// 匿名变量 配合函数返回值使用
	e, _, _ := test()
	fmt.Println(e)
}

func test() (int, int, int) {
	return 1, 2, 3
}
