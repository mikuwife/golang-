package main

import "fmt"

func main() {
	// 初始化数组
	var id [50]int
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	b := [5]int{1, 2, 3, 4, 5}
	// 部分初始化
	c := [5]int{1}
	// 指定某个下标元素初设化
	d := [5]int{1: 10}

	fmt.Println(id)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
