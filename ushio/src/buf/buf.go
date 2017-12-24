package buf

import (
    "fmt"
    "time"
)

func GetInt(a int, c chan int) {
    for i := 1; i < 5; i++ {

        c <- a + i
        fmt.Println("GetInt:", i)
    }
    close(c)
}

func Start() {
    c := make(chan int)

    go GetInt(0, c)

    for d := range c {
        time.Sleep(time.Second * 2)
        fmt.Println("Main: ", d)
    }

}