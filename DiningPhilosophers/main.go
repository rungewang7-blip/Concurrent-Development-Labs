// Dining Philosophers Template Code
// Author: Joseph Kehoe
// Created: 21/10/24
// GPL Licence (version not specified in the supplied Go template)
// Modified by: Runge Wang (C00304516)
package main

import (
	"fmt"
	"sync"
)

const philCount = 5
const rounds = 3

func think(index int) {
	fmt.Println("Philosopher", index, "is thinking")
}

func eat(index int) {
	fmt.Println("Philosopher", index, "is eating")
}

func getForks(index int, forks map[int]chan bool) {
	forks[index] <- true
	forks[(index+1)%philCount] <- true
}

func putForks(index int, forks map[int]chan bool) {
	<-forks[index]
	<-forks[(index+1)%philCount]
}

func doPhilStuff(index int, wg *sync.WaitGroup, forks map[int]chan bool, room chan bool) {
	defer wg.Done()
	for i := 0; i < rounds; i++ {
		think(index)
		// At most four philosophers may try to pick up forks at once.
		room <- true
		getForks(index, forks)
		eat(index)
		putForks(index, forks)
		<-room
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(philCount)

	forks := make(map[int]chan bool)
	for k := 0; k < philCount; k++ {
		forks[k] = make(chan bool, 1)
	}
	room := make(chan bool, philCount-1)
	for n := 0; n < philCount; n++ {
		go doPhilStuff(n, &wg, forks, room)
	}
	wg.Wait()
}
