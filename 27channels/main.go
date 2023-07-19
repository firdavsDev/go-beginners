package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Go")

	myCh := make(chan int, 1)

	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-ch
		fmt.Println(val)
		fmt.Println(val)
		defer wg.Done()
		//fmt.Println(<-ch)
		//fmt.Println(<-ch)
		fmt.Println("Channel Open", isChannelOpen)

	}(myCh, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		defer close(ch)

		ch <- 5
		ch <- 6
		defer wg.Done()
	}(myCh, wg)

	wg.Wait()
}
