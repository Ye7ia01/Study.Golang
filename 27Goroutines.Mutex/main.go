package main

import (
	"fmt"
	"sync"
)

func main() {

	/*
		An Application to count even numbers in set
		using multi threading, splitting performance accross threads
		testing race conditions and Mutexes for critical sections
	*/

	/*
		run with --race (gcc compiler needed) to check race conditions
	*/

	// Initialize a struct for WaitGroup
	var wg sync.WaitGroup

	// Initaliza Mutex
	var mx sync.Mutex

	// Define Large Input Set
	largeInputSet := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}

	// Split the large set into 3 for each thread with one global counter

	// Shared Memory
	// No need for locks on initialization
	var evenNumberCounter int

	// Start Threading the heavy computation
	wg.Add(3) // Add 3 Threads for the Group
	go thread_task(&wg, &mx, largeInputSet[:10], &evenNumberCounter)
	go thread_task(&wg, &mx, largeInputSet[10:20], &evenNumberCounter)
	go thread_task(&wg, &mx, largeInputSet[20:30], &evenNumberCounter)

	wg.Wait()
	// Read Shared Memory Final Result
	fmt.Println("Res : ", evenNumberCounter)
}

// WaitGroup, Mutex are passed as pointer to use the original object in functions
// and not to use copies of the object in the functions
func thread_task(wg *sync.WaitGroup, mx *sync.Mutex, inputSet []int, evenNumberCount *int) {

	// When leaving the function set as done
	for _, input := range inputSet {
		if input%2 == 0 {

			// Apply Mutex on critical section
			// Lock until goroutine finishes
			mx.Lock()
			// Critical Section
			// Can cause race condition
			*evenNumberCount++
			// Unlock after goroutine finishes
			mx.Unlock()
		}
	}

	wg.Done()
}
