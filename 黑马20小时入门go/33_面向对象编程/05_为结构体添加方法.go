package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

// 打印结构体信息
func (p *Person) PrintInfo() {
	fmt.Println("p = ", p)
}

// 设置结构体字段
func (p *Person) SetInfo(name string, sex byte, age int) {
	p.name = name
	p.sex = sex
	p.age = age
}

func main() {
	p1 := new(Person)
	p1.PrintInfo()
	p1.SetInfo("miku", '1', 24)
	p1.PrintInfo()
}
