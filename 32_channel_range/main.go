package main

import "fmt"

func squares() func() int {
	var x int

	return func() int {
		x++
		return x * x
	}
}

func main(){
	f := squares()
	squares := make(chan int, 20)

	for i := f(); i <= 100; i = f() {
		squares <- i
	} 

	// It's important to close the channel to use the range
	// You cannot still read in a open channel, it needs to be closed
	close(squares)

	for elem := range squares {
		fmt.Println(elem)
	}
}