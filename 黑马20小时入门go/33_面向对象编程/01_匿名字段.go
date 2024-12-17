package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person // 匿名字段，继承Person所有成员
	id     int
	addr   string
}

func main() {
	// 初始化
	s1 := Student{
		Person: Person{
			name: "miku",
		},
	}
	fmt.Printf("s1 = %+v\n", s1)

	// 成员操作
	s1.name = "mikucat"
	fmt.Println(s1)
}
