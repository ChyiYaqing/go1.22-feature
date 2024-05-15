package main

import (
	"fmt"
	"slices"
)

func main() {
	// go1.22将对缩小切片片段/大小的相关函数的结果行为进行调整
	// 切片经过缩小后新长度和旧长度之间的元素将会归位零
	s1 := []int{1, 2, 3, 4}
	s2 := slices.Delete(s1, 1, 3)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	/** 程序运行结果
	go 1.22
		s1: [1 4 0 0]
		s2: [1 4]

	go 1.21
		s1: [1 4 3 4]
		s2: [1 4]
	*/
}
