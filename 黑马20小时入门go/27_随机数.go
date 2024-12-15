package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 设置种子
	// 种子参数都一样，则每次程序生成的随机数都一样
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		fmt.Println("rand = ", rand.Intn(100))
	}
}
