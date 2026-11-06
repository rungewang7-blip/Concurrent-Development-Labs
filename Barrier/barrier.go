// Barrier.go Template Code
// Copyright (C) 2024 Dr. Joseph Kehoe
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Runge Wang (C00304516)
// Issues: None
//--------------------------------------------

// CLASSMATES I WORKED WITH:
// Dorian Nowacki (C00288613)

// CLASSMATES I HELPED:
// Sean McGrath
package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, wg *sync.WaitGroup, counter *int,
	theLock *sync.Mutex, sem *semaphore.Weighted,
	totalRoutines int, ctx context.Context) bool {

	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)

	theLock.Lock()
	(*counter)++

	if *counter == totalRoutines {
		// Last goroutine opens the barrier
		sem.Release(1)
		theLock.Unlock()
	} else {
		theLock.Unlock()

		// Wait until all goroutines finish Part A
		sem.Acquire(ctx, 1)

		// Allow the next goroutine to continue
		sem.Release(1)
	}

	fmt.Println("PartB", goNum)
	wg.Done()
	return true
}
func main() {
	counter := 0
	totalRoutines := 10

	var wg sync.WaitGroup
	wg.Add(totalRoutines)

	ctx := context.TODO()
	var theLock sync.Mutex

	// Create a semaphore with one permit
	sem := semaphore.NewWeighted(1)

	// Close the barrier before starting the goroutines
	sem.Acquire(ctx, 1)

	for i := range totalRoutines {
		go doStuff(i, &wg, &counter, &theLock, sem, totalRoutines, ctx)
	}

	wg.Wait()

	// Take back the final permit
	sem.Acquire(ctx, 1)
}
