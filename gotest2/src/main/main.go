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

	// slice

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

	a := [5]int{1, 2, 3, 4, 5}
	a2 := a[1:3]
	fmt.Println(a2)

	// map

	m := make(map[string]string)

	m["A"] = "Apple"
	m["B"] = "Banana"

	r, ok := m["A"] // r == "Apple", ok == true
	_, ok = m["C"]  // r == "", ok == false

	fmt.Printf("m=%s\n", m)
	fmt.Printf("r=%s, ok=%t\n", r, ok)

	// chan

	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	d, e, f := <-ch, <-ch, <-ch
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	ch <- 4

	close(ch)
}

func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}
