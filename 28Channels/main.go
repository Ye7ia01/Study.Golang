package main

import (
	"fmt"
	"sync"
)

func main() {

	/*
		Rule 1 : To send from one end of a channel, the other end has to be receiving from the channel
		& vice versa , to receive , the other end has to be sending
		otherwise a deadlock , or an error will happen

		Channels are queue data structures, First to be sent is first received

		Channel can be buffered or unbuffered
		Unbuffered Channels : make(chan T) can be used for many sends ,
			but once a message is sent it has to be received before continuing
			i.e. a send blocks until receave is ready
		Buffered channels :
			make(chan T, 10), allow for a buffering messages while not necessary reading it right away
			i.e you can send 10 messages (non-blocking) and can be received and read any time
	*/

	// Unbuffered channel
	// IMP : can send many values
	// 	but one send has to be recieved from the other end
	// 	any other behavior (no receiving end) will cause errors
	// Send only channel
	channel := make(chan int)
	// Read only channel
	// rChannel := make(<-chan int)
	// Define Waiting Group
	var wg sync.WaitGroup

	wg.Add(2)
	go thread_send(channel, &wg)
	go thread_rec(channel, &wg)

	wg.Wait()
	fmt.Println("Threads Executed, Terminating ..")
}

// Thread (goroutine) with inter-thread communication
// Thread can receive only in channel <- chan
func thread_rec(ch <-chan int, wg *sync.WaitGroup) {

	// Some Processing

	// All what the job / thread is doing is receiving
	for v := range ch {
		fmt.Println("Receive new value on channel : ", v)
	}
	// Some Processing
	wg.Done()
}

// Thread (gorouting) with inter-thread communication
// Thread can send only in channel chan <-
func thread_send(ch chan<- int, wg *sync.WaitGroup) {

	// Some Processing

	// send a message in channel
	// Will give errors if receiving end was not configured to receive
	for i := 0; i < 10; i++ {
		ch <- i
	}

	// Some Processing
	// Close channel
	close(ch)

	wg.Done()
}
