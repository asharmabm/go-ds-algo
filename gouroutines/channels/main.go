package main

import (
	"fmt"
)

func printA(ch chan string) {
	//defer wg.Done()
	//fmt.Println("Hello from printA")
	ch <- "Hello from printA"
}

func printB(ch chan string) {
	//defer wg.Done()
	//fmt.Println("Hello from printB")
	ch <- "Hello from printB"
}

func main() {
	//var wg sync.WaitGroup

	chA := make(chan string)
	chB := make(chan string)

	//wg.Add(1)
	go printA(chA)
	fmt.Println(<-chA)

	go func() {
		chA <- "dfg"
	}()

	fmt.Println(<-chA)

	//wg.Wait()

	//wg.Add(1)
	go printB(chB)
	fmt.Println(<-chB)

	//wg.Wait()

	fmt.Println("Done main function")
}
