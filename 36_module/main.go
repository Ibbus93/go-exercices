package main

import (
	"fmt"
	"github.com/cloudacademy/go/mod/demo/util"
	"github.com/cloudacademy/go/mod/demo/math"
)

func main() {
	fmt.Println(util.X)
	// fmt.Println(util.x) // compile error (unexported)
	
	fmt.Println(util.GetX())
	// fmt.Println(util.getX()) // compile error (unexported)

	fmt.Println(util.GetXY())

	x := util.GetX()
	y := util.GetY()
	sum := math.Add(x, y)
	fmt.Println(math.GetUuid())
	fmt.Println("sum:", sum)
}