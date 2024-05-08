package main

import (
	"fmt"
	"sync"
)

var counter int

func incrementCounter(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		counter++
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go incrementCounter(&wg)
	go incrementCounter(&wg)

	wg.Wait()

	fmt.Printf("Final counter value (without synchronization): %d\n", counter)
}
