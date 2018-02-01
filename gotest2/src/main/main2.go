package main

import (
	"fmt"
)

func init() {
	fmt.Println("main2 start")
}

func main2() {
	fmt.Println("main2")

	// point

	var p *int

	fmt.Println(p == nil) // => "true"

	var i int
	p2 := &i
	fmt.Printf("%T\n", p2)
	p3 := &p2
	fmt.Printf("%T\n", p3)

	i = 5
	*p2 += 5
	// pointerの値を参照する
	fmt.Println(i)

	p4 := &[3]int{1, 2, 3}
	s := (*p4)[0] + (*p4)[1] + (*p4)[2]
	fmt.Println(s)
	s += p4[0] + p4[1] + p4[2]
	fmt.Println(s)

	ip := &i
	fmt.Printf("Type=%T Address=%p Value=%d\n", ip, ip, *ip)

	mi := MyInt{A: 4, B: 5}
	mi.calc()

	fmt.Println(mi)

}

//type MyInt int
type MyInt struct {
	A   int
	B   int
	SUM int
}

func (m *MyInt) calc() {
	m.SUM = m.A + m.B
}
