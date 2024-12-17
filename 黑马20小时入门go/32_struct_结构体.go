package main

import "fmt"

// Student 定义结构体类型
type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	// 顺序初始化，每个成员都必须初始化
	var s1 Student = Student{1, "miku", '1', 24, "beijing"}
	fmt.Println("s1 = ", s1)

	// 指定成员初始化，不指定的赋值为0值
	s1 = Student{id: 10}
	fmt.Println("s1 = ", s1)

	// 结构体指针
	s2 := &Student{age: 20}
	fmt.Printf("s2 type = %T\n", s2)
	fmt.Printf("s2 = %v\n", s2)

	// 操作结构体成员
	var s3 Student
	s3.age = 20
	s3.name = "miku"
	fmt.Println("s3 = ", s3)

	// 指针变量
	var s4 Student
	var ps4 *Student
	ps4 = &s4
	// 1. 通过指针操作成员 ps4.id 等价于 (*ps4).id
	ps4.id = 20
	(*ps4).id = 40
	fmt.Println("ps4.id = ", ps4.id)
	// 2. 通过new申请一个结构体
	ps5 := new(Student)
	ps5.id = 50
	(*ps5).id = 100
	fmt.Println("ps5.id = ", ps5.id)

	// 结构体比较 只支持 == !=
	fmt.Println(s4 == s3)

	// 结构体作为函数参数，使用引用传递才能修改成员
	s6 := &Student{
		id:   10,
		name: "miku",
		sex:  '1',
		age:  18,
		addr: "jx",
	}
	fmt.Println("s6 = ", s6)
	changeSex(s6)
	fmt.Println("change s6 = ", s6) // 成功改变

}

func changeSex(p *Student) {
	p.sex = '0'
}
