// To start it: go run main.go
package main

import (
	"fmt"
)

// Global var in the entire app
// Bcs they are declared outside the main
// And bcs their name begins with a capital letter

// If declared outside main and the name begins with the capital letter
// Than the var is exported all outside
var (
	Name1 string = "global1"
	Name2 string = "global2"
)

// syntax to declare more variables with same type
var var1, var2 string = "local1", "local2"

func main() {
	var name1 string	
	
	// fast way to declare and init a var
	name2 := "Luffy"
	name3 := "KekOps"
	name4 := name2 + " " + name3
	
	fmt.Println(name1)
	fmt.Println("Hello kek", Name1, Name2)
	fmt.Println(name4)

	// string interpolation
	fmt.Printf("%v -- %v\n", name2, name3)
}