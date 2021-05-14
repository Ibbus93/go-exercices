package main

import "fmt"

type player struct {
	id int
	rankn int
}

func main () {
	map1 := make(map[string]string)

	map1["key1"] = "value1"
	map1["key2"] = "value2"
	map1["key3"] = "value3"

	value1 := map1["key1"]
	fmt.Println(value1)

	// map1["key1"] = 10 // compile error (unmatched type)
	
	delete(map1, "key2")
	map1["newKey"] = "value4"

	fmt.Println(map1)

	team := map[string]player{"John": {3, 10}, "Bob": {14, 11}}

	fmt.Println(team)


}