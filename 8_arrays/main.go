package main

import "fmt"

// Array cannot be expanded, the dimension is fixed
func main() {
	var arr1[4] int // init with 0s

	arr2 := [...]int{3, 1, 5, 10, 100}

	fmt.Println(arr1) // 4
	fmt.Println(arr2) // 5

	fmt.Println(len(arr1)) // 4
	fmt.Println(len(arr2)) // 5

	// underscore receive and discard index value
	for _, value := range arr2 {
		fmt.Println(value)
	}

	for index, value := range arr1 {
		fmt.Println(index,":", value)
	}

	arr1[0] = 10
	// arr1[4] = 1 compile error (out of bound)

	fmt.Println(arr1[0])

	// multidimensional array
	arr3 := [2][2]string{
		{"a1", "b1"},
		{"a2", "b2"}, // trailing comma required
	}

	fmt.Println(arr3)
}