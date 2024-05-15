package main

import (
	"fmt"
	"slices"
)

func main() {
	s1 := []int{1, 2, 3, 4}
	s2 := slices.Replace(s1, 1, 3, 99)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	/** 结果输出
	go 1.22
		s1: [1 99 4 0]
		s2: [1 99 4]

	go 1.21
		s1: [1 99 4 4]
		s2: [1 99 4]

	*/
}
