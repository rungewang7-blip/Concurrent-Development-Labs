package main

import (
	"fmt"
	"sync"
)

func main() {
	const workers = 100
	const updates = 1000
	total := 0

	// One token permits only one goroutine to update total at a time.
	semaphore := make(chan bool, 1)
	semaphore <- true
	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < updates; j++ {
				<-semaphore
				total++
				semaphore <- true
			}
		}()
	}

	wg.Wait()
	fmt.Println("Total:", total)
}
