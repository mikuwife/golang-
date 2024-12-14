package main

import "fmt"

func test3(a int) {
	if a != 0 {
		fmt.Println("a = ", a)
		test3(a - 1)
	}
	return
}

func sum(a int) (num int) {
	if a == 1 {
		num = 1
		return
	}
	return a + sum(a-1)
}

func main() {
	test3(3)
	fmt.Println(sum(100))
}
