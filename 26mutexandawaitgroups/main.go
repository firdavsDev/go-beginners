package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition")
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	var score = []int{0}
	wg.Add(3)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		defer wg.Done()
		fmt.Println("Routine 1")
		m.Lock()
		score = append(score, 1)
		m.Unlock()
	}(wg, mutex)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		defer wg.Done()
		fmt.Println("Routine 2")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
	}(wg, mutex)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		defer wg.Done()
		fmt.Println("Routine 3")
		m.Lock()
		score = append(score, 3)
		m.Unlock()
	}(wg, mutex)

	wg.Wait()
	fmt.Println(score)

}
