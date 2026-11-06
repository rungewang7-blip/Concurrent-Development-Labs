package main

import (
	"sync"
	"testing"
)

func TestEveryEventConsumedOnce(t *testing.T) {
	const eventCount = 20
	buffer := make(chan event, 5)
	consumed := make(chan event)
	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go consumer(buffer, consumed, &wg)
	}
	go producer(buffer, eventCount)
	go func() {
		wg.Wait()
		close(consumed)
	}()

	counts := make([]int, eventCount)
	for item := range consumed {
		if item.id < 0 || item.id >= eventCount {
			t.Errorf("unexpected event %d", item.id)
			continue
		}
		counts[item.id]++
	}
	for id, count := range counts {
		if count != 1 {
			t.Errorf("event %d consumed %d times; want 1", id, count)
		}
	}
}
