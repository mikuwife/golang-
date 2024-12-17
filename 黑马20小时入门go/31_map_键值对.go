package main

import (
	"fmt"
)

func main() {
	// 初始化map
	m1 := map[int]string{
		1: "hello",
	}
	fmt.Println(m1)

	m2 := make(map[int]string, 5)
	m2[1] = "world"
	fmt.Println(m2)

	// 赋值并初始化
	m3 := map[int]string{
		1: "a",
	}
	fmt.Println(m3)

	// 遍历map
	for key, value := range m3 {
		fmt.Println("key", key)
		fmt.Println("value", value)
	}

	// 添加，修改
	// key不存在则添加，存在则修改
	m3[4] = "111"

	// 查询键值对是否存在
	value, ok := m3[4]
	if ok {
		fmt.Println("value = ", value)
	} else {
		fmt.Println("value不存在")
	}

	// 删除key
	delete(m3, 4)
	fmt.Println(m3)

	// map作为函数参数传递
	test05(m3)
	fmt.Println("test() m3 = ", m3)
}

func test05(m map[int]string) {
	m[10] = "200"
}
