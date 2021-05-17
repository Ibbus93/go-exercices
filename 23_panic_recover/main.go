package main

import "fmt"

// Custom panic can be thrown
// when something unexceptional and unexpected has happened
// If not captured, a panic will cause the program to be exited with non-zero exit code

// A panic can be recovered iff a inbuilt defer function is present

func system() int {
	fmt.Println("system started...")

	defer func(msg string){
		if r := recover(); r != nil {
			fmt.Println("recovered")		
		}
		fmt.Println(msg)
	}("blah")

	var data[] int
	// data := [1]int{0} // this avoids the crash
	var x = data[0] // causes runtime panic
	x++

	fmt.Println("system finished!")

	return 1
}

func main() {
	data := system()
	fmt.Println(data)

	panic("die!")
}