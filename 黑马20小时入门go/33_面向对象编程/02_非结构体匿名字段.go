package main

import "fmt"

type mystr string

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person // 结构体匿名字段
	int    // 非结构体匿名字段
	mystr
}

func main() {
	s1 := Student{
		Person{"miku", '1', 18},
		666,
		"hehe",
	}

	fmt.Println("s1.int = ", s1.int)
}
