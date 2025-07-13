package main

import "fmt"

func main() {
	evenCh := make(chan int)
	oddChan := make(chan int)
	done := make(chan bool)

	go func() {
		for {
			i := <-evenCh
			fmt.Println("Even loop:", i)
			if i%2 == 0 {
				println("Even:", i)
			}
			if i == 10 {
				done <- true
				return
			}
			oddChan <- i + 1
		}
	}()

	go func() {
		for {
			i := <-oddChan
			fmt.Println("Odd loop:", i)
			if i%2 != 0 {
				println("Odd:", i)
			}
			evenCh <- i + 1
		}
	}()

	evenCh <- 0

	<-done

	close(evenCh)
	close(oddChan)
	close(done)
}
