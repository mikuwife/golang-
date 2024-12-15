package main

import "fmt"

func main() {
	var a int = 10
	fmt.Printf("a = %v\n", a)   // 变量存储的数据
	fmt.Printf("&a = %v\n", &a) // 变量的地址 &

	// *int 类型保存变量地址
	var pa *int
	pa = &a
	fmt.Println("pa =", pa)

	// 修改a的内存数据
	*pa = 200
	fmt.Println("a =", a)

	// 指向空指针
	var p *int
	fmt.Println("p =", p)

	*p = 200
}
