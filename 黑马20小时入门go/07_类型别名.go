package main

import "fmt"

type (
	myInt int
	char  byte
)

func main() {
	var a myInt = 20
	var ch char = '1'

	fmt.Printf("a type = %T, ch type = %T\n", a, ch)
}
