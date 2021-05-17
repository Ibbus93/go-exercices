package main

import "fmt"

func doSomething(msg string) {
	fmt.Println(msg)
}

// A defer function is good to doing cleanups
// If multiple defer are called, they are execute in LIFO order 
func system() int {
	fmt.Println("system started...")

	defer doSomething("cleanup")
	defer doSomething("stop")

	fmt.Println("system finished!")
	return 1
}

func main(){
	data := system()

	fmt.Println(data)
}