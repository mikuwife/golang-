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
	p.SetInfoP()

	// 方法值
	pInfoP := p.SetInfoP
	pInfoP() // 等价于 p.SetInfoP()
}
