package main

func main() {
	// With single goroutine and channels
	evenCh := make(chan int)
	oddChan := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i <= 10; i++ {
			if i%2 == 0 {
				evenCh <- i
			} else {
				oddChan <- i
			}

			if i == 10 {
				done <- true
				return
			}
		}
	}()

loop:
	for {
		select {
		case i := <-evenCh:
			println("Even loop:", i)
		case i := <-oddChan:
			println("Odd loop:", i)
		case <-done:
			println("job completed")
			break loop
		}
	}
}
