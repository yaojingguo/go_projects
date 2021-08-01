package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func(c chan int) {
		no := <-c
		fmt.Printf("no: %d\n", no)
	}(ch)
	ch <- 10
	time.Sleep(1 * time.Second)
}
