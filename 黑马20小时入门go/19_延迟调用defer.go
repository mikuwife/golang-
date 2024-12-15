package main

import "fmt"

func main() {
	// 后进先出，哪怕函数某一步崩溃，那剩下的defer也会执行
	defer fmt.Println("aaa")
	defer fmt.Println("bbb")
	defer fmt.Println("ccc")

	// defer和匿名函数结合
	a, b := 10, 20
	defer func() {
		fmt.Printf("a = %d, b = %d\n", a, b)
	}()

	a, b = 100, 200
	fmt.Printf("a = %d, b = %d\n", a, b)
}
