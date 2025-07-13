package main

import (
	"fmt"
	"sync"
)

var counter int

func increment(wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	mutex.Lock()

	counter++
	fmt.Println("Counter:", counter)
	mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup

	var mutex sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&wg, &mutex)
	}

	wg.Wait()

	fmt.Println("Final Counter Value:", counter)

}
