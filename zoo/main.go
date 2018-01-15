package main

import (
	"fmt"

	"./animals"
)

func main() {
	fmt.Println(animals.ElephantFeed()) // => "Grass"
	fmt.Println(animals.MonkeyFeed())   // => "Banana"
	fmt.Println(animals.RabbitFeed())   // => "Carrot"

	fmt.Printf("A=%d\n", 159) // => "A=159"
}
