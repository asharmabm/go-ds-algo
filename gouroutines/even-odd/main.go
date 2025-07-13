package main

import (
	"fmt"
)

func main() {
	oddChan := make(chan bool)
	evenChan := make(chan bool)
	done := make(chan bool)

	// Odd number goroutine
	go func() {
		for i := 1; i <= 10; i += 2 {
			<-oddChan        // wait for signal to print
			fmt.Println(i)   // print odd
			evenChan <- true // signal even goroutine
		}
	}()

	// Even number goroutine
	go func() {
		for i := 2; i <= 10; i += 2 {
			<-evenChan     // wait for signal to print
			fmt.Println(i) // print even
			if i == 10 {
				done <- true // we're done after 10
				return
			}
			oddChan <- true // signal odd goroutine
		}
	}()

	// Start the sequence by kicking off the odd goroutine
	oddChan <- true

	// Wait for completion
	<-done
}
