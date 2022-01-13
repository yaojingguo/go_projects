package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	for {
		for i := 1; i < 1000; i++ {
			fmt.Println(i)
		}
		time.Sleep(time.Second * 1)
	}
}
