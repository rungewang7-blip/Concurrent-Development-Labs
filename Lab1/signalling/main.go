package main

import (
	"fmt"
	"sync"
)

func main() {
	// No token means the second goroutine must wait for the signal.
	signal := make(chan bool, 1)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("First task finished")
		signal <- true
	}()

	go func() {
		defer wg.Done()
		<-signal
		fmt.Println("Second task starts")
	}()

	wg.Wait()
}
