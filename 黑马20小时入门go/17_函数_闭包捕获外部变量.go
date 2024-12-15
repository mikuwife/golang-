package main

import "fmt"

func main() {
	a := 1
	b := 2

	func() {
		a = 20
		b = 30
		fmt.Println("内部 a, b = ", a, b)
	}()

	fmt.Println("外部 a, b = ", a, b)
}
