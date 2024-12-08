package main

import "fmt"

const (
	a = iota
	b
	c
	d
)

func main() {
	// const 声明常量
	const N = "hello"
	fmt.Printf("N type is %T\n", N)

}
