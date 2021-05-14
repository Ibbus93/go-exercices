package main

import "fmt"
import "errors"

func divide(num1 int, num2 int) (int, error) {
	if num2 == 0 {
		return 0, errors.New("division by zero not allowed")
	} else {
		return (num1 / num2), nil
	}
}

func main() {
	if result, err := divide(10, 0); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}