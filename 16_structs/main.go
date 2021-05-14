package main

import "fmt"

type person struct {
	firstname string
	surname string
}

type lecture struct {
	name string
	instructor person
	duration int
}

func main() {
	johnInstructor := person{"John", "Elkann"}
	points := 300

	lectures := []lecture{
		{"Structs", johnInstructor, points},
		{"Pointers", johnInstructor, points},
		{"Functions", johnInstructor, points},
	}

	for _, lecture := range lectures {
		name := lecture.name
		instructor := fmt.Sprintf("%s %s", lecture.instructor.firstname, lecture.instructor.surname)
		duration := lecture.duration

		fmt.Printf("Lecture: '%s', Author: '%s', Duration: %d secs\n", name, instructor, duration)
	}

}