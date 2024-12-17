package main

import "fmt"

type mystr string

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	*Person // 指针结构体匿名字段
	int     // 非结构体匿名字段
	mystr
}

func main() {
	s1 := Student{
		&Person{"miku", '1', 18},
		666,
		"hehe",
	}

	fmt.Println("s1", s1)
	fmt.Println("s1.Person", s1.Person)

	var s2 Student
	s2.Person = new(Person)
	s2.age = 20
	fmt.Println("s2 = ", s2)
}
