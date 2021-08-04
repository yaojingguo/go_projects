package main

import (
	"sync"
	"time"
)


func main() {
	var g sync.WaitGroup
	g.Add(1)
	g.Add(1)
	g.Done()
	g.Done()
	g.Wait()
}

func foo(ch chan int) {

	select {
	case <- ch:

	case <-time.After(1 * time.Second):
	}
}

