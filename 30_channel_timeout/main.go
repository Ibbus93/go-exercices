package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	waitTime := 2

	go func(channel chan string) {
		seconds := 5
		time.Sleep(time.Duration(seconds) * time.Second)
		channel <- "cloud"
	}(channel)

	select {
	case msg1 := <-channel:
		fmt.Println(msg1)
	case <-time.After(time.Duration(waitTime) * time.Second):
		fmt.Println("timeout")
	}

}