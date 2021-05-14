package main

import "fmt"
import "strings"

type upstring string

func (msg upstring) up() string {
	s := string(msg)
	return strings.ToUpper(s)
}

func main() {
	message := upstring("clean")
	fmt.Println(message.up())
}