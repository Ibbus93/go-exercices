package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

import m "math"

func main() {
	name := "cloud"

	fmt.Println(strings.ToUpper(name))
	fmt.Println(uuid.New())

	f, _ := strconv.ParseFloat("3.1415", 64)
	fmt.Println(f)

	fmt.Println(m.Round(f))
}