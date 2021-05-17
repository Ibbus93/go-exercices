// When using channels as function parameters, as you often will, 
// by default can send and receive within the function. 

// To provide additional safety at compile time, 
// channel function parameters can be defined with a direction. 
// That is, they can be defined to be read-only or write-only. 
package main

import (
	"fmt"
	"time"
)

// function that locks the channel only in writing
func in(channel chan <- string, msg string) {
	channel <- msg
}

// functions that locks the channel only in reading
func out(channel <- chan string) {
	for {
		fmt.Println(<-channel)
	}
}

func main() {
	channel := make(chan string, 1)

	go out(channel)

	for i := 0; i < 10; i++ {
		in(channel, fmt.Sprintf("CloudAcademy - %d", i))
	}

	time.Sleep(time.Second * 10)
}