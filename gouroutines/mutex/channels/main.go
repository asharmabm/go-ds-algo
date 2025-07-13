package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
)

func increment(ch chan bool) {
	mutex.Lock()
	ch <- true
	counter++
	fmt.Println("Counter:", counter)
	mutex.Unlock()
}

func main() {
	ch := make(chan bool, 5)
	for i := 0; i < 10; i++ {
		go increment(ch)
		<-ch
	}

	fmt.Println("Final Counter Value:", counter)

}
