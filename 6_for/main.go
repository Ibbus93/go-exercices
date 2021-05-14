package main

import "fmt"

func main() {
	x:= 0
	y := 0

	// for is the only available cycle in Go
	for {
		if x++; x > 2 {
			fmt.Println("x", x)
			break
		}
	}

	for y < 3 {
		fmt.Println("y", y)
		y++
	}

	for z := 0; z < 10; z++ {
		if z < 8 {
			// this blocks the execution of the iteration and goes to the next one
			continue
		}
		fmt.Println("z", z)
	}
}