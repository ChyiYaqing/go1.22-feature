package main

import (
	"fmt"
	"slices"
)

func main() {
	s1 := []int{1, 2, 3}
	s2 := slices.Insert(s1, 4)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	/** 输出结果

	go 1.21
		s1: [1 2 3]
		s2: [1 2 3]

	go 1.22
		➜ go run pkg/slice_insert.go
		panic: runtime error: slice bounds out of range [4:3]

		goroutine 1 [running]:
		slices.Insert[...]({0x14000116018?, 0x1400010aed8?, 0x1007954c8?}, 0x60?, {0x0?, 0x60?, 0x140000021c0?})
		        /usr/local/go/src/slices/slices.go:133 +0x444
		main.main()
		        /Users/chyiyaqing/Chyi/github.com/go1.22-feature/pkg/slice_insert.go:10 +0x6c
		exit status 2
	*/
}
