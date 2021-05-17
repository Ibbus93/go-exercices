// Goroutines can be conceptually like lightweight threads
// They are scheduled and handled  by the Go runtime
// They are execute at the same time together
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pause() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(3) // n between 0 and 3
	time.Sleep(time.Duration(n) * time.Second)
}

func doSomething(msg string) {
	pause()
	fmt.Println(msg)	
}

func main() {
	doSomething("sync1")
	
	go doSomething("async 1")
	go doSomething("async 2")
	go doSomething("async 3")
	
	doSomething("sync2")

	time.Sleep(time.Second * 10)
}