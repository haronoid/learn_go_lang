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

	fruits := [3]string{"apple", "banana", "cherry"}

	for i, s := range fruits {
		fmt.Printf("Fruits[%d]=%s\n", i, s) // => "apple" "banana" "cherry"
	}

	for i, r := range "👪BC" {
		fmt.Printf("[%d] -> %d\n", i, r)
	}

	s := "A"

	switch s {
	case "A":
		s += "1"
		// gallthroughがない場合は、ここで処理が終了する。
	case "B":
		s += "2"
		fallthrough
	case "C":
		s += "3"
		fallthrough
	default:
	}
	fmt.Println(s) // A => "A1", B ==> "B23"

	anything("abc") // => "abc"
	anything(1.5)   // => "1.5"
	anything("日本語") // => "日本語"
	anything('a')   // => "97"

	var x interface{} = 3

	i, isInt := x.(int) // i == 3, isInt == true
	fmt.Printf("%d %t\n", i, isInt)

R:

	switch v := x.(type) {
	case bool:
		fmt.Println("x is bool")
	case int:
		fmt.Println("x is int")
	case string:
		fmt.Println("x is string")
	default:
		fmt.Printf("x is %#v\n", v)
	}

	if x == 3 {
		x = 4
		goto R
	}

	// panicが起きる処理より前じゃないと動かない
	defer func() {
		fmt.Println("recover")
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()

	runDefer()

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

func anything(a interface{}) {
	fmt.Println(a)
}

func runDefer() {

	// deferは関数処理の最後に実行される処理, 一周のFinally処理

	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	panic("runtime error") // throw new Exception

	defer fmt.Println("defer3")

	fmt.Println("runDefer Done")
}
