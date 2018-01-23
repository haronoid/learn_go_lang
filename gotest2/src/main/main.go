package main

import (
	"fmt"
	"runtime"
)

func init() {
	fmt.Println("初期化中")
}

func main() {
	go fmt.Println("hello world")

	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())
	fmt.Printf("Version: %s\n", runtime.Version())

	s := make([]int, 5, 10)
	s[2] = 2
	fmt.Println(s)
	fmt.Printf("len:%d, cap:%d\n", len(s), cap(s))
	s = append(s, 6, 7, 8)
	fmt.Println(s)
	fmt.Printf("len:%d, cap:%d\n", len(s), cap(s))
	s2 := []int{1, 2, 3, 4, 5}
	copy(s, s2)
	fmt.Println(s)
	fmt.Printf("len:%d, cap:%d\n", len(s), cap(s))

	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(1, 2, 3, 5, 6, 7))

	fmt.Println(sum(s2...))
}

func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}
