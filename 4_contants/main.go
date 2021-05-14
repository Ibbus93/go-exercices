package main

import (
	"fmt"
)

// const maxValue // compile error (it requires a value)

const population int = 10000

func main () {
	const name = "Cloud"

	// name = "Blah" // compile error (cannot reassign)

	if true {
		const color = "red"
		
		fmt.Println(population)
		fmt.Println(name)
	}

	// fmt.Println(color) // compile error (undefined)
}