package main

import "fmt"

func main() {
	num1 := 100
	num2 := 200
	str1 := "blah"

	var ptr1 *int = &num1
	// var ptr1 *int = &str1

	// If not init, the pointer takes nil
	var ptr3 *int
	fmt.Println(ptr3)

	fmt.Printf("mem address of num1 is %p\n", &num1)
	fmt.Printf("mem address of num2 is %p\n", &num2)
	fmt.Printf("mem address of str1 is %p\n", &str1)

	fmt.Printf("ptr1 points to mem address %p\n", ptr1)
	
	*ptr1 = 101
	fmt.Println(num1)
	
	ptr1 = &num2
	fmt.Printf("ptr1 points to mem address %p\n", ptr1)
	fmt.Println(*ptr1)
	
	// this function creates a new pointer
	ptr2 := new(int)
	ptr2 = ptr1

	fmt.Println(*ptr2)	
}