package main

import "fmt"

func test03() func() int {
	var a int
	return func() int {
		a++
		return a * a
	}
}

func main() {

	f := test03()
	fmt.Println("f = ", f())
	fmt.Println("f = ", f())
	fmt.Println("f = ", f())
}
