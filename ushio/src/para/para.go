package para

import (
	"fmt"
	"runtime"
)

func GetInt(a int, c chan int) {
    c <- a
    fmt.Println("GetInt", a)
}

func Start() {
    numCPU := runtime.GOMAXPROCS(0)
    fmt.Println("NUMCPU:", numCPU)
    c := make(chan int, numCPU)

    for i := 0; i < numCPU-1; i++ {
        go GetInt(i*10, c)
    }
    result03, result02, result01 := <-c, <-c, <-c
    fmt.Println("Main:", result01)
    fmt.Println("Main:", result02)
    fmt.Println("Main:", result03)
}