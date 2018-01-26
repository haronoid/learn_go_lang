package main

import (
	"fmt"
)

func init() {
	fmt.Println("main1 start")
}
func main1() {
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

	// select

	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)
	ch1 <- 1
	ch2 <- 2
	select {
	case <-ch1:
		fmt.Println("ch1")
	case <-ch2:
		fmt.Println("ch2")
	case ch3 <- 3:
		fmt.Println("ch3")
	default:
		fmt.Println("default")
	}

}

func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}
