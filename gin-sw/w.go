package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type httpPkg struct{}

func (httpPkg) Get(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	// robots, err := io.ReadAll(res.Body)
	// res.Body.Close()
	// if err != nil {
	//   log.Fatal(err)
	// }
	fmt.Printf("%s\n", res.Status)
}

var httpHandler httpPkg

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.baidu.com/",
		"https://httpbin.org/get",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			httpHandler.Get(url)
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}
