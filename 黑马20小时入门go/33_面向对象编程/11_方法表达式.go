package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func (p Person) SetInfoV() {
	fmt.Println("SetInfoV....")
}

func (p *Person) SetInfoP() {
	fmt.Println("SetInfoP....")
}

func main() {
	p := new(Person)

	// 方法表达式
	f := (*Person).SetInfoP
	f(p)

	f1 := (Person).SetInfoV
	f1(*p)
}
