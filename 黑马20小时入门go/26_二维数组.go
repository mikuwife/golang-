package main

import "fmt"

func main() {
	// 有多少[]就是多少维
	// 有多少[]就用多少个循环
	var a [3][4]int

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Printf("%v", a[i][j])
		}
		fmt.Println(" ")
	}
}
