package main

import (
	"fmt"
	"sync"
)

func main() {
	evenCh := make(chan int)
	oddCh := make(chan int)
	var wg sync.WaitGroup

	// Create a channel to coordinate printing in order
	doneCh := make(chan int, 10)

	// Start goroutines for each number
	for i := 0; i < 10; i++ {
		wg.Add(2)

		val := i
		if val%2 == 0 {
			go func(n int) {
				evenCh <- n
				wg.Done()
			}(val)
			go func(n int) {
				v := <-evenCh
				doneCh <- n // Send the number to indicate it's done
				fmt.Println("Even:", v)
				wg.Done()
			}(val)
		} else {
			go func(n int) {
				oddCh <- n
				wg.Done()
			}(val)
			go func(n int) {
				v := <-oddCh
				doneCh <- n // Send the number to indicate it's done
				fmt.Println("Odd:", v)
				wg.Done()
			}(val)
		}
	}

	// Process results in order
	for i := 0; i < 10; i++ {
		// Wait for each number
		for {
			num := <-doneCh
			if num == i {
				break
			}
			// Put back numbers that aren't in sequence yet
			doneCh <- num
		}
	}

	wg.Wait()
}
