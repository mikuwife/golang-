package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

// 接收者为值类型，是值的拷贝
func (p Person) SetInfoV() {
	fmt.Println("SetInfoV....")
}

// 接收者为指针类型，是传的指针类型
func (p *Person) SetInfoP() {

	fmt.Println("SetInfoP....")
}

func main() {
	p := new(Person)
	p.SetInfoP()

	// 不管是值还是指针，编译器都能调用到方法
	p.SetInfoV()
}
