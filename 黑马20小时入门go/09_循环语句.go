package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}

	fmt.Println("sum = ", sum)

	str1 := "hello, world!"
	for index, value := range str1 {
		fmt.Printf("i = %d, value = %c\n", index, value)
	}
}
