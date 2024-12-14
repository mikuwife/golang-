package main

import "fmt"

func MaxAndMin(a, b int) (max, min int) {
	if a > b {
		max, min = a, b
	} else {
		max, min = b, a
	}

	return
}

func main() {
	max, min := MaxAndMin(1, 2)
	fmt.Printf("max = %v, min = %v", max, min)
}
