package main

import "fmt"

// 1. 不定参数函数
func myFunc01(args ...int) {
	fmt.Println("%T", args)
}

func main() {
	myFunc01(1, 2, 3)
}
