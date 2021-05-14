package main

import "fmt"


type person struct {
	firstname string
	surname string
	age int
}

// This is a method
// The (p *person) reprents the receiver of the method
// Using a pointer, we allow to modify the parameters
// and we avoid copying the value on each method core
func (p *person) fullname() string {
	return p.firstname + " " + p.surname
}

func (p *person) canDrive() bool {
	if p.age >= 18 {
		return true
	}
	return false
}

func (p *person) updateAge(newAge int) {
	p.age = newAge
}

func (p* person) incrementAge() {
	p.age++
}

func main() {
	person1 := person{"John", "Elkann", 50}
	person2 := person{"Reato", "Di", 1}
	
	fmt.Printf("%s can drive: %t (age: %d)\n", person1.fullname(), person1.canDrive(), person1.age)
	fmt.Printf("%s can drive: %t (age: %d)\n", person2.fullname(), person2.canDrive(), person2.age)
	
	person2.updateAge(person2.age + 16)
	fmt.Printf("%s can drive: %t (age: %d)\n", person2.fullname(), person2.canDrive(), person2.age)
	
	person2.incrementAge()
	fmt.Printf("%s can drive: %t (age: %d)\n", person2.fullname(), person2.canDrive(), person2.age)
}