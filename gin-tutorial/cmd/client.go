package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	for i := 0 ; i < 100; i++ {
		for j := 0; j < 50; j++ {
			call()
		}
		time.Sleep(1 * time.Second)
	}
}

func call() {
	resp, err := http.Get("http://localhost:2112/albums")
	if err != nil {
		log.Print(err)
	} else {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Print(err)
		} else {
			log.Printf("%s", body)
		}
	}
}