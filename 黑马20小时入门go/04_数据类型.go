package main

import "fmt"

func main() {
	// bool 类型
	var b1 bool
	var b2 bool = true
	b3 := false
	fmt.Printf("b1 = %v, b2 = %v, b3 = %v\n", b1, b2, b3)

	// float 类型
	var f1 float64 = 3.1
	fmt.Printf("f1 = %T\n", f1)

	// byte 字符类型
	var c1 byte = 97
	fmt.Printf("c1 = %c, c1 = %d\n", c1, c1)

	// string 类型
	var s1 string = "hello, world"
	fmt.Printf("len(s1) = %d\n", len(s1))
	// 字符串是由字符组成的
	fmt.Printf("s1[0] = %c, s1[1] = %c\n", s1[0], s1[1])

	// complex 复数类型
	var t complex128 = 2.1 + 3.14i
	fmt.Println(t)

}
