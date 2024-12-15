package main

import "fmt"

func main() {
	var p *int

	// 传入一个type 返回一个指针 默认为type类型的0值
	p = new(int)

	*p = 666
	fmt.Println("p =", *p)
}
