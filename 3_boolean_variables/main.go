package main

import (
	"fmt"
)

func main() {
	var isBig bool // default to false
	var isFast, hasFourWheels = false, true

	fmt.Println(isBig)
	fmt.Println(isFast)
	fmt.Println(hasFourWheels)

	if !isFast {
		fmt.Println("so slow Bottas, kek")
	} else {
		fmt.Println("So fast, Ham")
	}

	if hasFourWheels {
		fmt.Println("It's a car, wow!")
	}
}