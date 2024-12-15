package main

import (
	"fmt"
	"os"
)

func main() {
	// 接收用户传递的参数，字符串切片
	list := os.Args

	n := len(list)
	fmt.Println("n = ", n)

	for index, data := range list {
		fmt.Printf("index = %d, data = %s\n", index, data)
	}
}
