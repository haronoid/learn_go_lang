package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	wait := new(sync.WaitGroup)
	urls := []string{
		"http://example.com",
		"http://example.net",
		"http://example.org",
	}

	for _, url := range urls {

		// add WaitGroup
		wait.Add(1)

		//gorutine execute
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			fmt.Println(url, res.Status)
			// add endWaitGroup
			wait.Done()
		}(url)

		//		res, err := http.Get(url)
		//		if err != nil {
		//			log.Fatal(err)
		//		}
		//		defer res.Body.Close()
		//		fmt.Println(url, res.Status)

	}
	// wait all WaitGroup Done
	wait.Wait()
	//	time.Sleep(time.Second)
}
