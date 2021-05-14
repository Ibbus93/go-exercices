package main

import "fmt"

func extend(val string) func() string {
	i := 0

	// this is a closure
	// The ability to reference variables declared outside of itself
	// In this case, it's the var i
	return func() string {
		i++
		return val[:i]
	}
}

func main() {
	ca := "cloudacademy"
	word := extend(ca)

	for i := 0; i < len(ca); i++ {
		fmt.Println(word())
	}
}