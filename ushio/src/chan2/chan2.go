package chan2

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetContent(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	s := string(body[:])

	if err != nil {
		fmt.Println(err)
		return
	}
	c <- s
}

func Start() {
	fmt.Println("Hello, playground")
	c := make(chan string)
	go GetContent("https://www.bing.com/", c)
	go GetContent("https://www.yahoo.co.jp", c)
	result01, result02 := <-c, <-c
	fmt.Println(result01)
	fmt.Println("********-----------")
	fmt.Println(result02)

}
