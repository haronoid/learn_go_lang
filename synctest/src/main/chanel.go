package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	urls := []string{
		"http://example.com",
		"http://example.net",
		"http://example.org",
	}
	statusChan := getStatus(urls)

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-statusChan)
	}
}

func getStatus(urls []string) <-chan string {
	statusChan := make(chan string)
	for _, url := range urls {

		//gorutine execute
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			statusChan <- res.Status
		}(url)
	}
	return statusChan
}
