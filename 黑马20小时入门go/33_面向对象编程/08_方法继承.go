package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func (p *Person) PrintInfo() {
	fmt.Printf("name = %s, sex = %c, age = %d", p.name, p.sex, p.age)
}

// 学生继承Person，成员和方法都会被继承
type Student struct {
	Person
	id   int
	addr string
}

func main() {
	s := &Student{
		Person{
			name: "miku",
			sex:  '1',
			age:  24,
		},
		10,
		"beijin",
	}
	s.PrintInfo()
}
