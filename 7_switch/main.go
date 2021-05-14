package main

import "fmt"

// break keyword not required in switch
func main() {
	switch signal := 1; signal {
		case 0:
			fmt.Println("red")
		case 1:
			fmt.Println("orange")
		case 2:
			fmt.Println("green")
	}

	score := 76

	switch {
		case score <= 25:
			fmt.Println("beginner")
		case score <= 75:
			fmt.Println("pro")
		default: 
			fmt.Println("expert")
	}
}