package main

import "fmt"

func print (id int, x int) {
	fmt.Printf("%d: x=%d\n", id, x)
}

var x int = 11

func main() {
	print(1, x)
	
	x := 2

	print(2, x)
}