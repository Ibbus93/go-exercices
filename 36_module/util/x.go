package util

const X int = 100 // exported
const x int = 200 // unexported

func GetX() int { // exported
	return x
}

func getX() int { // unexported
	return x
}

func GetXY() int {
	return x * getY() // visible bcs they are in the same package
}
