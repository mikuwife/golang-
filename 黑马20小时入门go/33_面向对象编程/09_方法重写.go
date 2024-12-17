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


type Student struct {
	Person
	id   int
	addr string
}

// 重写继承自Person的PrintInfo()方法
func (s *Student) PrintInfo() {
	fmt.Println("s = ", s)
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

	// 显示调用Person的方法
	s.Person.PrintInfo()
}