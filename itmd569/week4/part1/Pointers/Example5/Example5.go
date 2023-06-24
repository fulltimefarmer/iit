package main

import "fmt"

// Point is a structure with x, y coordinates.
type Point struct {
	x int
	y int
}

// ChangePointByValue modifies a copy of the provided point.
// As Go copies the value, the original point remains unchanged.
func ChangePointByValue(p Point) {
	p.x = 0
	p.y = 0
}

// ChangePointByReference modifies the original point by using a pointer.
// As Go directly accesses the memory location of p, the original point is modified.
func ChangePointByReference(p *Point) {
	p.x = 0
	p.y = 0
}

func main() {
	p := Point{x: 50, y: 25}
	fmt.Println("Initial Point:", p)

	ChangePointByValue(p)
	fmt.Println("After Modification by Value:", p)

	// Using "&" operator to pass memory address of the variable for modification.
	ChangePointByReference(&p)
	fmt.Println("After Modification by Reference:", p)
}

// Pass by pointer when:
// - You want the function to modify the variable
// - The data is large (e.g., a large struct) and you want to avoid copying it
// Pass by value when dealing with maps or slices as they are reference types by default.
