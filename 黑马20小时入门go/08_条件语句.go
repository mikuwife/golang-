package main

import "fmt"

func main() {

	// if 条件语句
	if a, _ := test1(); a != 10 {
		fmt.Printf("a不是10\n")
	} else {
		fmt.Printf("a等于10\n")
	}

	// switch 语句
	switch a, _ := test1(); a { // 可以不写条件
	case 1: // 可以同时匹配多个
		fmt.Println("1楼")
		//break // 跳出当前循环
		fallthrough // 强制继续执行下一个case
	case 2:
		fmt.Println("2楼")
	default:
		fmt.Println("地下室")
	}
}

func test1() (int, int) {
	a, b := 1, 20
	return a, b
}
