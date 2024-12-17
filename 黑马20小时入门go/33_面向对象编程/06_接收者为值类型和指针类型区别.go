package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

// 接收者为值类型，是值的拷贝
func (p Person) SetInfo(name string, sex byte, age int) {
	p.name = name
	p.sex = sex
	p.age = age
	fmt.Printf("&p = %p\n", &p)
}

// 接收者为指针类型，是传的指针类型
func (p *Person) SetInfo1(name string, sex byte, age int) {
	p.name = name
	p.sex = sex
	p.age = age
	fmt.Printf("&p = %p\n", &p)
}

func main() {
	var p1 Person
	fmt.Printf("&p1 = %p\n", &p1)

	// p1.SetInfo("miku", 1, 20)
	(&p1).SetInfo1("miku", 1, 20)
	fmt.Println(p1)

}
