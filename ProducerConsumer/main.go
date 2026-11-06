package main

import (
	"fmt"
	"sync"
)

// event is an item produced and then consumed.
type event struct {
	id int
}

// The buffered channel is the shared buffer.
func producer(buffer chan<- event, count int) {
	for i := 0; i < count; i++ {
		buffer <- event{id: i}
	}
	close(buffer) // No more events will be produced.
}

func consumer(buffer <-chan event, consumed chan<- event, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range buffer {
		consumed <- item
	}
}

func main() {
	const eventCount = 20
	const bufferSize = 5
	const consumerCount = 2

	buffer := make(chan event, bufferSize)
	consumed := make(chan event)
	var wg sync.WaitGroup
	wg.Add(consumerCount)
	for i := 0; i < consumerCount; i++ {
		go consumer(buffer, consumed, &wg)
	}
	go producer(buffer, eventCount)
	go func() {
		wg.Wait()
		close(consumed)
	}()
	for item := range consumed {
		fmt.Println("Consumed event", item.id)
	}
}
