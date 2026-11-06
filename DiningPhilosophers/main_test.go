package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestForksAreExclusive(t *testing.T) {
	forks := make(map[int]chan bool)
	for i := 0; i < philCount; i++ {
		forks[i] = make(chan bool, 1)
	}
	room := make(chan bool, philCount-1)
	var held [philCount]atomic.Int32
	var conflicts atomic.Int32
	var wg sync.WaitGroup
	wg.Add(philCount)
	for i := 0; i < philCount; i++ {
		go func(id int) {
			defer wg.Done()
			for n := 0; n < 100; n++ {
				room <- true
				getForks(id, forks)
				left, right := id, (id+1)%philCount
				leftCount := held[left].Add(1)
				rightCount := held[right].Add(1)
				if leftCount != 1 || rightCount != 1 {
					conflicts.Add(1)
				}
				held[left].Add(-1)
				held[right].Add(-1)
				putForks(id, forks)
				<-room
			}
		}(i)
	}
	wg.Wait()
	if conflicts.Load() != 0 {
		t.Fatal("a fork was held by more than one philosopher")
	}
}
