package main

import "fmt"

func main01() {
	arr1 := [5]int{10, 20, 30, 0, 0}
	/*
		arr1[low:high:max]
		len = high - low
		cap = max - low
	*/
	slice := arr1[1:4:5]

	// [20, 30, 0] 3 4
	fmt.Printf("slice = %v, len = %v, cap = %v\n", slice, len(slice), cap(slice))

	// 创建切片
	s := []int{}
	fmt.Printf("s = %v, len = %v, cap = %v\n", s, len(s), cap(s))
	// 给切片末尾添加元素
	s = append(s, 1)
	fmt.Printf("s = %v, len = %v, cap = %v\n", s, len(s), cap(s))

}

// 创建切片
func main02() {
	// 1. 自动推导类型，同时初始化
	s1 := []int{1, 2, 3}
	fmt.Println(s1)

	// 2. 使用make函数 make(type, len, cap)
	s2 := make([]int, 5, 6)
	fmt.Printf("s2 = %v, len = %v, cap = %v\n", s2, len(s2), cap(s2))
}

// 切片和底层数组关系
// 对切片的数据进行修改会直接影响到底层数组的数据
func main03() {
	arr1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := arr1[0:4] // 1, 2, 3, 4
	s1[0] = 10
	fmt.Println("arr1 = ", arr1) // [10 2 3 4 5 6 7 8 9 10]

	s2 := s1[:10] // [10 2 3 4 5 6 7 8 9 10]
	fmt.Println("s2 = ", s2)
	s2[len(s2)-1] = 100
	fmt.Println("arr1 = ", arr1)
}

/*
切片内建函数
append：原切片末尾添加元素(扩容是以两倍速度扩容 原容量*2)
copy: 会对齐？[3,4] , [1, 2, 1, 1, 1] ====> [3 4 1 1 1]
*/
func main04() {
	s1 := make([]int, 0)
	s1 = append(s1, 1, 2)
	fmt.Println("s1 = ", s1)
	s1 = append(s1, 3)
	fmt.Println("s1 cap = ", cap(s1))

	src := []int{3, 4}
	dst := []int{1, 2, 1, 1, 1}
	copy(dst, src)
	fmt.Println(dst)
}

func main() {
	s1 := []int{4, 2, 3, 6, 7, 1}
	bubbleSort1(s1)
	fmt.Println(s1)
}

func bubbleSort1(a []int) {
	len := len(a)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
