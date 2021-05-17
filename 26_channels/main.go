package main

import "fmt"

// Channels are communication objects used by goroutines to communicate and share data with each other. 
// One goroutine will write messages into the channel
// while another goroutine will read the same messages out of the channel. 

func main() {
	msgChan := make(chan string)

	go func(){
		msgChan <- "Cloud"
		msgChan <- "Academy"
		msgChan <- "2020"
	}()

	msg1 := <- msgChan
	msg2 := <- msgChan
	msg3 := <- msgChan

	fmt.Println(msg1, msg2, msg3)
}