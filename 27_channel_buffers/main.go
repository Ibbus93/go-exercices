package main

import (
	"time"
)

// channels can be buffered or unbuffered
// to be buffered, it is enough to add the size parameter to the make func

// An unbuffered channel causes the sender to block immediately 
// after writing its message into the channel, 
// and until the receiver has read it back off the channel. 
// In a buffered channel scenario, the sender blocks only once the buffer 
// has filled up and waits until space becomes available before unblocking. 
// Put another way, when the channel is full, the sender blocks and waits 
// until another goroutine reads off the channel making room by receiving. 
func main() {
	size := 3

	var buffChan = make(chan int, size)

	// reader
	go func(){
		for {
			_ = <- buffChan
			time.Sleep(time.Second * 3)
		}
	}()

	// writer
	writer := func() {
		for i := 1 ; i <= 10; i++ {
			buffChan <- i
			println(i)
		}
	}

	writer()
}