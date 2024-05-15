package main

import (
	"fmt"
	"slices"
)

func main() {
	s1 := []int{1, 2, 2, 2, 4}
	s2 := slices.Compact(s1)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	/** 输出结果
	go 1.22
		s1: [1 2 4 0 0]
		s2: [1 2 4]

	go 1.21
		s1: [1 2 4 2 4]
		s2: [1 2 4]

	*/
}
