package main

import "fmt"

func main() {
	num := 0
	for {
		num++
		if num == 20 {
			// break    // 跳出最近的那个循环体
			continue // 跳出当前循环条件，继续执行下一个循环
		}
		fmt.Println("num = ", num)
	}
}
