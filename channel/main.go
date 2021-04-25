package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	// create a channel
	ch := make(chan int, 1)

	wg.Add(2)

	// receive-only channel <-chan
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// receive a message from channel
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)

	// send-only channel ->chan
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// send a message to channel
		ch <- 42
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
