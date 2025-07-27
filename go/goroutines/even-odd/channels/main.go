package main

import (
	"fmt"
)

func main() {
	done := make(chan bool)
	evenCh := make(chan bool)
	oddCh := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			<-evenCh
			fmt.Println("Even loop:", i)
			if i%2 == 0 {
				fmt.Println("Even:", i)
			}
			oddCh <- true
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			<-oddCh
			fmt.Println("Odd loop:", i)
			if i%2 != 0 {
				fmt.Println("Odd:", i)
			}

			if i == 9 {
				done <- true
				return
			}
			evenCh <- true
		}
	}()

	evenCh <- true
	<-done
}
