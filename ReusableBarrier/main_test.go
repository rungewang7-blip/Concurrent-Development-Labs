package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestBarrierAcrossRounds(t *testing.T) {
	const workers = 5
	const rounds = 500
	b := newBarrier(workers)
	var arrivals [rounds]atomic.Int32
	var failures atomic.Int32
	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for round := 0; round < rounds; round++ {
				arrivals[round].Add(1)
				b.wait()
				if arrivals[round].Load() != workers {
					failures.Add(1)
				}
			}
		}()
	}
	wg.Wait()
	if failures.Load() != 0 {
		t.Fatal("a goroutine passed before everyone arrived")
	}
}
