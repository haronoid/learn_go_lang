package main

import (
	"fmt"
	"mypkg"
)

func main() {
	fmt.Println(mypkg.Message)
	var s []string
	s = append(s, "a")
	s = append(s, "b")
	s = append(s, " abc")

	fmt.Println(s[1:])

	if s[1] == "b" {
		fmt.Println("OK")
	}

	fmt.Println(sum(1, 2, 4, 5, 6))

	var id ID = 1
	var pri PRIORITY = 2
	task(id, pri)

	Print(id)
}

func sum(num ...int) (result int) {
	for _, n := range num {
		result += n
	}
	return
}

type ID int
type PRIORITY int
type Task struct {
	id     ID
	pri    PRIORITY
	detail string
	done   bool
}

func task(id ID, pri PRIORITY) {
	fmt.Println(string(id) + "-" + string(pri))

	var task Task = Task{
		id:     id,
		pri:    pri,
		detail: "myTask",
		done:   false,
	}

	fmt.Println(task)
}

func Print(value interface{}) {

	_, ok := value.(int)
	fmt.Println(ok)

	switch v := value.(type) {
	case string:
		fmt.Printf("string type is : %s\n", v)
	case int:
		fmt.Printf("int type is : %d\n", v)
	case ID:
		fmt.Printf("ID type is : %d\n", v)
	}
}
