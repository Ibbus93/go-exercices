package main

import "fmt"

// slice allows either to expand or to restrict the length of the array at runtime
func main() {
	arr1 := [6]int{1, 3, 5, 7, 11, 13}
	slice1 := []int{1, 3, 5, 7, 11, 13} // first way to perform slice is not inserting the length
	slice2 := slice1[1:2]

	// make takes the type of the slice, the length of the slice and the capacity of the slice
	var slice3 = make([]int, 2, 3)

	fmt.Println(arr1)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	slice3 = slice1[1:4]

	fmt.Println(slice3)
	fmt.Println(len(slice3))

	// append adds elements to slice1
	slice1 = append(slice1, 200, 300, 400)
	fmt.Println(slice1)

	// the three points unpack the array
	slice2 = append(slice2, []int{7,8,9}...)
	fmt.Println(slice2)

	// copy makes a copy of slice1 in copyslice
	// make here creats an array with the same length of slice1
	copyslice := make([]int, len(slice1))
	copy(copyslice, slice1)
	fmt.Println(copyslice)

}