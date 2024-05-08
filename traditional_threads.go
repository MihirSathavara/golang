package main

import (
	"fmt"
	"sync"
)

func computePrimes(start, end int) {
	// Implementation for computing prime numbers
	// ...
}

func main() {
	const numGoroutines = 10
	const maxNumber = 1000000

	// Using goroutines
	var wg sync.WaitGroup
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			computePrimes(i*(maxNumber/numGoroutines), (i+1)*(maxNumber/numGoroutines))
		}()
	}
	wg.Wait()

	// Using traditional threads (if needed)
	// ...

	fmt.Println("Done!")
}
