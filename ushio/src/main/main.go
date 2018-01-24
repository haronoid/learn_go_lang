package main

import (
	"fmt"

	"buf"
	"chan2"
	"para"
	"time"
)

func main() {
	fmt.Println("start main")

	chan2.Start()

	buf.Start()

	para.Start()

	time.Sleep(3 * time.Second)
}
