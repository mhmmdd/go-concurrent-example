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
		for msg := range ch {
			fmt.Println(msg)
		}
		wg.Done()
	}(ch, wg)

	// send-only channel ->chan
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// send a message to channel
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
