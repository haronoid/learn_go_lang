package main

import (
	"fmt"

	ani "./animals"
)

const (
	A = iota // A == 0
	B        // B == 1
	C        // C == 2
)

const (
	onYourMark = "内部定数スコープ"
	OnYourMakr = "公開定数スコープ"
)

// 内部参照のみの関数
// アルファベットじゃない日本語名の場合はすべて内部関数になる
func myFunc() {
	fmt.Println(onYourMark)
}

// 外部に公開される関数
func MyFunc() {
	fmt.Println(OnYourMakr)
}

func main() {
	fmt.Println(ani.ElephantFeed()) // => "Grass"
	fmt.Println(ani.MonkeyFeed())   // => "Banana"
	fmt.Println(ani.RabbitFeed())   // => "Carrot"

	fmt.Printf("A=%d\n", 159) // => "A=159"

	f := func(x, y int) int { return x + y }

	fmt.Printf("f=%d\n", f(5, 6)) // => "f=11"

	f2 := funcReturn()

	f2("a") // store == " a"
	f2("b") // store == " a b"
	f2("c") // => "store= a b c"

	callFunc(func() { fmt.Print("Hell!") }) // => "Hell!Hell!"

	fmt.Println(A + B + C) // => 3

	myFunc() // => "内部定数スコープ"
	MyFunc() // => "公開定数スコープ"
}

// 戻り値を関数にする
func funcReturn() func(string) {
	var store string
	var cnt int
	return func(next string) {
		cnt++
		store += " " + next

		if cnt == 3 {
			fmt.Println("store=" + store)
		}
	}
}

// 関数をパラメータで受け取る
func callFunc(f func()) {
	f()
	f()
	fmt.Println()
}
