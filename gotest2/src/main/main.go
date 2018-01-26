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

	s := 2

	if s == 1 {
		main1()
	}

	if s == 2 {
		main2()
	}

}
