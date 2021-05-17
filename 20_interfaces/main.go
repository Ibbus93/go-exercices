package main

import "fmt"
import "strings"

// interfaces are collections of methods
type device interface {
	turnOn() string
	update(version float32)
}

type iphone struct {
	name string
	model string
	version float32
}

type imac struct {
	name string
	model string	
	version float32
}

type company struct {
	name string
}

func (phone iphone) turnOn() string {
	return "iOS starting up..."
}

func (mac imac) turnOn() string {
	return "macOS starting up..."
}

func (phone *iphone) update(version float32) {
	phone.version = version
}

func (mac *imac) update(version float32) {
	mac.version = version
}

func main() {
	dev1 := iphone{"iPhone", "11 Pro", 13.1}
	dev2 := imac{"iMac", "27 5k Retina", 10.15}

	// devices := []device{dev1, dev2}
	// pointer required bcs methods takes pointers
	devices := []device{&dev1, &dev2}

	for _, dev := range devices {
		if strings.Contains(dev.turnOn(), "iOS") {
			dev.update(14.0)
		} else if strings.Contains(dev.turnOn(), "macOS") {
			dev.update(11.0)
		}
	}

	fmt.Println(dev1)
	fmt.Println(dev2)

	// empty interfaces
	var a, b, c, d interface{}

	a = 42
	b = "blah"
	c = true
	d = company{"cloud"}

	func(list ...interface{}){
		for _, v := range list {
			fmt.Printf("%v, %T\n", v, v)
		}
	}(a, b, c, d)

	// type assertion
	c1 := d.(company) // d.(type) asserts that d is of type company and is not nil, else it panics
	fmt.Println(c1)

	if c2, ok := c.(company); ok {
		fmt.Println(c2, ok)
	} else {
		fmt.Printf("Expected company type, found: %T\n", c)
	}

	// n := d.(int); n++ panic
}