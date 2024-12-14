package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

type FuncType func(int, int) int

func main() {
	fmt.Println(Add(1, 2))

	// 多态
	var fTest FuncType
	fTest = Add
	fmt.Println(fTest(1, 2))

	fTest = Max
	fmt.Println(fTest(10, 2))
}
