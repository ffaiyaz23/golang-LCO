package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("This is all about channels in Golang")

	// myChannel := make(chan int, 2) // buffer channel
	myChannel := make(chan int)
	wg := &sync.WaitGroup{}

	// myChannel <- 5
	// fmt.Println(<-myChannel)

	wg.Add(2)
	// RECEIVE ONLY
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myChannel
		if isChannelOpen {
			fmt.Println(val)
		}

		val1, isChannelOpen := <-myChannel
		if isChannelOpen {
			fmt.Println(val1)
		}

		wg.Done()
	}(myChannel, wg)
	// SEND ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myChannel <- 0
		close(myChannel)
		// myChannel <- 6

		wg.Done()
	}(myChannel, wg)

	wg.Wait()
}
