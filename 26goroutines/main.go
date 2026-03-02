package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mutx sync.Mutex

func main() {

	score := make([]int, 8)

	wg.Add(3)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// Shared Resource
		// Can Cause Race conditions
		mutx.Lock()
		score = append(score, 1)
		mutx.Unlock()
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// Shared Resource
		// Can Cause Race conditions
		mutx.Lock()
		score = append(score, 2)
		mutx.Unlock()

	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// Shared Resource
		// Can Cause Race conditions
		mutx.Lock()
		score = append(score, 3)
		mutx.Unlock()

	}(&wg)

	wg.Wait()
	mutx.Lock()
	fmt.Println(score)
	mutx.Unlock()
}

// import (
// 	"fmt"
// 	"net/http"
// 	"sync"
// )

// var wg sync.WaitGroup
// var mutex sync.Mutex

// func main() {

// 	fmt.Println("Go Routines using WaitGroup & recommneded behavior from the documentation")

// 	domains := []string{
// 		"https://www.google.com",
// 		"http://www.golang.org/",
// 	}

// 	var score int = 0

// 	for _, domain := range domains {
// 		wg.Add(5) // Add delta +1 . i.e. increment routines pool
// 		/*
// 			Better approach based on documentation
// 			wg.Go(func ())
// 			But It requires the function definition to be written inside (not readable)

// 			Another note, better to specify wg.Add() before the firing up of the thread
// 		*/
// 		go GetRequest(domain, &score) // Fireup go routine
// 		go GetRequest(domain, &score) // Fireup go routine
// 		go GetRequest(domain, &score) // Fireup go routine
// 		go GetRequest(domain, &score) // Fireup go routine
// 		go GetRequest(domain, &score) // Fireup go routine
// 	}

// 	wg.Wait() // wait for all threads (routines) to finish before exit
// 	fmt.Println(score)
// }

// func GetRequest(s string, raceCondition *int) {
// 	defer wg.Done() // Set Thread as done once finished func execution
// 	res, err := http.Get(s)
// 	if err != nil {
// 		fmt.Println("Failed to get domain ", s)
// 	}
// 	// content, err := io.ReadAll(res.Body)
// 	// if err != nil {
// 	// fmt.Println("Failed to get domain ", s)

// 	// }
// 	// Lock Before Shared Resource
// 	mutex.Lock()
// 	// Shared Resource
// 	*raceCondition++
// 	// Unlock
// 	mutex.Unlock()
// 	fmt.Printf("Success : Domain %v has code %v \n", s, res.StatusCode)
// }
