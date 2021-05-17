// When designing Go applications involving goroutines which read and write to multiple channels,
// it should seriously considered to use the select and case statements
// to simplify the management and readability of waits across multiple channels
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pause() {
	n := rand.Intn(5) // n between 0 and 5
	time.Sleep(time.Duration(n) * time.Second)
}

func func1(c chan <- string) {
	for {
		pause()
		c <- "cloud"
	}
}

func func2(c chan <- string) {
	for {
		pause()
		c <- "academy"
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	chan1 := make(chan string)
	chan2 := make(chan string)

	go func1(chan1)
	go func2(chan2)

	for {
		select {
		case msg1 := <- chan1:
			fmt.Println("chan1", msg1)
		case msg2 := <- chan2:
			fmt.Println("chan2", msg2)	
		}
	}
}