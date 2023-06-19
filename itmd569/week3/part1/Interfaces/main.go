package main

import (
	"fmt"
)

// An interface type in Go is kind of like a definition. It defines and describes the exact methods that some other type must have.

// shapes interface, meaning that anything that implements it MUST also implement an area function
type shapes interface {
	area() float64
}

// declare square and triangle structs, with area methods for each
// this means that they satisfy the definition of our shapes interface
type square struct {
	width  float64
	height float64
}

func (s square) area() float64 {
	squareArea := s.width * s.height
	return squareArea
}

type triangle struct {
	width  float64
	height float64
}

func (t triangle) area() float64 {
	triangleArea := (t.width * t.height) / 2
	return triangleArea

}

// we can declare functions that take types of interfaces
// here, this calculateArea function will accept any struct which statifies the shapes interface type

func calcArea(s shapes) float64 {
	shapeArea := s.area() // can be triangle or square or any new struct we create that has an area method
	fmt.Println(shapeArea)
	return shapeArea
}

func main() {

	// notice how we calculate the area for both the square and the triangle by calling the same function
	square := square{width: 4, height: 4}
	squareArea := calcArea(square)
	fmt.Println("Area of the square is ", squareArea)

	triangle := triangle{width: 4, height: 6}
	triangleArea := calcArea(triangle)
	fmt.Println("Area of the triangle is ", triangleArea)
}
