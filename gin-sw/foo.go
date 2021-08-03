package main

import (
	"fmt"
	"time"
)

func main() {
	// ch := make(chan int, 1)
	var ch chan int
	fmt.Printf("channel: %v\n", ch)
	// go foo(ch)
	go func() {
		ch <- 100
	}()
	time.Sleep(1 * time.Second)
}

// func foo(ch chan int) {
//   v, ok := <-ch
//   if !ok {
//     fmt.Println("receive message error")
//   } else {
//     fmt.Printf("v: %d\n", v)
//   }
// }
