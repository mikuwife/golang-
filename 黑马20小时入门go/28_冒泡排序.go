package main

import "fmt"

func main() {
	a := [5]int{5, 2, 1, 4, 3}

	bubbleSort(&a)

	fmt.Println(a)
}

func bubbleSort(a *[5]int) {
	len := len(a)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
