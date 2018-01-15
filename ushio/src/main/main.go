package main

import (
	"fmt"

	"buf"
	"chan2"
	"para"
)

func main() {
	fmt.Println("start main")

	chan2.Start()

	buf.Start()

	para.Start()
}
