package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mypkg"
	"os"
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
	Id     ID
	Pri    PRIORITY
	Detail string
	Done   bool
}

func task(id ID, pri PRIORITY) {
	fmt.Println(string(id) + "-" + string(pri))

	var task Task = Task{
		Id:     id,
		Pri:    pri,
		Detail: "myTask",
		Done:   false,
	}

	b, err := json.Marshal(task)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("json:")
	fmt.Println(string(b))

	file, err := os.Create("./file.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	message := []byte("hello world\n")
	message2 := "hello world2\n"

	_, err = file.Write(message)
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(message2)
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fprint(file, "hello world3\n")
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadFile("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)

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
