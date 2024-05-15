package main

import (
	"fmt"
	"slices"
)

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{11, 12, 13}
	s3 := []int{21, 22, 23}
	// go1.22 之前
	//s := append(s1, s2...)
	//s4 := append(s, s3...)

	// go1.22 slices.Concat
	s4 := slices.Concat(s1, s2, s3)

	fmt.Println(s4)
}
