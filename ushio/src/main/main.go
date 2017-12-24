package main

import (
	"buf"
	"chan2"
	"fmt"
	"para"
)

func main() {
	fmt.Println("start main")

	chan2.Start()

	buf.Start()

	para.Start()
}
