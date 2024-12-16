package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreatNum(p *int) {
	// 随机数种子
	rand.Seed(time.Now().UnixNano())
	var num int
	for {
		num = rand.Intn(10000)
		if num >= 1000 {
			break
		}
	}

	*p = num
}

// GetNum 保存四位数的每一位
func GetNum(randSlice []int, randNum int) {
	randSlice[0] = randNum / 1000       // 千位
	randSlice[1] = randNum % 1000 / 100 // 百位
	randSlice[2] = randNum % 100 / 10   // 十位
	randSlice[3] = randNum % 10         // 个位
}

// OnGame 游戏核心逻辑
func OnGame(randSlice []int) {
	var myNum int
	mySlice := make([]int, 4)

	for {
		for {
			fmt.Println("请输入你要猜测的数字: ")
			fmt.Scan(&myNum)
			if 999 < myNum && myNum <= 10000 {
				break
			}
			fmt.Println("输入的数不符合要求")
		}
		fmt.Println("myNum = ", myNum)

		GetNum(mySlice, myNum)

		n := 0
		for i := 0; i < 4; i++ {
			if mySlice[i] > randSlice[i] {
				fmt.Printf("第%d位大了\n", i+1)
			} else if mySlice[i] < randSlice[i] {
				fmt.Printf("第%d位小了\n", i+1)
			} else {
				fmt.Printf("第%d位猜对了\n", i+1)
				n++
			}
		}

		// 4位数都猜对了
		if n == 4 {
			fmt.Println("全部都猜对...")
			break
		}
	}

}

func main() {
	var randNum int

	// 生成一个随机4位数
	CreatNum(&randNum)
	fmt.Println("randNum: ", randNum)

	randSlice := make([]int, 4)
	// 保存四位数的每一位
	GetNum(randSlice, randNum)
	fmt.Println(randSlice)

	OnGame(randSlice)
}
