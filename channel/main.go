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
		// if the channel is open then ok is to be true
		if msg, ok := <-ch; ok {
			fmt.Println(msg, ok)
		}
		wg.Done()
	}(ch, wg)

	// send-only channel ->chan
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// send a message to channel
		close(ch)
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
