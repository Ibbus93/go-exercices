package main

import (
	"fmt"
)

func main() {
	var count1 int
	var count2 int = 100

	count3 := 200
	count4 := count2 + count3

	fmt.Println(count1)
	fmt.Println(count3)
	fmt.Println(count4)

	count1++

	fmt.Println(count1)
}