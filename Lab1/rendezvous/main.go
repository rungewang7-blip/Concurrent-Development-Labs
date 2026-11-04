package main

import (
	"fmt"
	"sync"
)

func main() {
	firstArrived := make(chan bool, 1)
	secondArrived := make(chan bool, 1)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("A1")
		firstArrived <- true
		<-secondArrived
		fmt.Println("A2")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("B1")
		secondArrived <- true
		<-firstArrived
		fmt.Println("B2")
	}()

	wg.Wait()
}
