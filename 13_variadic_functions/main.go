package main

import "fmt"

// the variadic signature must be in the last parameter
// It is passed as a slice
func displayCount(id int, letters ...string) {
	count := len(letters)

	// count := 0
	// for range letters {
	// 	count++
	// }

	fmt.Printf("%d - %d - %T - %s\n", id, count, letters, letters)
}

func main() {
	displayCount(1, "c", "l", "o", "u", "d")
	displayCount(2, "k", "e", "k")

	cloud := []string{"d", "e", "v", "o", "p", "s"}

	displayCount(3, cloud...)
}