package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// A buffered channel with one token acts as a turnstile.
type barrier struct {
	total      int32
	arrived    atomic.Int32
	turnstile1 chan bool
	turnstile2 chan bool
}

func newBarrier(total int32) *barrier {
	b := &barrier{
		total:      total,
		turnstile1: make(chan bool, 1),
		turnstile2: make(chan bool, 1),
	}
	// The first turnstile starts closed; the second starts open.
	b.turnstile2 <- true
	return b
}

func (b *barrier) wait() {
	if b.arrived.Add(1) == b.total {
		<-b.turnstile2 // Close the second turnstile.
		b.turnstile1 <- true
	}

	<-b.turnstile1
	b.turnstile1 <- true // Let the next goroutine through.

	if b.arrived.Add(-1) == 0 {
		<-b.turnstile1 // Close the first turnstile for the next round.
		b.turnstile2 <- true
	}

	<-b.turnstile2
	b.turnstile2 <- true
}

func main() {
	const workers = 5
	const rounds = 3
	b := newBarrier(workers)
	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func(id int) {
			defer wg.Done()
			for round := 1; round <= rounds; round++ {
				fmt.Printf("Round %d: Part A %d\n", round, id)
				b.wait()
				fmt.Printf("Round %d: Part B %d\n", round, id)
			}
		}(i)
	}
	wg.Wait()
}
