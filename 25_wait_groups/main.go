// The WaitGroup manages to wait
/// until all goroutines have completed

package main

import (
	"fmt"
	"sync"
)

func main () {
	var wg sync.WaitGroup

	wg.Add(1)
	go func(msg string) {
		defer wg.Done()
		fmt.Println(msg)
	}("cloud")

	wg.Add(1)
	go func(msg string) {
		defer wg.Done()
		fmt.Println(msg)
	}("academy")

	wg.Wait()
}